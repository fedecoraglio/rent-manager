package document

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListDocumentsUseCase struct {
	contracts port.RentalContractRepository
	documents port.DocumentGateway
}

func NewListDocumentsUseCase(contracts port.RentalContractRepository, documents port.DocumentGateway) *ListDocumentsUseCase {
	return &ListDocumentsUseCase{contracts: contracts, documents: documents}
}

func (uc *ListDocumentsUseCase) Execute(ctx context.Context, id int64) ([]domain.Document, error) {
	scope, err := resolveDocumentScope(ctx, uc.contracts, id)
	if err != nil {
		return nil, err
	}
	return uc.documents.List(ctx, scope)
}
