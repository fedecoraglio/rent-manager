package tenant

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetTenantByDocumentNumberUseCase struct {
	tenantRepository port.TenantRepository
}

func NewGetTenantByDocumentNumberUseCase(tenantRepository port.TenantRepository) *GetTenantByDocumentNumberUseCase {
	return &GetTenantByDocumentNumberUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *GetTenantByDocumentNumberUseCase) GetTenantByDocumentNumber(
	ctx context.Context,
	documentNumber string,
) (*domain.Tenant, error) {
	tenant, err := uc.tenantRepository.GetTenantByDocumentNumber(ctx, documentNumber)
	if err != nil {
		slog.Error("[GetTenantByDocumentNumberUseCase] repository get by document number failed", "err", err)
		return nil, err
	}

	return tenant, nil
}
