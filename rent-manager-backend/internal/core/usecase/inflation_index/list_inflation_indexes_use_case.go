package inflation_index

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListInflationIndexesUseCase struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewListInflationIndexesUseCase(
	inflationIndexRepository port.InflationIndexRepository,
) *ListInflationIndexesUseCase {
	return &ListInflationIndexesUseCase{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (uc *ListInflationIndexesUseCase) Execute(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.InflationIndex, error) {
	return uc.inflationIndexRepository.List(ctx, page, limit)
}
