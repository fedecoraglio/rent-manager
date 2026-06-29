package inflation_index

import (
	"context"
	"errors"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type InflationIndexUpdatePolicy struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewInflationIndexUpdatePolicy(
	inflationIndexRepository port.InflationIndexRepository,
) *InflationIndexUpdatePolicy {
	return &InflationIndexUpdatePolicy{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (p *InflationIndexUpdatePolicy) Validate(
	ctx context.Context,
	inflationIndex *domain.InflationIndex,
) error {
	if inflationIndex.ID <= 0 {
		return domain.ErrInflationIndexNotFound
	}

	if inflationIndex.Period.IsZero() {
		return domain.ErrInflationIndexPeriodRequired
	}

	if inflationIndex.Percentage < 0 {
		return domain.ErrInflationIndexPercentageInvalid
	}

	currentInflationIndex, err := p.inflationIndexRepository.GetByID(ctx, inflationIndex.ID)
	if err != nil {
		return err
	}

	existingInflationIndex, err := p.inflationIndexRepository.GetByPeriod(ctx, inflationIndex.Period)
	if err == nil &&
		existingInflationIndex != nil &&
		existingInflationIndex.ID != currentInflationIndex.ID {
		return domain.ErrInflationIndexAlreadyExists
	}

	if err != nil && !errors.Is(err, domain.ErrInflationIndexNotFound) {
		return err
	}

	return nil
}
