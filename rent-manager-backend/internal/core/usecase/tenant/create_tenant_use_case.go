package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/tenant"
	"rent-manager-backend/internal/core/port"
)

type CreateTenantUseCase struct {
	tenantRepository   port.TenantRepository
	tenantCreatePolicy *policy.TenantCreatePolicy
}

func NewCreateTenantUseCase(
	tenantRepository port.TenantRepository,
	tenantCreatePolicy *policy.TenantCreatePolicy,
) *CreateTenantUseCase {
	return &CreateTenantUseCase{
		tenantRepository:   tenantRepository,
		tenantCreatePolicy: tenantCreatePolicy,
	}
}

func (uc *CreateTenantUseCase) CreateTenant(
	ctx context.Context,
	tenant *domain.Tenant,
) (*domain.Tenant, error) {
	if err := uc.tenantCreatePolicy.Execute(ctx, tenant); err != nil {
		slog.Error("[CreateTenantUseCase] policy validation failed", "err", err)
		return nil, err
	}

	createdTenant, err := uc.tenantRepository.CreateTenant(ctx, tenant)
	if err != nil {
		slog.Error("[CreateTenantUseCase] repository create failed", "err", err)
		return nil, err
	}

	return createdTenant, nil
}
