package rental_contract

import (
	"context"
	"log/slog"
	"rent-manager-backend/internal/core/policy/shared"

	"rent-manager-backend/internal/core/domain"
	policyRentContract "rent-manager-backend/internal/core/policy/rent_contract"
	"rent-manager-backend/internal/core/port"
)

type UpdateRentalContractUseCase struct {
	rentalContractRepository   port.RentalContractRepository
	rentalContractUpdatePolicy *policyRentContract.RentalContractUpdatePolicy
}

func NewUpdateRentalContractUseCase(
	rentalContractRepository port.RentalContractRepository,
	rentalContractUpdatePolicy *policyRentContract.RentalContractUpdatePolicy,
) *UpdateRentalContractUseCase {
	return &UpdateRentalContractUseCase{
		rentalContractRepository:   rentalContractRepository,
		rentalContractUpdatePolicy: rentalContractUpdatePolicy,
	}
}

func (uc *UpdateRentalContractUseCase) UpdateRentalContract(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) (*domain.RentalContract, error) {
	rentalContract.TotalPayments = shared.CalculateTotalPayments(&rentalContract.StartDate, &rentalContract.EndDate)
	if err := uc.rentalContractUpdatePolicy.Execute(ctx, rentalContract); err != nil {
		slog.Error("[UpdateRentalContractUseCase] policy validation failed", "err", err)
		return nil, err
	}

	updatedRentalContract, err := uc.rentalContractRepository.UpdateRentalContract(
		ctx,
		rentalContract,
	)
	if err != nil {
		slog.Error("[UpdateRentalContractUseCase] repository update failed", "err", err)
		return nil, err
	}

	return updatedRentalContract, nil
}
