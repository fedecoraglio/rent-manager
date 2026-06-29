package inflation_index

import (
	"context"
	"errors"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type InflationIndexCreatePolicy struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewInflationIndexCreatePolicy(
	inflationIndexRepository port.InflationIndexRepository,
) *InflationIndexCreatePolicy {
	return &InflationIndexCreatePolicy{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (p *InflationIndexCreatePolicy) Validate(ctx context.Context, inflationIndex *domain.InflationIndex) error {
	if inflationIndex.Period.IsZero() {
		return domain.ErrInflationIndexPeriodRequired
	}

	if inflationIndex.Percentage < 0 {
		return domain.ErrInflationIndexPercentageInvalid
	}

	existingInflationIndex, err := p.inflationIndexRepository.GetByPeriod(ctx, inflationIndex.Period)
	if err == nil && existingInflationIndex != nil {
		return domain.ErrInflationIndexAlreadyExists
	}

	if err != nil && !errors.Is(err, domain.ErrInflationIndexNotFound) {
		return err
	}

	return nil
}
