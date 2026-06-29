package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type PropertyRepository interface {
	CreateProperty(ctx context.Context, property *domain.Property) (*domain.Property, error)
	GetPropertyByID(ctx context.Context, id int64) (*domain.Property, error)
	GetPropertyByCode(ctx context.Context, code string) (*domain.Property, error)
	ListProperties(ctx context.Context, page uint64, limit uint64) ([]domain.Property, error)
	SearchProperties(ctx context.Context, value string, page uint64, limit uint64) ([]domain.Property, error)
	UpdateProperty(ctx context.Context, property *domain.Property) (*domain.Property, error)
}
