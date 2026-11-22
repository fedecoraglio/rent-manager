package document

import (
	"context"
	documentPolicy "rent-manager-backend/internal/core/policy/document"

	"rent-manager-backend/internal/core/port"
)

type DeleteDocumentUseCase struct {
	contracts port.RentalContractRepository
	documents port.DocumentGateway
}

func NewDeleteDocumentUseCase(contracts port.RentalContractRepository, documents port.DocumentGateway) *DeleteDocumentUseCase {
	return &DeleteDocumentUseCase{contracts: contracts, documents: documents}
}

func (uc *DeleteDocumentUseCase) Execute(ctx context.Context, id int64, docID string) error {
	if err := documentPolicy.ValidateDocumentID(docID); err != nil {
		return err
	}
	scope, err := resolveDocumentScope(ctx, uc.contracts, id)
	if err != nil {
		return err
	}
	return uc.documents.Delete(ctx, scope, docID)
}
