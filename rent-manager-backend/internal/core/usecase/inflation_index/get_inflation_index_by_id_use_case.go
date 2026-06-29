package inflation_index

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetInflationIndexByIDUseCase struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewGetInflationIndexByIDUseCase(
	inflationIndexRepository port.InflationIndexRepository,
) *GetInflationIndexByIDUseCase {
	return &GetInflationIndexByIDUseCase{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (uc *GetInflationIndexByIDUseCase) Execute(ctx context.Context, id int64) (*domain.InflationIndex, error) {
	return uc.inflationIndexRepository.GetByID(ctx, id)
}
