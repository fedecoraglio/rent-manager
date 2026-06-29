package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	propertyPolicy "rent-manager-backend/internal/core/policy/property"
	"rent-manager-backend/internal/core/port"
)

type CreatePropertyUseCase struct {
	propertyRepository   port.PropertyRepository
	propertyCreatePolicy *propertyPolicy.PropertyCreatePolicy
}

func NewCreatePropertyUseCase(
	propertyRepository port.PropertyRepository,
	propertyCreatePolicy *propertyPolicy.PropertyCreatePolicy,
) *CreatePropertyUseCase {
	return &CreatePropertyUseCase{
		propertyRepository:   propertyRepository,
		propertyCreatePolicy: propertyCreatePolicy,
	}
}

func (uc *CreatePropertyUseCase) CreateProperty(
	ctx context.Context,
	property *domain.Property,
) (*domain.Property, error) {
	if err := uc.propertyCreatePolicy.Execute(ctx, property); err != nil {
		slog.Error("[CreatePropertyUseCase] policy validation failed", "err", err)
		return nil, err
	}

	createdProperty, err := uc.propertyRepository.CreateProperty(ctx, property)
	if err != nil {
		slog.Error("[CreatePropertyUseCase] repository create failed", "err", err)
		return nil, err
	}

	return createdProperty, nil
}
