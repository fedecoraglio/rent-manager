package location

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListCountriesUseCase struct {
	locationRepository port.LocationRepository
}

func NewListCountriesUseCase(
	locationRepository port.LocationRepository,
) *ListCountriesUseCase {
	return &ListCountriesUseCase{
		locationRepository: locationRepository,
	}
}

func (uc *ListCountriesUseCase) ListCountries(
	ctx context.Context,
) ([]domain.Country, error) {
	countries, err := uc.locationRepository.ListCountries(ctx)
	if err != nil {
		slog.Error("[ListCountriesUseCase] repository list countries failed", "err", err)
		return nil, err
	}

	return countries, nil
}
