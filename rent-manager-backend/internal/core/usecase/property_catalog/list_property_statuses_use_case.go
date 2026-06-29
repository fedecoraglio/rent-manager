package property_catalog

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListPropertyStatusesUseCase struct {
	propertyCatalogRepository port.PropertyCatalogRepository
}

func NewListPropertyStatusesUseCase(
	propertyCatalogRepository port.PropertyCatalogRepository,
) *ListPropertyStatusesUseCase {
	return &ListPropertyStatusesUseCase{
		propertyCatalogRepository: propertyCatalogRepository,
	}
}

func (uc *ListPropertyStatusesUseCase) ListPropertyStatuses(
	ctx context.Context,
) ([]domain.PropertyStatus, error) {
	propertyStatuses, err := uc.propertyCatalogRepository.ListPropertyStatuses(ctx)
	if err != nil {
		slog.Error("[ListPropertyStatusesUseCase] repository list property statuses failed", "err", err)
		return nil, err
	}

	return propertyStatuses, nil
}
