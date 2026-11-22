package document

import (
	"context"
	documentPolicy "rent-manager-backend/internal/core/policy/document"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type DownloadDocumentUseCase struct {
	contracts port.RentalContractRepository
	documents port.DocumentGateway
}

func NewDownloadDocumentUseCase(contracts port.RentalContractRepository, documents port.DocumentGateway) *DownloadDocumentUseCase {
	return &DownloadDocumentUseCase{contracts: contracts, documents: documents}
}

func (uc *DownloadDocumentUseCase) Execute(ctx context.Context, id int64, docID string, attachment bool) (*domain.DocumentDownload, error) {
	if err := documentPolicy.ValidateDocumentID(docID); err != nil {
		return nil, err
	}
	scope, err := resolveDocumentScope(ctx, uc.contracts, id)
	if err != nil {
		return nil, err
	}
	return uc.documents.Download(ctx, scope, docID, attachment)
}
