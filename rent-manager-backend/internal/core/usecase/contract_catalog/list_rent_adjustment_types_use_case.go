package contract_catalog

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListRentAdjustmentTypesUseCase struct {
	contractCatalogRepository port.ContractCatalogRepository
}

func NewListRentAdjustmentTypesUseCase(
	contractCatalogRepository port.ContractCatalogRepository,
) *ListRentAdjustmentTypesUseCase {
	return &ListRentAdjustmentTypesUseCase{
		contractCatalogRepository: contractCatalogRepository,
	}
}

func (uc *ListRentAdjustmentTypesUseCase) ListRentAdjustmentTypes(
	ctx context.Context,
) ([]domain.RentAdjustmentType, error) {
	types, err := uc.contractCatalogRepository.ListRentAdjustmentTypes(ctx)
	if err != nil {
		slog.Error("[ListRentAdjustmentTypesUseCase] repository list rent adjustment types failed", "err", err)
		return nil, err
	}

	return types, nil
}
