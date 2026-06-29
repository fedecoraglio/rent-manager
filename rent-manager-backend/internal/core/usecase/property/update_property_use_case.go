package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	propertyPolicy "rent-manager-backend/internal/core/policy/property"
	"rent-manager-backend/internal/core/port"
)

type UpdatePropertyUseCase struct {
	propertyRepository   port.PropertyRepository
	propertyUpdatePolicy *propertyPolicy.PropertyUpdatePolicy
}

func NewUpdatePropertyUseCase(
	propertyRepository port.PropertyRepository,
	propertyUpdatePolicy *propertyPolicy.PropertyUpdatePolicy,
) *UpdatePropertyUseCase {
	return &UpdatePropertyUseCase{
		propertyRepository:   propertyRepository,
		propertyUpdatePolicy: propertyUpdatePolicy,
	}
}

func (uc *UpdatePropertyUseCase) UpdateProperty(
	ctx context.Context,
	property *domain.Property,
) (*domain.Property, error) {
	if err := uc.propertyUpdatePolicy.Execute(ctx, property); err != nil {
		slog.Error("[UpdatePropertyUseCase] policy validation failed", "err", err)
		return nil, err
	}

	updatedProperty, err := uc.propertyRepository.UpdateProperty(ctx, property)
	if err != nil {
		slog.Error("[UpdatePropertyUseCase] repository update failed", "err", err)
		return nil, err
	}

	return updatedProperty, nil
}
