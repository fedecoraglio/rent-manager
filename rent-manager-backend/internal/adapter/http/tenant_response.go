package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type tenantResponse struct {
	ID             int64  `json:"id"`
	CountryID      int64  `json:"country_id"`
	StateID        int64  `json:"state_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DocumentNumber string `json:"document_number"`

	City         string `json:"city"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
	PostalCode   string `json:"postal_code"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newTenantResponse(tenant *domain.Tenant) tenantResponse {
	return tenantResponse{
		ID:             tenant.ID,
		CountryID:      tenant.CountryID,
		StateID:        tenant.StateID,
		Name:           tenant.Name,
		Email:          tenant.Email,
		Phone:          tenant.Phone,
		DocumentNumber: tenant.DocumentNumber,
		City:           tenant.City,
		Street:         tenant.Street,
		StreetNumber:   tenant.StreetNumber,
		Floor:          tenant.Floor,
		Apartment:      tenant.Apartment,
		PostalCode:     tenant.PostalCode,
		CreatedAt:      tenant.CreatedAt,
		UpdatedAt:      tenant.UpdatedAt,
	}
}

func newTenantsResponse(tenants []domain.Tenant) []tenantResponse {
	response := make([]tenantResponse, 0, len(tenants))

	for _, tenant := range tenants {
		tenantCopy := tenant
		response = append(response, newTenantResponse(&tenantCopy))
	}

	return response
}
