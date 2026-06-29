package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type createRentalContractRequest struct {
	PropertyID int64 `json:"property_id" binding:"required,min=1"`
	TenantID   int64 `json:"tenant_id" binding:"required,min=1"`
	StatusID   int64 `json:"status_id" binding:"required,min=1"`

	InterestCalculationTypeID int64 `json:"interest_calculation_type_id" binding:"required,min=1"`
	AdjustmentTypeID          int64 `json:"adjustment_type_id" binding:"required,min=1"`

	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`

	MonthlyAmount float64 `json:"monthly_amount" binding:"required,gt=0"`
	DepositAmount float64 `json:"deposit_amount"`
	Currency      string  `json:"currency" binding:"required"`

	DueDay                    int64   `json:"due_day" binding:"required,min=1,max=31"`
	DailyInterestPercentage   float64 `json:"daily_interest_percentage"`
	AdjustmentFrequencyMonths int64   `json:"adjustment_frequency_months" binding:"required,min=1"`

	Notes string `json:"notes"`
}

type updateRentalContractRequest struct {
	PropertyID int64 `json:"property_id" binding:"required,min=1"`
	TenantID   int64 `json:"tenant_id" binding:"required,min=1"`
	StatusID   int64 `json:"status_id" binding:"required,min=1"`

	InterestCalculationTypeID int64 `json:"interest_calculation_type_id" binding:"required,min=1"`
	AdjustmentTypeID          int64 `json:"adjustment_type_id" binding:"required,min=1"`

	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`

	MonthlyAmount float64 `json:"monthly_amount" binding:"required,gt=0"`
	DepositAmount float64 `json:"deposit_amount"`
	Currency      string  `json:"currency" binding:"required"`

	DueDay                    int64   `json:"due_day" binding:"required,min=1,max=31"`
	DailyInterestPercentage   float64 `json:"daily_interest_percentage"`
	AdjustmentFrequencyMonths int64   `json:"adjustment_frequency_months" binding:"required,min=1"`

	Notes string `json:"notes"`
}

type listRentalContractsRequest struct {
	PropertyID int64  `form:"property_id"`
	Page       uint64 `form:"page" binding:"required,min=1"`
	Limit      uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newRentalContractFromCreateRequest(
	req createRentalContractRequest,
) (*domain.RentalContract, error) {
	startDate, err := parseDate(req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := parseDate(req.EndDate)
	if err != nil {
		return nil, err
	}

	return &domain.RentalContract{
		PropertyID:                req.PropertyID,
		TenantID:                  req.TenantID,
		StatusID:                  req.StatusID,
		InterestCalculationTypeID: req.InterestCalculationTypeID,
		AdjustmentTypeID:          req.AdjustmentTypeID,
		StartDate:                 startDate,
		EndDate:                   endDate,
		MonthlyAmount:             req.MonthlyAmount,
		DepositAmount:             req.DepositAmount,
		Currency:                  req.Currency,
		DueDay:                    req.DueDay,
		DailyInterestPercentage:   req.DailyInterestPercentage,
		AdjustmentFrequencyMonths: req.AdjustmentFrequencyMonths,
		Notes:                     req.Notes,
	}, nil
}

func newRentalContractFromUpdateRequest(
	id int64,
	req updateRentalContractRequest,
) (*domain.RentalContract, error) {
	rentalContract, err := newRentalContractFromCreateRequest(
		createRentalContractRequest(req),
	)
	if err != nil {
		return nil, err
	}

	rentalContract.ID = id

	return rentalContract, nil
}

func parseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}
