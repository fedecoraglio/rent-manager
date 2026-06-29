package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type propertyResponse struct {
	ID           int64     `json:"id"`
	OwnerID      int64     `json:"owner_id"`
	TypeID       int64     `json:"type_id"`
	StatusID     int64     `json:"status_id"`
	CountryID    int64     `json:"country_id"`
	StateID      int64     `json:"state_id"`
	Code         string    `json:"code"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Street       string    `json:"street"`
	StreetNumber string    `json:"street_number"`
	Floor        string    `json:"floor"`
	Apartment    string    `json:"apartment"`
	City         string    `json:"city"`
	PostalCode   string    `json:"postal_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type propertySummaryResponse struct {
	PropertyID            int64                         `json:"id"`
	PropertyTitle         string                        `json:"title"`
	RentalContractSummary rentalContractSummaryResponse `json:"summary"`
}

func newPropertyResponse(property *domain.Property) propertyResponse {
	return propertyResponse{
		ID:           property.ID,
		OwnerID:      property.OwnerID,
		TypeID:       property.TypeID,
		StatusID:     property.StatusID,
		CountryID:    property.CountryID,
		StateID:      property.StateID,
		Code:         property.Code,
		Title:        property.Title,
		Description:  property.Description,
		Street:       property.Street,
		StreetNumber: property.StreetNumber,
		Floor:        property.Floor,
		Apartment:    property.Apartment,
		City:         property.City,
		PostalCode:   property.PostalCode,
		CreatedAt:    property.CreatedAt,
		UpdatedAt:    property.UpdatedAt,
	}
}

func newPropertySummaryResponse(property *domain.PropertySummary) propertySummaryResponse {
	return propertySummaryResponse{
		PropertyID:            property.PropertyID,
		PropertyTitle:         property.PropertyTitle,
		RentalContractSummary: newRentalContractSummaryResponse(property.RentalContractSummary),
	}
}

func newPropertiesResponse(properties []domain.Property) []propertyResponse {
	response := make([]propertyResponse, 0, len(properties))

	for _, property := range properties {
		propertyCopy := property
		response = append(response, newPropertyResponse(&propertyCopy))
	}

	return response
}

func newPropertiesSummaryResponse(properties []domain.PropertySummary) []propertySummaryResponse {
	response := make([]propertySummaryResponse, 0, len(properties))
	for _, propertySummary := range properties {
		propertySummaryCopy := newPropertySummaryResponse(&propertySummary)
		response = append(response, propertySummaryCopy)
	}
	return response
}
