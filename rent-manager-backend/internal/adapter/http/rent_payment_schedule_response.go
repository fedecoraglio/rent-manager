package http

import (
	"rent-manager-backend/internal/core/domain"
)

type rentPaymentScheduleItemResponse struct {
	RentalContractID              int64   `json:"rental_contract_id"`
	RentPaymentID                 int64   `json:"rent_payment_id"`
	Period                        string  `json:"period"`
	DueDate                       string  `json:"due_date"`
	BaseAmount                    float64 `json:"base_amount"`
	SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
	SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
	SuggestedTotalAmount          float64 `json:"suggested_total_amount"`
	AppliedAdjustmentPercentage   float64 `json:"applied_adjustment_percentage"`
	AppliedInterestAmount         float64 `json:"applied_interest_amount"`
	TotalAmount                   float64 `json:"total_amount"`
	PaidAmount                    float64 `json:"paid_amount"`
	PaymentDate                   *string `json:"payment_date"`
	IsPaid                        bool    `json:"is_paid"`
}

func newRentPaymentScheduleItemResponse(
	item domain.RentPaymentScheduleItem,
) rentPaymentScheduleItemResponse {
	var paymentDate *string

	if item.PaymentDate != nil {
		formatted := item.PaymentDate.Format("2006-01-02")
		paymentDate = &formatted
	}

	return rentPaymentScheduleItemResponse{
		RentalContractID:              item.RentalContractID,
		RentPaymentID:                 item.RentPaymentID,
		Period:                        item.Period.Format("2006-01"),
		DueDate:                       item.DueDate.Format("2006-01-02"),
		BaseAmount:                    item.BaseAmount,
		SuggestedAdjustmentPercentage: item.SuggestedAdjustmentPercentage,
		SuggestedInterestAmount:       item.SuggestedInterestAmount,
		SuggestedTotalAmount:          item.SuggestedTotalAmount,
		AppliedAdjustmentPercentage:   item.AppliedAdjustmentPercentage,
		AppliedInterestAmount:         item.AppliedInterestAmount,
		TotalAmount:                   item.TotalAmount,
		PaidAmount:                    item.PaidAmount,
		PaymentDate:                   paymentDate,
		IsPaid:                        item.IsPaid,
	}
}

func newRentPaymentScheduleResponse(
	items []domain.RentPaymentScheduleItem,
) []rentPaymentScheduleItemResponse {
	response := make([]rentPaymentScheduleItemResponse, 0, len(items))

	for _, item := range items {
		response = append(response, newRentPaymentScheduleItemResponse(item))
	}

	return response
}
