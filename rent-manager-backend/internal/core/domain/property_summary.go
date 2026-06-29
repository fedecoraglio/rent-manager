package domain

type PropertySummary struct {
	PropertyID            int64
	PropertyTitle         string
	RentalContractSummary *RentalContractSummary
}
