package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type SearchOwnersUseCase struct {
	ownerRepository port.OwnerRepository
}

func NewSearchOwnersUseCase(ownerRepository port.OwnerRepository) *SearchOwnersUseCase {
	return &SearchOwnersUseCase{
		ownerRepository: ownerRepository,
	}
}

func (uc *SearchOwnersUseCase) SearchOwners(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Owner, error) {
	owners, err := uc.ownerRepository.SearchOwners(ctx, value, page, limit)
	if err != nil {
		slog.Error("[SearchOwnersUseCase] repository search failed", "err", err)
		return nil, err
	}

	return owners, nil
}
