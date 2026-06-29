package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListPropertiesUseCase struct {
	propertyRepository port.PropertyRepository
}

func NewListPropertiesUseCase(propertyRepository port.PropertyRepository) *ListPropertiesUseCase {
	return &ListPropertiesUseCase{
		propertyRepository: propertyRepository,
	}
}

func (uc *ListPropertiesUseCase) ListProperties(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.Property, error) {
	properties, err := uc.propertyRepository.ListProperties(ctx, page, limit)
	if err != nil {
		slog.Error("[ListPropertiesUseCase] repository list failed", "err", err)
		return nil, err
	}

	return properties, nil
}
