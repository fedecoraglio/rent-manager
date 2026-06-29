package domain

type RentPaymentCalculation struct {
	BaseAmount           float64
	AdjustmentPercentage float64
	AdjustmentAmount     float64
	AdjustmentStatus     string
	AdjustmentMessage    string
	InterestAmount       float64
	TotalAmount          float64
}
