package property

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
	"rent-manager-backend/internal/core/usecase/rent_payment"
)

type ListPropertiesSummariesUseCase struct {
	rentalContractRepository        port.RentalContractRepository
	getRentalContractSummaryUseCase *rent_payment.GetRentalContractSummaryUseCase
}

func NewListPropertiesSummariesUseCase(
	rentalContractRepository port.RentalContractRepository,
	getRentalContractSummaryUseCase *rent_payment.GetRentalContractSummaryUseCase,
) *ListPropertiesSummariesUseCase {
	return &ListPropertiesSummariesUseCase{
		rentalContractRepository:        rentalContractRepository,
		getRentalContractSummaryUseCase: getRentalContractSummaryUseCase,
	}
}

func (uc *ListPropertiesSummariesUseCase) Execute(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.PropertySummary, error) {
	contracts, err := uc.rentalContractRepository.ListActiveRentalContracts(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	items := make([]domain.PropertySummary, 0, len(contracts))

	for _, contract := range contracts {
		summary, err := uc.getRentalContractSummaryUseCase.GetRentalContractSummary(ctx, contract.ID)
		if err != nil {
			return nil, err
		}

		item := domain.PropertySummary{
			PropertyID:            contract.PropertyID,
			RentalContractSummary: summary,
		}

		if contract.Property != nil {
			item.PropertyTitle = contract.Property.Title
		}

		items = append(items, item)
	}

	return items, nil
}
