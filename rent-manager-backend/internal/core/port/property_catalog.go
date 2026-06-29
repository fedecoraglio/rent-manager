package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type PropertyCatalogRepository interface {
	ListPropertyTypes(ctx context.Context) ([]domain.PropertyType, error)
	ListPropertyStatuses(ctx context.Context) ([]domain.PropertyStatus, error)
}
