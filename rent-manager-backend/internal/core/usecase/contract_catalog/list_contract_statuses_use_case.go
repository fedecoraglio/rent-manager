package contract_catalog

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListContractStatusesUseCase struct {
	contractCatalogRepository port.ContractCatalogRepository
}

func NewListContractStatusesUseCase(
	contractCatalogRepository port.ContractCatalogRepository,
) *ListContractStatusesUseCase {
	return &ListContractStatusesUseCase{
		contractCatalogRepository: contractCatalogRepository,
	}
}

func (uc *ListContractStatusesUseCase) ListContractStatuses(
	ctx context.Context,
) ([]domain.ContractStatus, error) {
	statuses, err := uc.contractCatalogRepository.ListContractStatuses(ctx)
	if err != nil {
		slog.Error("[ListContractStatusesUseCase] repository list contract statuses failed", "err", err)
		return nil, err
	}

	return statuses, nil
}
