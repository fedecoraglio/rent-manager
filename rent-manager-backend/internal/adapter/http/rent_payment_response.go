package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type rentPaymentResponse struct {
	ID                            int64     `json:"id"`
	RentalContractID              int64     `json:"rental_contract_id"`
	Period                        string    `json:"period"`
	DueDate                       string    `json:"due_date"`
	PaymentDate                   *string   `json:"payment_date"`
	BaseAmount                    float64   `json:"base_amount"`
	SuggestedAdjustmentPercentage float64   `json:"suggested_adjustment_percentage"`
	AppliedAdjustmentPercentage   float64   `json:"applied_adjustment_percentage"`
	SuggestedInterestAmount       float64   `json:"suggested_interest_amount"`
	AppliedInterestAmount         float64   `json:"applied_interest_amount"`
	TotalAmount                   float64   `json:"total_amount"`
	PaidAmount                    float64   `json:"paid_amount"`
	IsPaid                        bool      `json:"is_paid"`
	Notes                         string    `json:"notes"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
}

func newRentPaymentResponse(
	rentPayment *domain.RentPayment,
) rentPaymentResponse {
	var paymentDate *string

	if rentPayment.PaymentDate != nil {
		formatted := rentPayment.PaymentDate.Format("2006-01-02")
		paymentDate = &formatted
	}

	return rentPaymentResponse{
		ID:                            rentPayment.ID,
		RentalContractID:              rentPayment.RentalContractID,
		Period:                        rentPayment.Period.Format("2006-01"),
		DueDate:                       rentPayment.DueDate.Format("2006-01-02"),
		PaymentDate:                   paymentDate,
		BaseAmount:                    rentPayment.BaseAmount,
		SuggestedAdjustmentPercentage: rentPayment.SuggestedAdjustmentPercentage,
		AppliedAdjustmentPercentage:   rentPayment.AppliedAdjustmentPercentage,
		SuggestedInterestAmount:       rentPayment.SuggestedInterestAmount,
		AppliedInterestAmount:         rentPayment.AppliedInterestAmount,
		TotalAmount:                   rentPayment.TotalAmount,
		PaidAmount:                    rentPayment.PaidAmount,
		IsPaid:                        rentPayment.IsPaid,
		Notes:                         rentPayment.Notes,
		CreatedAt:                     rentPayment.CreatedAt,
		UpdatedAt:                     rentPayment.UpdatedAt,
	}
}

func newRentPaymentsResponse(
	rentPayments []domain.RentPayment,
) []rentPaymentResponse {
	response := make([]rentPaymentResponse, 0, len(rentPayments))

	for _, rentPayment := range rentPayments {
		rentPaymentCopy := rentPayment
		response = append(response, newRentPaymentResponse(&rentPaymentCopy))
	}

	return response
}
