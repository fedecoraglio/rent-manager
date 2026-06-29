package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type ContractCatalogRepository interface {
	ListContractStatuses(ctx context.Context) ([]domain.ContractStatus, error)
	ListInterestCalculationTypes(ctx context.Context) ([]domain.InterestCalculationType, error)
	ListRentAdjustmentTypes(ctx context.Context) ([]domain.RentAdjustmentType, error)
}
