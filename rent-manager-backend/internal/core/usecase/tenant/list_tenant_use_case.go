package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListTenantUseCase struct {
	tenantRepository port.TenantRepository
}

func NewListTenantsUseCase(tenantRepository port.TenantRepository) *ListTenantUseCase {
	return &ListTenantUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *ListTenantUseCase) ListTenants(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.Tenant, error) {
	tenants, err := uc.tenantRepository.ListTenants(ctx, page, limit)
	if err != nil {
		slog.Error("[ListTenantUseCase] repository list failed", "err", err)
		return nil, err
	}

	return tenants, nil
}
