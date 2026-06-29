package domain

import "time"

type RentPaymentSuggestion struct {
	RentalContractID              int64
	Period                        time.Time
	DueDate                       time.Time
	PaymentDate                   time.Time
	BaseAmount                    float64
	SuggestedAdjustmentPercentage float64
	SuggestedAdjustmentAmount     float64
	SuggestedInterestAmount       float64
	SuggestedTotalAmount          float64
}
