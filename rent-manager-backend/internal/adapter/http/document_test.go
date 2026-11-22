package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"rent-manager-backend/internal/adapter/ai"
	"rent-manager-backend/internal/core/domain"
	documentPolicy "rent-manager-backend/internal/core/policy/document"
	"rent-manager-backend/internal/core/port"
	documentUseCase "rent-manager-backend/internal/core/usecase/document"
	"strings"
	"testing"
	"time"
)

type documentContractRepository struct {
	port.RentalContractRepository
	calls int
}

func (r *documentContractRepository) GetRentalContractByID(ctx context.Context, id int64) (*domain.RentalContract, error) {
	r.calls++
	if id == 99 {
		return nil, domain.ErrRentalContractNotFound
	}
	return &domain.RentalContract{ID: id, PropertyID: 12}, nil
}

type documentTokenProvider struct{ port.TokenProvider }

func (documentTokenProvider) ValidateToken(token string) (*domain.TokenClaims, error) {
	if token != "valid" {
		return nil, domain.ErrUnauthorized
	}
	return &domain.TokenClaims{}, nil
}
func TestDocumentRoutesAuthenticateAndResolveScope(t *testing.T) {
	gin.SetMode(gin.TestMode)
	calls := 0
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Header.Get("Authorization") != "" {
			t.Error("JWT leaked upstream")
		}
		if r.URL.Path != "/properties/12/contracts/34/documents" {
			t.Errorf("unexpected scope: %s", r.URL.Path)
		}
		io.WriteString(w, "[]")
	}))
	defer upstream.Close()
	client, _ := ai.NewClient(upstream.URL, time.Second)
	repository := &documentContractRepository{}
	handler := NewDocumentHandler(
		documentUseCase.NewUploadDocumentUseCase(repository, client, documentPolicy.NewDocumentUploadPolicy()),
		documentUseCase.NewListDocumentsUseCase(repository, client),
		documentUseCase.NewDownloadDocumentUseCase(repository, client),
		documentUseCase.NewDeleteDocumentUseCase(repository, client),
		documentUseCase.NewAskContractQuestionUseCase(repository, client, documentPolicy.NewContractQuestionPolicy()),
	)
	router := gin.New()
	group := router.Group("/v1/rental-contracts", AuthMiddleware(documentTokenProvider{}))
	handler.Register(group)
	for _, item := range []struct {
		method, path, token, body string
		status                    int
	}{
		{"GET", "/v1/rental-contracts/34/documents", "", "", 401},
		{"GET", "/v1/rental-contracts/34/documents", "invalid", "", 401},
		{"POST", "/v1/rental-contracts/34/documents", "", "", 401},
		{"DELETE", "/v1/rental-contracts/34/documents/abc", "", "", 401},
		{"GET", "/v1/rental-contracts/34/documents/abc", "", "", 401},
		{"POST", "/v1/rental-contracts/34/chat", "", "", 401},
		{"GET", "/v1/rental-contracts/99/documents", "valid", "", 404},
		{"GET", "/v1/rental-contracts/0/documents", "valid", "", 400},
		{"POST", "/v1/rental-contracts/34/chat", "valid", `{"question":" "}`, 400},
		{"GET", "/v1/rental-contracts/34/documents", "valid", "", 200},
	} {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(item.method, item.path, strings.NewReader(item.body))
		if item.token != "" {
			request.Header.Set("Authorization", "Bearer "+item.token)
		}
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		if response.Code != item.status {
			t.Errorf("%s %s: status %d body=%s", item.method, item.path, response.Code, response.Body.String())
		}
		if item.status == 200 && strings.TrimSpace(response.Body.String()) != `{"data":[]}` {
			t.Errorf("unexpected envelope: %s", response.Body.String())
		}
	}
	if calls != 1 {
		t.Errorf("unexpected upstream calls: %d", calls)
	}
}
