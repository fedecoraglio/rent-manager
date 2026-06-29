package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type SearchTenantsUseCase struct {
	tenantRepository port.TenantRepository
}

func NewSearchTenantsUseCase(tenantRepository port.TenantRepository) *SearchTenantsUseCase {
	return &SearchTenantsUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *SearchTenantsUseCase) SearchTenants(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Tenant, error) {
	tenants, err := uc.tenantRepository.SearchTenants(ctx, value, page, limit)
	if err != nil {
		slog.Error("[SearchTenantsUseCase] repository search failed", "err", err)
		return nil, err
	}

	return tenants, nil
}
