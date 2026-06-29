package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type TenantRepository interface {
	CreateTenant(ctx context.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	GetTenantByID(ctx context.Context, id int64) (*domain.Tenant, error)
	GetTenantByDocumentNumber(ctx context.Context, documentNumber string) (*domain.Tenant, error)
	ListTenants(ctx context.Context, page uint64, limit uint64) ([]domain.Tenant, error)
	SearchTenants(ctx context.Context, value string, page uint64, limit uint64) ([]domain.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *domain.Tenant) (*domain.Tenant, error)
}
