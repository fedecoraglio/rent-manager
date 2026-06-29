package http

import "rent-manager-backend/internal/core/domain"

type rentPaymentSuggestionResponse struct {
	RentalContractID              int64   `json:"rental_contract_id"`
	Period                        string  `json:"period"`
	DueDate                       string  `json:"due_date"`
	PaymentDate                   string  `json:"payment_date"`
	BaseAmount                    float64 `json:"base_amount"`
	SuggestedAdjustmentPercentage float64 `json:"suggested_adjustment_percentage"`
	SuggestedAdjustmentAmount     float64 `json:"suggested_adjustment_amount"`
	SuggestedInterestAmount       float64 `json:"suggested_interest_amount"`
	SuggestedTotalAmount          float64 `json:"suggested_total_amount"`
}

func newRentPaymentSuggestionResponse(
	suggestion *domain.RentPaymentSuggestion,
) rentPaymentSuggestionResponse {
	return rentPaymentSuggestionResponse{
		RentalContractID:              suggestion.RentalContractID,
		Period:                        suggestion.Period.Format("2006-01"),
		DueDate:                       suggestion.DueDate.Format("2006-01-02"),
		PaymentDate:                   suggestion.PaymentDate.Format("2006-01-02"),
		BaseAmount:                    suggestion.BaseAmount,
		SuggestedAdjustmentPercentage: suggestion.SuggestedAdjustmentPercentage,
		SuggestedAdjustmentAmount:     suggestion.SuggestedAdjustmentAmount,
		SuggestedInterestAmount:       suggestion.SuggestedInterestAmount,
		SuggestedTotalAmount:          suggestion.SuggestedTotalAmount,
	}
}
