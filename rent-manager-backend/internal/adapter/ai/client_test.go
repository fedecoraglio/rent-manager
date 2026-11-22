package ai

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"rent-manager-backend/internal/core/domain"
	"strings"
	"testing"
	"time"
)

func TestUploadPreservesBytesAndDoesNotForwardCredentials(t *testing.T) {
	pdf := "%PDF-1.4\noriginal bytes\x00\xff"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/properties/12/contracts/34/documents" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("credentials forwarded")
		}
		if err := r.ParseMultipartForm(1024); err != nil {
			t.Error(err)
			w.WriteHeader(400)
			return
		}
		defer r.MultipartForm.RemoveAll()
		file, header, err := r.FormFile("file")
		if err != nil {
			t.Error(err)
			w.WriteHeader(400)
			return
		}
		defer file.Close()
		content, _ := io.ReadAll(file)
		if string(content) != pdf || header.Filename != "contract.pdf" || header.Header.Get("Content-Type") != "application/pdf" || r.FormValue("type") != "RENT_CONTRACT" {
			t.Error("multipart changed")
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"id":"12_34_RENT_CONTRACT","type":"RENT_CONTRACT","filename":"contract.pdf","contentType":"application/pdf","size":20,"createdAt":"2026-11-22T00:00:00Z"}`)
	}))
	defer server.Close()
	client, err := NewClient(server.URL, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.Upload(context.Background(), domain.DocumentScope{PropertyID: 12, ContractID: 34}, domain.DocumentUpload{Filename: "contract.pdf", Type: "RENT_CONTRACT", Size: int64(len(pdf)), Content: strings.NewReader(pdf)})
	if err != nil || result.ID != "12_34_RENT_CONTRACT" {
		t.Fatalf("result=%v err=%v", result, err)
	}
}

func TestUpstreamErrorsAreTranslatedWithoutLeakingBody(t *testing.T) {
	for _, item := range []struct {
		status int
		code   domain.ErrorCode
	}{{404, domain.ErrCodeDataNotFound}, {413, domain.ErrCodeDocumentTooLarge}, {422, domain.ErrCodeDocumentUnprocessable}, {500, domain.ErrCodeAIUnavailable}, {503, domain.ErrCodeAIUnavailable}} {
		t.Run(http.StatusText(item.status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(item.status)
				io.WriteString(w, "secret internal stack trace")
			}))
			defer server.Close()
			client, _ := NewClient(server.URL, time.Second)
			_, err := client.List(context.Background(), domain.DocumentScope{PropertyID: 1, ContractID: 2})
			var appErr *domain.AppError
			if !errors.As(err, &appErr) || appErr.Code != item.code || strings.Contains(err.Error(), "secret") {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestTimeoutAndRedirect(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { time.Sleep(100 * time.Millisecond) }))
		defer server.Close()
		client, _ := NewClient(server.URL, 50*time.Millisecond)
		_, err := client.Ask(context.Background(), domain.DocumentScope{PropertyID: 1, ContractID: 2}, "question")
		var appErr *domain.AppError
		if !errors.As(err, &appErr) || appErr.Code != domain.ErrCodeAITimeout {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("redirect", func(t *testing.T) {
		called := false
		target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { called = true }))
		defer target.Close()
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, target.URL, 307) }))
		defer server.Close()
		client, _ := NewClient(server.URL, time.Second)
		_, err := client.Ask(context.Background(), domain.DocumentScope{PropertyID: 1, ContractID: 2}, "question")
		if err == nil || called {
			t.Fatal("redirect was followed")
		}
	})
}
