package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListOwnersUseCase struct {
	ownerRepository port.OwnerRepository
}

func NewListOwnersUseCase(ownerRepository port.OwnerRepository) *ListOwnersUseCase {
	return &ListOwnersUseCase{
		ownerRepository: ownerRepository,
	}
}

func (uc *ListOwnersUseCase) ListOwners(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.Owner, error) {
	owners, err := uc.ownerRepository.ListOwners(ctx, page, limit)
	if err != nil {
		slog.Error("[ListOwnersUseCase] repository list failed", "err", err)
		return nil, err
	}

	return owners, nil
}
