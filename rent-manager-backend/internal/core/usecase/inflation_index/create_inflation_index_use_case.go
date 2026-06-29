package inflation_index

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/inflation_index"
	"rent-manager-backend/internal/core/port"
)

type CreateInflationIndexUseCase struct {
	inflationIndexRepository   port.InflationIndexRepository
	inflationIndexCreatePolicy *inflation_index.InflationIndexCreatePolicy
}

func NewCreateInflationIndexUseCase(
	inflationIndexRepository port.InflationIndexRepository,
	inflationIndexCreatePolicy *inflation_index.InflationIndexCreatePolicy,
) *CreateInflationIndexUseCase {
	return &CreateInflationIndexUseCase{
		inflationIndexRepository:   inflationIndexRepository,
		inflationIndexCreatePolicy: inflationIndexCreatePolicy,
	}
}

func (uc *CreateInflationIndexUseCase) Execute(
	ctx context.Context,
	inflationIndex *domain.InflationIndex,
) (*domain.InflationIndex, error) {
	if err := uc.inflationIndexCreatePolicy.Validate(ctx, inflationIndex); err != nil {
		return nil, err
	}

	if err := uc.inflationIndexRepository.Create(ctx, inflationIndex); err != nil {
		return nil, err
	}

	return inflationIndex, nil
}
