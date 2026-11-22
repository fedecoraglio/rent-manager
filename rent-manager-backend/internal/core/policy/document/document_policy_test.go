package policy

import (
	"context"
	"errors"
	"rent-manager-backend/internal/core/domain"
	"strings"
	"testing"
)

func assertCode(t *testing.T, err error, code domain.ErrorCode) {
	t.Helper()
	if code == "" {
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	var appErr *domain.AppError
	if !errors.As(err, &appErr) || appErr.Code != code {
		t.Fatalf("expected %s, received %v", code, err)
	}
}

func TestUploadPolicy(t *testing.T) {
	policy := NewDocumentUploadPolicy()
	valid := func() *domain.DocumentUpload {
		return &domain.DocumentUpload{Size: 1, Content: strings.NewReader("pdf"), Type: "RENT_CONTRACT"}
	}
	for _, kind := range []string{"RENT_CONTRACT", "GUARANTEE", "ANNEX", "INVENTORY", "HANDOVER", "OTHER"} {
		t.Run(kind, func(t *testing.T) {
			upload := valid()
			upload.Type = kind
			assertCode(t, policy.Execute(context.Background(), upload), "")
		})
	}
	for _, test := range []struct {
		name   string
		change func(*domain.DocumentUpload)
		code   domain.ErrorCode
	}{
		{"empty", func(u *domain.DocumentUpload) { u.Size = 0 }, domain.ErrCodeInvalidData},
		{"negative size", func(u *domain.DocumentUpload) { u.Size = -1 }, domain.ErrCodeInvalidData},
		{"missing content", func(u *domain.DocumentUpload) { u.Content = nil }, domain.ErrCodeInvalidData},
		{"unknown type", func(u *domain.DocumentUpload) { u.Type = "UNKNOWN" }, domain.ErrCodeInvalidData},
		{"exact size limit", func(u *domain.DocumentUpload) { u.Size = domain.MaxDocumentBytes }, ""},
		{"oversize", func(u *domain.DocumentUpload) { u.Size = domain.MaxDocumentBytes + 1 }, domain.ErrCodeDocumentTooLarge},
	} {
		t.Run(test.name, func(t *testing.T) {
			u := valid()
			test.change(u)
			assertCode(t, policy.Execute(context.Background(), u), test.code)
		})
	}
	assertCode(t, policy.Execute(context.Background(), nil), domain.ErrCodeInvalidData)
}

func TestQuestionPolicy(t *testing.T) {
	policy := NewContractQuestionPolicy()
	for _, test := range []struct {
		input, normalized string
		code              domain.ErrorCode
	}{
		{"  When does it end? \n", "When does it end?", ""},
		{" \t\n", "", domain.ErrCodeInvalidData},
		{strings.Repeat("a", 8000), strings.Repeat("a", 8000), ""},
		{strings.Repeat("a", 8001), "", domain.ErrCodeInvalidData},
		{strings.Repeat("é", 4001), "", domain.ErrCodeInvalidData},
	} {
		question := &domain.ContractQuestion{Question: test.input}
		assertCode(t, policy.Execute(context.Background(), question), test.code)
		if test.code == "" && question.Question != test.normalized {
			t.Fatal("question was not normalized")
		}
	}
	assertCode(t, policy.Execute(context.Background(), nil), domain.ErrCodeInvalidData)
}

func TestScopeIdentifiers(t *testing.T) {
	for _, id := range []string{"12_34_RENT_CONTRACT", "A~B_C_ANNEX_123-456", strings.Repeat("x", 256)} {
		assertCode(t, ValidateDocumentID(id), "")
	}
	for _, id := range []string{"", "../secret", "a/b", "id?download=true", strings.Repeat("x", 257)} {
		assertCode(t, ValidateDocumentID(id), domain.ErrCodeInvalidData)
	}
	assertCode(t, ValidateRentalContractID(1), "")
	assertCode(t, ValidateRentalContractID(0), domain.ErrCodeInvalidData)
	assertCode(t, ValidateRentalContractID(-1), domain.ErrCodeInvalidData)
}
