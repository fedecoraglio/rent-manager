package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type OwnerRepository interface {
	CreateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error)
	GetOwnerByID(ctx context.Context, id int64) (*domain.Owner, error)
	GetOwnerByDocumentNumber(ctx context.Context, documentNumber string) (*domain.Owner, error)
	ListOwners(ctx context.Context, page uint64, limit uint64) ([]domain.Owner, error)
	SearchOwners(ctx context.Context, value string, page uint64, limit uint64) ([]domain.Owner, error)
	UpdateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error)
}
