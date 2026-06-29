package port

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
)

type InflationIndexRepository interface {
	Create(ctx context.Context, inflationIndex *domain.InflationIndex) error
	Update(ctx context.Context, inflationIndex *domain.InflationIndex) error
	GetByID(ctx context.Context, id int64) (*domain.InflationIndex, error)
	GetByPeriod(ctx context.Context, period time.Time) (*domain.InflationIndex, error)
	List(ctx context.Context, page uint64, limit uint64) ([]domain.InflationIndex, error)
	ListByPeriodRange(ctx context.Context, from time.Time, to time.Time) ([]domain.InflationIndex, error)
	ListLatestBeforePeriod(ctx context.Context, period time.Time, limit uint64) ([]domain.InflationIndex, error)
}
