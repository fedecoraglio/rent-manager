package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type rentalContractResponse struct {
	ID                        int64            `json:"id"`
	PropertyID                int64            `json:"property_id"`
	TenantID                  int64            `json:"tenant_id"`
	StatusID                  int64            `json:"status_id"`
	InterestCalculationTypeID int64            `json:"interest_calculation_type_id"`
	AdjustmentTypeID          int64            `json:"adjustment_type_id"`
	StartDate                 string           `json:"start_date"`
	EndDate                   string           `json:"end_date"`
	TotalPayments             int64            `json:"total_payments" binding:"required,min=0"`
	MonthlyAmount             float64          `json:"monthly_amount"`
	DepositAmount             float64          `json:"deposit_amount"`
	Currency                  string           `json:"currency"`
	DueDay                    int64            `json:"due_day"`
	DailyInterestPercentage   float64          `json:"daily_interest_percentage"`
	AdjustmentFrequencyMonths int64            `json:"adjustment_frequency_months"`
	Property                  propertyResponse `json:"property"`
	Tenant                    tenantResponse   `json:"tenant"`
	Notes                     string           `json:"notes"`
	CreatedAt                 time.Time        `json:"created_at"`
	UpdatedAt                 time.Time        `json:"updated_at"`
}

func newRentalContractResponse(
	rentalContract *domain.RentalContract,
) rentalContractResponse {
	property := propertyResponse{}
	if rentalContract.Property != nil {
		property.ID = rentalContract.Property.ID
		property.Title = rentalContract.Property.Title
	}
	tenant := tenantResponse{}
	if rentalContract.Tenant != nil {
		tenant.ID = rentalContract.Tenant.ID
		tenant.Name = rentalContract.Tenant.Name
	}
	return rentalContractResponse{
		ID:                        rentalContract.ID,
		PropertyID:                rentalContract.PropertyID,
		TenantID:                  rentalContract.TenantID,
		StatusID:                  rentalContract.StatusID,
		InterestCalculationTypeID: rentalContract.InterestCalculationTypeID,
		AdjustmentTypeID:          rentalContract.AdjustmentTypeID,
		StartDate:                 formatDate(rentalContract.StartDate),
		EndDate:                   formatDate(rentalContract.EndDate),
		MonthlyAmount:             rentalContract.MonthlyAmount,
		DepositAmount:             rentalContract.DepositAmount,
		Currency:                  rentalContract.Currency,
		DueDay:                    rentalContract.DueDay,
		DailyInterestPercentage:   rentalContract.DailyInterestPercentage,
		AdjustmentFrequencyMonths: rentalContract.AdjustmentFrequencyMonths,
		Notes:                     rentalContract.Notes,
		Property:                  property,
		Tenant:                    tenant,
		CreatedAt:                 rentalContract.CreatedAt,
		UpdatedAt:                 rentalContract.UpdatedAt,
	}
}

func newRentalContractsResponse(
	rentalContracts []domain.RentalContract,
) []rentalContractResponse {
	response := make([]rentalContractResponse, 0, len(rentalContracts))

	for _, rentalContract := range rentalContracts {
		rentalContractCopy := rentalContract
		response = append(response, newRentalContractResponse(&rentalContractCopy))
	}

	return response
}

func formatDate(value time.Time) string {
	if value.IsZero() {
		return ""
	}

	return value.Format("2006-01-02")
}
