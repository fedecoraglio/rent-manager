package http

import "rent-manager-backend/internal/core/domain"

type rentalContractSummaryResponse struct {
	RentalContractID     int64   `json:"rental_contract_id"`
	TotalPayments        int64   `json:"total_payments"`
	PaidPayments         int64   `json:"paid_payments"`
	RemainingPayments    int64   `json:"remaining_payments"`
	CurrentAmount        float64 `json:"current_amount"`
	NextSuggestedAmount  float64 `json:"next_suggested_amount"`
	NextPendingPeriod    *string `json:"next_pending_period"`
	NextAdjustmentPeriod *string `json:"next_adjustment_period"`
}

func newRentalContractSummaryResponse(
	summary *domain.RentalContractSummary,
) rentalContractSummaryResponse {
	var nextPendingPeriod *string
	var nextAdjustmentPeriod *string

	if summary.NextPendingPeriod != nil {
		formatted := summary.NextPendingPeriod.Format("2006-01")
		nextPendingPeriod = &formatted
	}

	if summary.NextAdjustmentPeriod != nil {
		formatted := summary.NextAdjustmentPeriod.Format("2006-01")
		nextAdjustmentPeriod = &formatted
	}

	return rentalContractSummaryResponse{
		RentalContractID:     summary.RentalContractID,
		TotalPayments:        summary.TotalPayments,
		PaidPayments:         summary.PaidPayments,
		RemainingPayments:    summary.RemainingPayments,
		CurrentAmount:        summary.CurrentAmount,
		NextSuggestedAmount:  summary.NextSuggestedAmount,
		NextPendingPeriod:    nextPendingPeriod,
		NextAdjustmentPeriod: nextAdjustmentPeriod,
	}
}
