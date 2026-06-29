package inflation_index

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetInflationIndexByPeriodUseCase struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewGetInflationIndexByPeriodUseCase(
	inflationIndexRepository port.InflationIndexRepository,
) *GetInflationIndexByPeriodUseCase {
	return &GetInflationIndexByPeriodUseCase{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (uc *GetInflationIndexByPeriodUseCase) Execute(
	ctx context.Context,
	period time.Time,
) (*domain.InflationIndex, error) {
	return uc.inflationIndexRepository.GetByPeriod(
		ctx,
		period,
	)
}
