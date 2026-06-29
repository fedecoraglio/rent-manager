package policy

import (
	"context"
	"errors"
	"rent-manager-backend/internal/core/policy/shared"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type TenantUpdatePolicy struct {
	tenantRepository port.TenantRepository
}

func NewTenantUpdatePolicy(tenantRepository port.TenantRepository) *TenantUpdatePolicy {
	return &TenantUpdatePolicy{
		tenantRepository: tenantRepository,
	}
}

func (tenantPolicy *TenantUpdatePolicy) Execute(ctx context.Context, tenant *domain.Tenant) error {
	if tenant == nil {
		return domain.ErrTenantNil
	}

	if tenant.ID <= 0 {
		return domain.ErrTenantNotFound
	}

	tenant.Name = strings.TrimSpace(tenant.Name)
	tenant.Email = strings.TrimSpace(strings.ToLower(tenant.Email))
	tenant.Phone = strings.TrimSpace(tenant.Phone)
	tenant.DocumentNumber = strings.TrimSpace(tenant.DocumentNumber)
	tenant.City = strings.TrimSpace(tenant.City)
	tenant.Street = strings.TrimSpace(tenant.Street)
	tenant.StreetNumber = strings.TrimSpace(tenant.StreetNumber)
	tenant.Floor = strings.TrimSpace(tenant.Floor)
	tenant.Apartment = strings.TrimSpace(tenant.Apartment)
	tenant.PostalCode = strings.TrimSpace(tenant.PostalCode)

	if tenant.Name == "" {
		return domain.ErrTenantNameEmpty
	}

	if tenant.DocumentNumber == "" {
		return domain.ErrTenantDocumentNumberEmpty
	}

	if tenant.Email != "" && !shared.IsValidEmail(tenant.Email) {
		return domain.ErrTenantInvalidEmail
	}

	existingTenant, err := tenantPolicy.tenantRepository.GetTenantByDocumentNumber(
		ctx,
		tenant.DocumentNumber,
	)
	if err != nil {
		var appErr *domain.AppError

		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return domain.WrapAppError(
			domain.ErrCodeTenantDocumentLookupError,
			"error getting tenant by document number",
			err,
		)
	}

	if existingTenant != nil && existingTenant.ID != tenant.ID {
		return domain.ErrTenantAlreadyExists
	}

	return nil
}
