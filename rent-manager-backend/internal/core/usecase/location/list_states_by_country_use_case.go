package location

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListStatesByCountryUseCase struct {
	locationRepository port.LocationRepository
}

func NewListStatesByCountryUseCase(
	locationRepository port.LocationRepository,
) *ListStatesByCountryUseCase {
	return &ListStatesByCountryUseCase{
		locationRepository: locationRepository,
	}
}

func (uc *ListStatesByCountryUseCase) ListStatesByCountry(
	ctx context.Context,
	countryID int64,
) ([]domain.State, error) {
	states, err := uc.locationRepository.ListStatesByCountry(
		ctx,
		countryID,
	)
	if err != nil {
		slog.Error("[ListStatesByCountryUseCase] repository list states failed", "err", err)
		return nil, err
	}

	return states, nil
}
