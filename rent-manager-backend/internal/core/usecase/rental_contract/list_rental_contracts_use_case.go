package rental_contract

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListRentalContractsUseCase struct {
	rentalContractRepository port.RentalContractRepository
}

func NewListRentalContractsUseCase(
	rentalContractRepository port.RentalContractRepository,
) *ListRentalContractsUseCase {
	return &ListRentalContractsUseCase{
		rentalContractRepository: rentalContractRepository,
	}
}

func (uc *ListRentalContractsUseCase) ListRentalContracts(
	ctx context.Context,
	propertyID int64,
	page uint64,
	limit uint64,
) ([]domain.RentalContract, error) {
	rentalContracts, err := uc.rentalContractRepository.ListRentalContracts(
		ctx,
		propertyID,
		page,
		limit,
	)
	if err != nil {
		slog.Error("[ListRentalContractsUseCase] repository list failed", "err", err)
		return nil, err
	}

	return rentalContracts, nil
}
