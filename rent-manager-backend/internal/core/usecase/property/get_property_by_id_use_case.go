package property

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetPropertyByIDUseCase struct {
	propertyRepository port.PropertyRepository
}

func NewGetPropertyByIDUseCase(propertyRepository port.PropertyRepository) *GetPropertyByIDUseCase {
	return &GetPropertyByIDUseCase{
		propertyRepository: propertyRepository,
	}
}

func (uc *GetPropertyByIDUseCase) GetPropertyByID(ctx context.Context, id int64) (*domain.Property, error) {
	property, err := uc.propertyRepository.GetPropertyByID(ctx, id)
	if err != nil {
		slog.Error("[GetPropertyByIDUseCase] repository get by id failed", "err", err)
		return nil, err
	}

	return property, nil
}
