package rental_contract

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetRentalContractByIDUseCase struct {
	rentalContractRepository port.RentalContractRepository
}

func NewGetRentalContractByIDUseCase(
	rentalContractRepository port.RentalContractRepository,
) *GetRentalContractByIDUseCase {
	return &GetRentalContractByIDUseCase{
		rentalContractRepository: rentalContractRepository,
	}
}

func (uc *GetRentalContractByIDUseCase) GetRentalContractByID(
	ctx context.Context,
	id int64,
) (*domain.RentalContract, error) {
	rentalContract, err := uc.rentalContractRepository.GetRentalContractByID(ctx, id)
	if err != nil {
		slog.Error("[GetRentalContractByIDUseCase] repository get by id failed", "err", err)
		return nil, err
	}

	return rentalContract, nil
}
