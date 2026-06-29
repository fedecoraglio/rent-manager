package domain

import "time"

type ContractStatus struct {
	ID   int64
	Code string
	Name string
}

type InterestCalculationType struct {
	ID   int64
	Code string
	Name string
}

type RentAdjustmentType struct {
	ID   int64
	Code string
	Name string
}

type RentalContract struct {
	ID                        int64
	PropertyID                int64
	TenantID                  int64
	StatusID                  int64
	InterestCalculationTypeID int64
	AdjustmentTypeID          int64
	StartDate                 time.Time
	EndDate                   time.Time
	TotalPayments             int64
	MonthlyAmount             float64
	DepositAmount             float64
	Currency                  string
	DueDay                    int64
	DailyInterestPercentage   float64
	AdjustmentFrequencyMonths int64
	Notes                     string
	Property                  *Property
	Tenant                    *Tenant
	Status                    *ContractStatus
	InterestCalculationType   *InterestCalculationType
	AdjustmentType            *RentAdjustmentType
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}
