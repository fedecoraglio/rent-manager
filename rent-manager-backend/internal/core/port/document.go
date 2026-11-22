package port

import (
	"context"
	"rent-manager-backend/internal/core/domain"
)

// DocumentGateway is the outbound port to the private document and RAG service.
type DocumentGateway interface {
	Upload(context.Context, domain.DocumentScope, domain.DocumentUpload) (*domain.Document, error)
	List(context.Context, domain.DocumentScope) ([]domain.Document, error)
	Download(context.Context, domain.DocumentScope, string, bool) (*domain.DocumentDownload, error)
	Delete(context.Context, domain.DocumentScope, string) error
	Ask(context.Context, domain.DocumentScope, string) (string, error)
}
