package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetOwnerByDocumentNumberUseCase struct {
	ownerRepository port.OwnerRepository
}

func NewGetOwnerByDocumentNumberUseCase(
	ownerRepository port.OwnerRepository,
) *GetOwnerByDocumentNumberUseCase {
	return &GetOwnerByDocumentNumberUseCase{
		ownerRepository: ownerRepository,
	}
}

func (uc *GetOwnerByDocumentNumberUseCase) GetOwnerByDocumentNumber(
	ctx context.Context,
	documentNumber string,
) (*domain.Owner, error) {
	owner, err := uc.ownerRepository.GetOwnerByDocumentNumber(ctx, documentNumber)
	if err != nil {
		slog.Error("[GetOwnerByDocumentNumberUseCase] repository get by document number failed", "err", err)
		return nil, err
	}

	return owner, nil
}
