package rental_contract

import (
	"context"
	"log/slog"
	"rent-manager-backend/internal/core/policy/shared"

	"rent-manager-backend/internal/core/domain"
	policyRentContract "rent-manager-backend/internal/core/policy/rent_contract"
	"rent-manager-backend/internal/core/port"
)

type CreateRentalContractUseCase struct {
	rentalContractRepository   port.RentalContractRepository
	rentalContractCreatePolicy *policyRentContract.RentalContractCreatePolicy
}

func NewCreateRentalContractUseCase(
	rentalContractRepository port.RentalContractRepository,
	rentalContractCreatePolicy *policyRentContract.RentalContractCreatePolicy,
) *CreateRentalContractUseCase {
	return &CreateRentalContractUseCase{
		rentalContractRepository:   rentalContractRepository,
		rentalContractCreatePolicy: rentalContractCreatePolicy,
	}
}

func (uc *CreateRentalContractUseCase) CreateRentalContract(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) (*domain.RentalContract, error) {

	rentalContract.TotalPayments = shared.CalculateTotalPayments(&rentalContract.StartDate, &rentalContract.EndDate)
	if err := uc.rentalContractCreatePolicy.Execute(ctx, rentalContract); err != nil {
		slog.Error("[CreateRentalContractUseCase] policy validation failed", "err", err)
		return nil, err
	}

	createdRentalContract, err := uc.rentalContractRepository.CreateRentalContract(
		ctx,
		rentalContract,
	)
	if err != nil {
		slog.Error("[CreateRentalContractUseCase] repository create failed", "err", err)
		return nil, err
	}

	return createdRentalContract, nil
}
