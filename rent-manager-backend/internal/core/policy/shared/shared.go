package shared

import (
	"rent-manager-backend/internal/core/domain"
	"strings"
	"time"
)

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}

	atIndex := strings.Index(email, "@")
	dotIndex := strings.LastIndex(email, ".")

	return atIndex > 0 &&
		dotIndex > atIndex+1 &&
		dotIndex < len(email)-1
}

func CalculateTotalPayments(startDate *time.Time, endDate *time.Time) int64 {
	if startDate == nil || endDate == nil || startDate.IsZero() || endDate.IsZero() {
		return 0
	}

	years := endDate.Year() - startDate.Year()
	months := int(endDate.Month()) - int(startDate.Month())

	totalMonths := years*12 + months + 1

	if totalMonths < 0 {
		return 0
	}

	return int64(totalMonths)
}

func IsPeriodWithinContract(period time.Time, rentalContract *domain.RentalContract) bool {
	normalizedPeriod := FirstDayOfMonth(period)
	startPeriod := FirstDayOfMonth(rentalContract.StartDate)
	endPeriod := FirstDayOfMonth(rentalContract.EndDate)
	return !normalizedPeriod.Before(startPeriod) && !normalizedPeriod.After(endPeriod)
}

func FirstDayOfMonth(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), 1, 0, 0, 0, 0, time.UTC)
}
