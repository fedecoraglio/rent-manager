package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type SearchPropertiesUseCase struct {
	propertyRepository port.PropertyRepository
}

func NewSearchPropertiesUseCase(propertyRepository port.PropertyRepository) *SearchPropertiesUseCase {
	return &SearchPropertiesUseCase{
		propertyRepository: propertyRepository,
	}
}

func (uc *SearchPropertiesUseCase) SearchProperties(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Property, error) {
	properties, err := uc.propertyRepository.SearchProperties(ctx, value, page, limit)
	if err != nil {
		slog.Error("[SearchPropertiesUseCase] repository search failed", "err", err)
		return nil, err
	}

	return properties, nil
}
