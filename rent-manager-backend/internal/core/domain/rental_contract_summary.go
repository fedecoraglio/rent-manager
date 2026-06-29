package domain

import "time"

type RentalContractSummary struct {
	RentalContractID     int64
	TotalPayments        int64
	PaidPayments         int64
	RemainingPayments    int64
	CurrentAmount        float64
	NextSuggestedAmount  float64
	NextPendingPeriod    *time.Time
	NextAdjustmentPeriod *time.Time
}
