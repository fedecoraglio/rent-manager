package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type LocationRepository interface {
	ListCountries(ctx context.Context) ([]domain.Country, error)
	ListStatesByCountry(ctx context.Context, countryID int64) ([]domain.State, error)
}
