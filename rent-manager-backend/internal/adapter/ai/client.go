package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	http    *http.Client
}

var _ port.DocumentGateway = (*Client)(nil)

func NewClient(baseURL string, timeout time.Duration) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") || u.User != nil || u.RawQuery != "" || u.Fragment != "" || timeout <= 0 {
		return nil, fmt.Errorf("invalid AI service URL or timeout")
	}
	return &Client{baseURL: strings.TrimRight(baseURL, "/"), http: &http.Client{
		Timeout: timeout,
		// Never send document contents to an unexpected redirect destination.
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}}, nil
}
func (c *Client) path(scope domain.DocumentScope) string {
	return fmt.Sprintf("%s/properties/%d/contracts/%d", c.baseURL, scope.PropertyID, scope.ContractID)
}
func unavailable(cause error) error {
	return domain.WrapAppError(domain.ErrCodeAIUnavailable, "The AI service is unavailable", cause)
}
func (c *Client) request(ctx context.Context, method, endpoint, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, unavailable(err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	// Deliberately do not forward cookies, Authorization, or arbitrary caller headers.
	response, err := c.http.Do(req)
	if err != nil {
		var timeout interface{ Timeout() bool }
		if errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &timeout) && timeout.Timeout()) {
			return nil, domain.WrapAppError(domain.ErrCodeAITimeout, "The AI service timed out", err)
		}
		return nil, unavailable(err)
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return response, nil
	}
	defer response.Body.Close()
	// Translate status only: never leak upstream stack traces or internal URLs.
	code, message := domain.ErrCodeAIUnavailable, "The AI service is unavailable"
	switch response.StatusCode {
	case 400:
		code, message = domain.ErrCodeInvalidData, "Invalid document or question"
	case 404:
		code, message = domain.ErrCodeDataNotFound, "Document not found"
	case 409:
		code, message = domain.ErrCodeConflictingData, "Conflicting document operation"
	case 413:
		code, message = domain.ErrCodeDocumentTooLarge, "PDF files must not exceed 20 MiB"
	case 415:
		code, message = domain.ErrCodeDocumentUnsupported, "Only PDF documents are supported"
	case 422:
		code, message = domain.ErrCodeDocumentUnprocessable, "The PDF must contain readable text and must not be encrypted"
	case 504:
		code, message = domain.ErrCodeAITimeout, "The AI service timed out"
	}
	return nil, domain.NewAppError(code, message)
}

type documentResponse struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"contentType"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (d documentResponse) domain() domain.Document {
	return domain.Document{ID: d.ID, Type: d.Type, Filename: d.Filename, ContentType: d.ContentType, Size: d.Size, CreatedAt: d.CreatedAt}
}
func decode(response *http.Response, target any) error {
	defer response.Body.Close()
	if err := json.NewDecoder(io.LimitReader(response.Body, 4*1024*1024)).Decode(target); err != nil {
		return unavailable(err)
	}
	return nil
}
func (c *Client) Upload(ctx context.Context, scope domain.DocumentScope, upload domain.DocumentUpload) (*domain.Document, error) {
	// Bounded by the HTTP adapter and use case; rebuilding multipart preserves PDF bytes.
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	header := make(textproto.MIMEHeader)
	// FormatMediaType quotes the supplied name and prevents multipart header injection.
	header.Set("Content-Disposition", mime.FormatMediaType("form-data", map[string]string{"name": "file", "filename": upload.Filename}))
	header.Set("Content-Type", "application/pdf")
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, unavailable(err)
	}
	n, err := io.Copy(part, io.LimitReader(upload.Content, domain.MaxDocumentBytes+1))
	if err != nil {
		return nil, unavailable(err)
	}
	if n > domain.MaxDocumentBytes {
		return nil, domain.NewAppError(domain.ErrCodeDocumentTooLarge, "PDF files must not exceed 20 MiB")
	}
	if err = writer.WriteField("type", upload.Type); err != nil {
		return nil, unavailable(err)
	}
	if err = writer.Close(); err != nil {
		return nil, unavailable(err)
	}
	response, err := c.request(ctx, http.MethodPost, c.path(scope)+"/documents", writer.FormDataContentType(), &body)
	if err != nil {
		return nil, err
	}
	var result documentResponse
	if err = decode(response, &result); err != nil {
		return nil, err
	}
	document := result.domain()
	return &document, nil
}
func (c *Client) List(ctx context.Context, scope domain.DocumentScope) ([]domain.Document, error) {
	response, err := c.request(ctx, http.MethodGet, c.path(scope)+"/documents", "", nil)
	if err != nil {
		return nil, err
	}
	var result []documentResponse
	if err = decode(response, &result); err != nil {
		return nil, err
	}
	documents := make([]domain.Document, 0, len(result))
	for _, item := range result {
		documents = append(documents, item.domain())
	}
	return documents, nil
}
func (c *Client) Download(ctx context.Context, scope domain.DocumentScope, id string, attachment bool) (*domain.DocumentDownload, error) {
	response, err := c.request(ctx, http.MethodGet, c.path(scope)+"/documents/"+url.PathEscape(id)+fmt.Sprintf("?download=%t", attachment), "", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := io.ReadAll(io.LimitReader(response.Body, domain.MaxDocumentBytes+1))
	if err != nil {
		return nil, unavailable(err)
	}
	if int64(len(content)) > domain.MaxDocumentBytes {
		return nil, unavailable(fmt.Errorf("upstream PDF exceeds limit"))
	}
	return &domain.DocumentDownload{Content: content, ContentDisposition: response.Header.Get("Content-Disposition")}, nil
}
func (c *Client) Delete(ctx context.Context, scope domain.DocumentScope, id string) error {
	response, err := c.request(ctx, http.MethodDelete, c.path(scope)+"/documents/"+url.PathEscape(id), "", nil)
	if err == nil {
		response.Body.Close()
	}
	return err
}
func (c *Client) Ask(ctx context.Context, scope domain.DocumentScope, question string) (string, error) {
	body, err := json.Marshal(struct {
		Question string `json:"question"`
	}{question})
	if err != nil {
		return "", unavailable(err)
	}
	response, err := c.request(ctx, http.MethodPost, c.path(scope)+"/chat", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	var result struct {
		Answer string `json:"answer"`
	}
	if err = decode(response, &result); err != nil {
		return "", err
	}
	return result.Answer, nil
}
