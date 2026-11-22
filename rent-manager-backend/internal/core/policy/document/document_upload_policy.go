package policy

import (
	"context"
	"rent-manager-backend/internal/core/domain"
)

type DocumentUploadPolicy struct{}

func NewDocumentUploadPolicy() *DocumentUploadPolicy { return &DocumentUploadPolicy{} }

func (policy *DocumentUploadPolicy) Execute(_ context.Context, upload *domain.DocumentUpload) error {
	if upload == nil {
		return domain.NewAppError(domain.ErrCodeInvalidData, "A nonempty PDF is required")
	}
	if upload.Size > domain.MaxDocumentBytes {
		return domain.NewAppError(domain.ErrCodeDocumentTooLarge, "PDF files must not exceed 20 MiB")
	}
	if upload.Size <= 0 || upload.Content == nil {
		return domain.NewAppError(domain.ErrCodeInvalidData, "A nonempty PDF is required")
	}
	switch upload.Type {
	case "RENT_CONTRACT", "GUARANTEE", "ANNEX", "INVENTORY", "HANDOVER", "OTHER":
		return nil
	default:
		return domain.NewAppError(domain.ErrCodeInvalidData, "Invalid document type")
	}
}
