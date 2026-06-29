package domain

import "time"

type RentPayment struct {
	ID                            int64
	RentalContractID              int64
	Period                        time.Time
	DueDate                       time.Time
	PaymentDate                   *time.Time
	BaseAmount                    float64
	SuggestedAdjustmentPercentage float64
	AppliedAdjustmentPercentage   float64
	SuggestedInterestAmount       float64
	AppliedInterestAmount         float64
	TotalAmount                   float64
	PaidAmount                    float64
	IsPaid                        bool
	Notes                         string
	RentalContract                *RentalContract
	CreatedAt                     time.Time
	UpdatedAt                     time.Time
}
