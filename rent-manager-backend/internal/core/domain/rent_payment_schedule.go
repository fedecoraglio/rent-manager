package domain

import "time"

type RentPaymentScheduleItem struct {
	RentalContractID int64
	RentPaymentID    int64

	Period  time.Time
	DueDate time.Time

	BaseAmount float64

	SuggestedAdjustmentPercentage float64
	SuggestedInterestAmount       float64
	SuggestedTotalAmount          float64

	AppliedAdjustmentPercentage float64
	AppliedInterestAmount       float64
	TotalAmount                 float64
	PaidAmount                  float64
	PaymentDate                 *time.Time

	IsPaid bool
}
