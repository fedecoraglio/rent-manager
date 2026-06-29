package inflation_index

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	inflationIndexPolicy "rent-manager-backend/internal/core/policy/inflation_index"
	"rent-manager-backend/internal/core/port"
)

type UpdateInflationIndexUseCase struct {
	inflationIndexRepository   port.InflationIndexRepository
	inflationIndexUpdatePolicy *inflationIndexPolicy.InflationIndexUpdatePolicy
}

func NewUpdateInflationIndexUseCase(
	inflationIndexRepository port.InflationIndexRepository,
	inflationIndexUpdatePolicy *inflationIndexPolicy.InflationIndexUpdatePolicy,
) *UpdateInflationIndexUseCase {
	return &UpdateInflationIndexUseCase{
		inflationIndexRepository:   inflationIndexRepository,
		inflationIndexUpdatePolicy: inflationIndexUpdatePolicy,
	}
}

func (uc *UpdateInflationIndexUseCase) Execute(
	ctx context.Context,
	inflationIndex *domain.InflationIndex,
) (*domain.InflationIndex, error) {
	if err := uc.inflationIndexUpdatePolicy.Validate(ctx, inflationIndex); err != nil {
		return nil, err
	}

	if err := uc.inflationIndexRepository.Update(ctx, inflationIndex); err != nil {
		return nil, err
	}

	return inflationIndex, nil
}
