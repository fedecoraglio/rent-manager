package document

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	documentPolicy "rent-manager-backend/internal/core/policy/document"
	"rent-manager-backend/internal/core/port"
)

type UploadDocumentUseCase struct {
	contracts    port.RentalContractRepository
	documents    port.DocumentGateway
	uploadPolicy *documentPolicy.DocumentUploadPolicy
}

func NewUploadDocumentUseCase(contracts port.RentalContractRepository, documents port.DocumentGateway, uploadPolicy *documentPolicy.DocumentUploadPolicy) *UploadDocumentUseCase {
	return &UploadDocumentUseCase{contracts: contracts, documents: documents, uploadPolicy: uploadPolicy}
}

func (uc *UploadDocumentUseCase) Execute(ctx context.Context, id int64, upload domain.DocumentUpload) (*domain.Document, error) {
	scope, err := resolveDocumentScope(ctx, uc.contracts, id)
	if err != nil {
		return nil, err
	}
	if err := uc.uploadPolicy.Execute(ctx, &upload); err != nil {
		return nil, err
	}
	return uc.documents.Upload(ctx, scope, upload)
}
