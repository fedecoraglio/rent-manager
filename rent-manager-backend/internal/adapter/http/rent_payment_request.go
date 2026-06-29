package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type createRentPaymentRequest struct {
	RentalContractID              int64   `json:"rental_contract_id" binding:"required,min=1"`
	Period                        string  `json:"period" binding:"required"`
	DueDate                       string  `json:"due_date" binding:"required"`
	PaymentDate                   string  `json:"payment_date"`
	BaseAmount                    float64 `json:"base_amount" binding:"required,gt=0"`
	SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
	AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
	SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
	AppliedInterestAmount         float64 `json:"applied_interest_amount"`
	TotalAmount                   float64 `json:"total_amount" binding:"required,gt=0"`
	PaidAmount                    float64 `json:"paid_amount"`
	IsPaid                        bool    `json:"is_paid"`
	Notes                         string  `json:"notes"`
}

type updateRentPaymentRequest struct {
	RentalContractID              int64   `json:"rental_contract_id" binding:"required,min=1"`
	Period                        string  `json:"period" binding:"required"`
	DueDate                       string  `json:"due_date" binding:"required"`
	PaymentDate                   string  `json:"payment_date"`
	BaseAmount                    float64 `json:"base_amount" binding:"required,gt=0"`
	SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
	AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
	SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
	AppliedInterestAmount         float64 `json:"applied_interest_amount"`
	TotalAmount                   float64 `json:"total_amount" binding:"required,gt=0"`
	PaidAmount                    float64 `json:"paid_amount"`
	IsPaid                        bool    `json:"is_paid"`
	Notes                         string  `json:"notes"`
}

type listRentPaymentsRequest struct {
	RentalContractID int64  `form:"rental_contract_id"`
	Page             uint64 `form:"page" binding:"required,min=1"`
	Limit            uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newRentPaymentFromCreateRequest(
	req createRentPaymentRequest,
) (*domain.RentPayment, error) {
	period, err := parsePeriod(req.Period)
	if err != nil {
		return nil, err
	}

	dueDate, err := parseDate(req.DueDate)
	if err != nil {
		return nil, err
	}

	paymentDate, err := parseOptionalDate(req.PaymentDate)
	if err != nil {
		return nil, err
	}

	if req.PaidAmount == 0 {
		req.PaidAmount = req.TotalAmount
	}
	return &domain.RentPayment{
		RentalContractID:              req.RentalContractID,
		Period:                        period,
		DueDate:                       dueDate,
		PaymentDate:                   paymentDate,
		BaseAmount:                    req.BaseAmount,
		SuggestedAdjustmentPercentage: req.SuggestedAdjustmentPercentage,
		AppliedAdjustmentPercentage:   req.AppliedAdjustmentPercentage,
		SuggestedInterestAmount:       req.SuggestedInterestAmount,
		AppliedInterestAmount:         req.AppliedInterestAmount,
		TotalAmount:                   req.TotalAmount,
		PaidAmount:                    req.PaidAmount,
		IsPaid:                        req.IsPaid,
		Notes:                         req.Notes,
	}, nil
}

func newRentPaymentFromUpdateRequest(
	id int64,
	req updateRentPaymentRequest,
) (*domain.RentPayment, error) {
	rentPayment, err := newRentPaymentFromCreateRequest(createRentPaymentRequest(req))
	if err != nil {
		return nil, err
	}

	rentPayment.ID = id

	return rentPayment, nil
}

func parsePeriod(value string) (time.Time, error) {
	return time.Parse("2006-01", value)
}

func parseOptionalDate(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	parsedDate, err := parseDate(value)
	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}
