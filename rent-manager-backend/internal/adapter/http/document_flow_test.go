package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"rent-manager-backend/internal/adapter/ai"
	documentPolicy "rent-manager-backend/internal/core/policy/document"
	documentUseCase "rent-manager-backend/internal/core/usecase/document"
	"strings"
	"testing"
	"time"
)

func TestDocumentHTTPFlow(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pdf := []byte("%PDF-1.4\ncontract bytes\x00\xff")
	upstreamCalls := 0
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upstreamCalls++
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("user credentials forwarded")
		}
		switch {
		case r.Method == "POST" && r.URL.Path == "/properties/12/contracts/34/documents":
			if err := r.ParseMultipartForm(1024); err != nil {
				t.Error(err)
				w.WriteHeader(400)
				return
			}
			defer r.MultipartForm.RemoveAll()
			file, _, err := r.FormFile("file")
			if err != nil {
				t.Error(err)
				return
			}
			defer file.Close()
			content, _ := io.ReadAll(file)
			if !bytes.Equal(pdf, content) {
				t.Error("PDF bytes changed")
			}
			w.WriteHeader(201)
			io.WriteString(w, `{"id":"12_34_RENT_CONTRACT","type":"RENT_CONTRACT","filename":"contract.pdf","contentType":"application/pdf","size":25,"createdAt":"2026-11-22T00:00:00Z"}`)
		case r.Method == "GET" && strings.HasSuffix(r.URL.Path, "/documents/12_34_RENT_CONTRACT"):
			if r.URL.Query().Get("download") != "true" {
				t.Error("download flag lost")
			}
			w.Header().Set("Content-Disposition", `attachment; filename="contract.pdf"`)
			w.Write(pdf)
		case r.Method == "DELETE":
			w.WriteHeader(204)
		case r.Method == "POST" && r.URL.Path == "/properties/12/contracts/34/chat":
			var body struct {
				Question string `json:"question"`
			}
			json.NewDecoder(r.Body).Decode(&body)
			if body.Question != "When does the contract end?" {
				t.Errorf("unexpected question: %q", body.Question)
			}
			io.WriteString(w, `{"answer":"December 2027"}`)
		default:
			t.Errorf("unexpected upstream request: %s %s", r.Method, r.URL)
			w.WriteHeader(500)
		}
	}))
	defer upstream.Close()
	client, _ := ai.NewClient(upstream.URL, time.Second)
	handler := NewDocumentHandler(
		documentUseCase.NewUploadDocumentUseCase(&documentContractRepository{}, client, documentPolicy.NewDocumentUploadPolicy()),
		documentUseCase.NewListDocumentsUseCase(&documentContractRepository{}, client),
		documentUseCase.NewDownloadDocumentUseCase(&documentContractRepository{}, client),
		documentUseCase.NewDeleteDocumentUseCase(&documentContractRepository{}, client),
		documentUseCase.NewAskContractQuestionUseCase(&documentContractRepository{}, client, documentPolicy.NewContractQuestionPolicy()),
	)
	router := gin.New()
	handler.Register(router.Group("/v1/rental-contracts", AuthMiddleware(documentTokenProvider{})))
	var form bytes.Buffer
	writer := multipart.NewWriter(&form)
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", `form-data; name="file"; filename="contract.pdf"`)
	header.Set("Content-Type", "application/pdf")
	part, _ := writer.CreatePart(header)
	part.Write(pdf)
	writer.WriteField("type", "RENT_CONTRACT")
	writer.WriteField("propertyId", "999")
	writer.Close()
	request := httptest.NewRequest("POST", "/v1/rental-contracts/34/documents", &form)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer valid")
	request.Header.Set("Cookie", "session=secret")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != 201 || response.Header().Get("Location") != "/v1/rental-contracts/34/documents/12_34_RENT_CONTRACT" || !strings.Contains(response.Body.String(), `"data":`) {
		t.Fatalf("upload status=%d body=%s", response.Code, response.Body.String())
	}
	for _, item := range []struct {
		method, path, body string
		status             int
	}{
		{"GET", "/v1/rental-contracts/34/documents/12_34_RENT_CONTRACT?download=true", "", 200},
		{"POST", "/v1/rental-contracts/34/chat", `{"question":"When does the contract end?"}`, 200},
		{"DELETE", "/v1/rental-contracts/34/documents/12_34_RENT_CONTRACT", "", 204},
	} {
		req := httptest.NewRequest(item.method, item.path, strings.NewReader(item.body))
		req.Header.Set("Authorization", "Bearer valid")
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		if res.Code != item.status {
			t.Fatalf("%s: status=%d body=%s", item.path, res.Code, res.Body.String())
		}
		if item.method == "GET" && (!bytes.Equal(res.Body.Bytes(), pdf) || res.Header().Get("X-Content-Type-Options") != "nosniff" || res.Header().Get("Content-Disposition") == "") {
			t.Error("PDF response changed")
		}
		if strings.HasSuffix(item.path, "/chat") && strings.TrimSpace(res.Body.String()) != `{"data":{"answer":"December 2027"}}` {
			t.Error("chat envelope changed")
		}
	}
	if upstreamCalls != 4 {
		t.Errorf("expected 4 calls, received %d", upstreamCalls)
	}
}
