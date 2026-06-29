package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetOwnerByIDUseCase struct {
	ownerRepository port.OwnerRepository
}

func NewGetOwnerByIDUseCase(ownerRepository port.OwnerRepository) *GetOwnerByIDUseCase {
	return &GetOwnerByIDUseCase{
		ownerRepository: ownerRepository,
	}
}

func (uc *GetOwnerByIDUseCase) GetOwnerByID(ctx context.Context, id int64) (*domain.Owner, error) {
	owner, err := uc.ownerRepository.GetOwnerByID(ctx, id)
	if err != nil {
		slog.Error("[GetOwnerByIDUseCase] repository get by id failed", "err", err)
		return nil, err
	}

	return owner, nil
}
