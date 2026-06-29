package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/tenant"
	"rent-manager-backend/internal/core/port"
)

type UpdateTenantUseCase struct {
	tenantRepository   port.TenantRepository
	tenantUpdatePolicy *policy.TenantUpdatePolicy
}

func NewUpdateTenantUseCase(
	tenantRepository port.TenantRepository,
	tenantUpdatePolicy *policy.TenantUpdatePolicy,
) *UpdateTenantUseCase {
	return &UpdateTenantUseCase{
		tenantRepository:   tenantRepository,
		tenantUpdatePolicy: tenantUpdatePolicy,
	}
}

func (uc *UpdateTenantUseCase) UpdateTenant(
	ctx context.Context,
	tenant *domain.Tenant,
) (*domain.Tenant, error) {
	if err := uc.tenantUpdatePolicy.Execute(ctx, tenant); err != nil {
		slog.Error("[UpdateTenantUseCase] policy validation failed", "err", err)
		return nil, err
	}

	updatedTenant, err := uc.tenantRepository.UpdateTenant(ctx, tenant)
	if err != nil {
		slog.Error("[UpdateTenantUseCase] repository update failed", "err", err)
		return nil, err
	}

	return updatedTenant, nil
}
