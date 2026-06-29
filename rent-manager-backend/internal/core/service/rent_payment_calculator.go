package service

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
)

const (
	interestTypeDailyFromDueDayNextDay = "daily_from_due_day_next_day"
	interestTypeDailyFromMonthFirstDay = "daily_from_month_first_day"
)

type InflationAdjustmentCalculator interface {
	CalculateSuggestedAdjustment(
		ctx context.Context,
		rentalContract *domain.RentalContract,
		period time.Time,
	) (*domain.AdjustmentSuggestion, error)
}

type RentPaymentCalculator struct {
	inflationAdjustmentCalculator InflationAdjustmentCalculator
}

func NewRentPaymentCalculator(
	inflationAdjustmentCalculator InflationAdjustmentCalculator,
) *RentPaymentCalculator {
	return &RentPaymentCalculator{
		inflationAdjustmentCalculator: inflationAdjustmentCalculator,
	}
}

func (calculator *RentPaymentCalculator) Calculate(
	ctx context.Context,
	rentalContract *domain.RentalContract,
	period time.Time,
	paymentDate time.Time,
) (*domain.RentPaymentCalculation, error) {
	baseAmount := rentalContract.MonthlyAmount

	adjustmentSuggestion := &domain.AdjustmentSuggestion{
		Percentage: 0,
		Status:     domain.AdjustmentStatusNotApplicable,
		Message:    "",
	}
	if calculator.inflationAdjustmentCalculator != nil {
		calculatedSuggestion, err := calculator.inflationAdjustmentCalculator.
			CalculateSuggestedAdjustment(ctx, rentalContract, period)
		if err != nil {
			return nil, err
		}

		if calculatedSuggestion != nil {
			adjustmentSuggestion = calculatedSuggestion
		}
	}
	adjustmentPercentage := adjustmentSuggestion.Percentage
	adjustmentAmount := baseAmount * adjustmentPercentage / 100

	interestAmount := calculateInterestAmount(
		rentalContract,
		period,
		paymentDate,
		baseAmount+adjustmentAmount,
	)

	totalAmount := baseAmount + adjustmentAmount + interestAmount

	return &domain.RentPaymentCalculation{
		BaseAmount:           baseAmount,
		AdjustmentPercentage: adjustmentPercentage,
		AdjustmentAmount:     adjustmentAmount,
		AdjustmentStatus:     adjustmentSuggestion.Status,
		AdjustmentMessage:    adjustmentSuggestion.Message,
		InterestAmount:       interestAmount,
		TotalAmount:          totalAmount,
	}, nil
}

func calculateInterestAmount(
	rentalContract *domain.RentalContract,
	period time.Time,
	paymentDate time.Time,
	amount float64,
) float64 {
	if rentalContract.DailyInterestPercentage <= 0 {
		return 0
	}

	if rentalContract.InterestCalculationType == nil {
		return 0
	}

	startDate := getInterestStartDate(
		rentalContract,
		period,
	)

	if paymentDate.Before(startDate) {
		return 0
	}

	lateDays := int(paymentDate.Sub(startDate).Hours()/24) + 1
	if lateDays <= 0 {
		return 0
	}

	return amount * (rentalContract.DailyInterestPercentage / 100) * float64(lateDays)
}

func getInterestStartDate(rentalContract *domain.RentalContract, period time.Time) time.Time {
	normalizedPeriod := firstDayOfMonth(period)

	switch rentalContract.InterestCalculationType.Code {
	case interestTypeDailyFromDueDayNextDay:
		return time.Date(
			normalizedPeriod.Year(),
			normalizedPeriod.Month(),
			int(rentalContract.DueDay)+1,
			0,
			0,
			0,
			0,
			time.UTC,
		)

	case interestTypeDailyFromMonthFirstDay:
		return normalizedPeriod

	default:
		return time.Date(
			normalizedPeriod.Year(),
			normalizedPeriod.Month(),
			int(rentalContract.DueDay)+1,
			0,
			0,
			0,
			0,
			time.UTC,
		)
	}
}

func firstDayOfMonth(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), 1, 0, 0, 0, 0, time.UTC)
}
