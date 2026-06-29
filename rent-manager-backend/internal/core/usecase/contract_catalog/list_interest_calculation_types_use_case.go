package contract_catalog

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListInterestCalculationTypesUseCase struct {
	contractCatalogRepository port.ContractCatalogRepository
}

func NewListInterestCalculationTypesUseCase(
	contractCatalogRepository port.ContractCatalogRepository,
) *ListInterestCalculationTypesUseCase {
	return &ListInterestCalculationTypesUseCase{
		contractCatalogRepository: contractCatalogRepository,
	}
}

func (uc *ListInterestCalculationTypesUseCase) ListInterestCalculationTypes(
	ctx context.Context,
) ([]domain.InterestCalculationType, error) {
	types, err := uc.contractCatalogRepository.ListInterestCalculationTypes(ctx)
	if err != nil {
		slog.Error("[ListInterestCalculationTypesUseCase] repository list interest calculation types failed", "err", err)
		return nil, err
	}

	return types, nil
}
