package http

import "rent-manager-backend/internal/core/domain"

type createTenantRequest struct {
	CountryID      int64  `json:"country_id"`
	StateID        int64  `json:"state_id"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"omitempty,email"`
	Phone          string `json:"phone"`
	DocumentNumber string `json:"document_number" binding:"required"`

	City         string `json:"city"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
	PostalCode   string `json:"postal_code"`
}

type updateTenantRequest struct {
	CountryID      int64  `json:"country_id"`
	StateID        int64  `json:"state_id"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"omitempty,email"`
	Phone          string `json:"phone"`
	DocumentNumber string `json:"document_number" binding:"required"`

	City         string `json:"city"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
	PostalCode   string `json:"postal_code"`
}

type listTenantsRequest struct {
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

type searchTenantsRequest struct {
	Value string `form:"value" binding:"required"`
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newTenantFromCreateRequest(req createTenantRequest) *domain.Tenant {
	return &domain.Tenant{
		CountryID:      req.CountryID,
		StateID:        req.StateID,
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		DocumentNumber: req.DocumentNumber,

		City:         req.City,
		Street:       req.Street,
		StreetNumber: req.StreetNumber,
		Floor:        req.Floor,
		Apartment:    req.Apartment,
		PostalCode:   req.PostalCode,
	}
}

func newTenantFromUpdateRequest(id int64, req updateTenantRequest) *domain.Tenant {
	return &domain.Tenant{
		ID:             id,
		CountryID:      req.CountryID,
		StateID:        req.StateID,
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		DocumentNumber: req.DocumentNumber,

		City:         req.City,
		Street:       req.Street,
		StreetNumber: req.StreetNumber,
		Floor:        req.Floor,
		Apartment:    req.Apartment,
		PostalCode:   req.PostalCode,
	}
}
