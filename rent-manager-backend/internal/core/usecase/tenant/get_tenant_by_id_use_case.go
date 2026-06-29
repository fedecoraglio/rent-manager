package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetTenantByIDUseCase struct {
	tenantRepository port.TenantRepository
}

func NewGetTenantByIDUseCase(tenantRepository port.TenantRepository) *GetTenantByIDUseCase {
	return &GetTenantByIDUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *GetTenantByIDUseCase) GetTenantByID(ctx context.Context, id int64) (*domain.Tenant, error) {
	tenant, err := uc.tenantRepository.GetTenantByID(ctx, id)
	if err != nil {
		slog.Error("[GetTenantByIDUseCase] repository get by id failed", "err", err)
		return nil, err
	}

	return tenant, nil
}
