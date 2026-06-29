package http

import "rent-manager-backend/internal/core/domain"

type createPropertyRequest struct {
	OwnerID  int64 `json:"owner_id" binding:"required,min=1"`
	TypeID   int64 `json:"type_id" binding:"required,min=1"`
	StatusID int64 `json:"status_id" binding:"required,min=1"`

	CountryID int64 `json:"country_id" binding:"required,min=1"`
	StateID   int64 `json:"state_id" binding:"required,min=1"`

	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`

	Street       string `json:"street" binding:"required"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
	City         string `json:"city" binding:"required"`
	PostalCode   string `json:"postal_code"`
}

type updatePropertyRequest struct {
	OwnerID  int64 `json:"owner_id" binding:"required,min=1"`
	TypeID   int64 `json:"type_id" binding:"required,min=1"`
	StatusID int64 `json:"status_id" binding:"required,min=1"`

	CountryID int64 `json:"country_id" binding:"required,min=1"`
	StateID   int64 `json:"state_id" binding:"required,min=1"`

	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`

	Street       string `json:"street" binding:"required"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
	City         string `json:"city" binding:"required"`
	PostalCode   string `json:"postal_code"`
}

type listPropertiesRequest struct {
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

type searchPropertiesRequest struct {
	Value string `form:"value" binding:"required"`
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newPropertyFromCreateRequest(req createPropertyRequest) *domain.Property {
	return &domain.Property{
		OwnerID:      req.OwnerID,
		TypeID:       req.TypeID,
		StatusID:     req.StatusID,
		CountryID:    req.CountryID,
		StateID:      req.StateID,
		Code:         req.Code,
		Title:        req.Title,
		Description:  req.Description,
		Street:       req.Street,
		StreetNumber: req.StreetNumber,
		Floor:        req.Floor,
		Apartment:    req.Apartment,
		City:         req.City,
		PostalCode:   req.PostalCode,
	}
}

func newPropertyFromUpdateRequest(
	id int64,
	req updatePropertyRequest,
) *domain.Property {
	return &domain.Property{
		ID:           id,
		OwnerID:      req.OwnerID,
		TypeID:       req.TypeID,
		StatusID:     req.StatusID,
		CountryID:    req.CountryID,
		StateID:      req.StateID,
		Code:         req.Code,
		Title:        req.Title,
		Description:  req.Description,
		Street:       req.Street,
		StreetNumber: req.StreetNumber,
		Floor:        req.Floor,
		Apartment:    req.Apartment,
		City:         req.City,
		PostalCode:   req.PostalCode,
	}
}
