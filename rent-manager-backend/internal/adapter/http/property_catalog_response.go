package http

import "rent-manager-backend/internal/core/domain"

type propertyTypeResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type propertyStatusResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func newPropertyTypeResponse(propertyType domain.PropertyType) propertyTypeResponse {
	return propertyTypeResponse{
		ID:   propertyType.ID,
		Code: propertyType.Code,
		Name: propertyType.Name,
	}
}

func newPropertyTypesResponse(propertyTypes []domain.PropertyType) []propertyTypeResponse {
	response := make([]propertyTypeResponse, 0, len(propertyTypes))

	for _, propertyType := range propertyTypes {
		response = append(response, newPropertyTypeResponse(propertyType))
	}

	return response
}

func newPropertyStatusResponse(propertyStatus domain.PropertyStatus) propertyStatusResponse {
	return propertyStatusResponse{
		ID:   propertyStatus.ID,
		Code: propertyStatus.Code,
		Name: propertyStatus.Name,
	}
}

func newPropertyStatusesResponse(propertyStatuses []domain.PropertyStatus) []propertyStatusResponse {
	response := make([]propertyStatusResponse, 0, len(propertyStatuses))

	for _, propertyStatus := range propertyStatuses {
		response = append(response, newPropertyStatusResponse(propertyStatus))
	}

	return response
}
