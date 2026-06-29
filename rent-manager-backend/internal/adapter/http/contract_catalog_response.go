package http

import "rent-manager-backend/internal/core/domain"

type contractStatusResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type interestCalculationTypeResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type rentAdjustmentTypeResponse struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func newContractStatusResponse(status domain.ContractStatus) contractStatusResponse {
	return contractStatusResponse{
		ID:   status.ID,
		Code: status.Code,
		Name: status.Name,
	}
}

func newContractStatusesResponse(statuses []domain.ContractStatus) []contractStatusResponse {
	response := make([]contractStatusResponse, 0, len(statuses))

	for _, status := range statuses {
		response = append(response, newContractStatusResponse(status))
	}

	return response
}

func newInterestCalculationTypeResponse(
	interestType domain.InterestCalculationType,
) interestCalculationTypeResponse {
	return interestCalculationTypeResponse{
		ID:   interestType.ID,
		Code: interestType.Code,
		Name: interestType.Name,
	}
}

func newInterestCalculationTypesResponse(
	interestTypes []domain.InterestCalculationType,
) []interestCalculationTypeResponse {
	response := make([]interestCalculationTypeResponse, 0, len(interestTypes))

	for _, interestType := range interestTypes {
		response = append(response, newInterestCalculationTypeResponse(interestType))
	}

	return response
}

func newRentAdjustmentTypeResponse(
	adjustmentType domain.RentAdjustmentType,
) rentAdjustmentTypeResponse {
	return rentAdjustmentTypeResponse{
		ID:   adjustmentType.ID,
		Code: adjustmentType.Code,
		Name: adjustmentType.Name,
	}
}

func newRentAdjustmentTypesResponse(adjustmentTypes []domain.RentAdjustmentType) []rentAdjustmentTypeResponse {
	response := make([]rentAdjustmentTypeResponse, 0, len(adjustmentTypes))

	for _, adjustmentType := range adjustmentTypes {
		response = append(response, newRentAdjustmentTypeResponse(adjustmentType))
	}

	return response
}
