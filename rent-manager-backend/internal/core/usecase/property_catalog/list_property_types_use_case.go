package property_catalog

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListPropertyTypesUseCase struct {
	propertyCatalogRepository port.PropertyCatalogRepository
}

func NewListPropertyTypesUseCase(
	propertyCatalogRepository port.PropertyCatalogRepository,
) *ListPropertyTypesUseCase {
	return &ListPropertyTypesUseCase{
		propertyCatalogRepository: propertyCatalogRepository,
	}
}

func (uc *ListPropertyTypesUseCase) ListPropertyTypes(
	ctx context.Context,
) ([]domain.PropertyType, error) {
	propertyTypes, err := uc.propertyCatalogRepository.ListPropertyTypes(ctx)
	if err != nil {
		slog.Error("[ListPropertyTypesUseCase] repository list property types failed", "err", err)
		return nil, err
	}

	return propertyTypes, nil
}
