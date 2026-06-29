package inflation_index

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListInflationIndexesByPeriodRangeUseCase struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewListInflationIndexesByPeriodRangeUseCase(
	inflationIndexRepository port.InflationIndexRepository,
) *ListInflationIndexesByPeriodRangeUseCase {
	return &ListInflationIndexesByPeriodRangeUseCase{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (uc *ListInflationIndexesByPeriodRangeUseCase) Execute(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.InflationIndex, error) {
	return uc.inflationIndexRepository.ListByPeriodRange(ctx, from, to)
}
