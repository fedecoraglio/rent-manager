package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetPropertyByCodeUseCase struct {
	propertyRepository port.PropertyRepository
}

func NewGetPropertyByCodeUseCase(propertyRepository port.PropertyRepository) *GetPropertyByCodeUseCase {
	return &GetPropertyByCodeUseCase{
		propertyRepository: propertyRepository,
	}
}

func (uc *GetPropertyByCodeUseCase) GetPropertyByCode(ctx context.Context, code string) (*domain.Property, error) {
	property, err := uc.propertyRepository.GetPropertyByCode(ctx, code)
	if err != nil {
		slog.Error("[GetPropertyByCodeUseCase] repository get by code failed", "err", err)
		return nil, err
	}

	return property, nil
}
