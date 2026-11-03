package service

import (
	"math"
	"testing"
	"time"

	"rent-manager-backend/internal/core/domain"
)

func TestInterestStartsOnlyAfterDueDate(t *testing.T) {
	period := time.Date(2026, time.November, 1, 0, 0, 0, 0, time.UTC)
	for _, mode := range []struct {
		code              string
		firstLateInterest float64
	}{
		{interestTypeDailyFromMonthFirstDay, 80649.36},
		{interestTypeDailyFromDueDayNextDay, 7331.76},
	} {
		t.Run(mode.code, func(t *testing.T) {
			contract := &domain.RentalContract{
				DueDay:                  10,
				DailyInterestPercentage: 1,
				InterestCalculationType: &domain.InterestCalculationType{Code: mode.code},
			}
			for _, tc := range []struct {
				name        string
				paymentDate time.Time
				want        float64
			}{
				{"before_due_date", period.AddDate(0, 0, 2), 0},
				{"on_due_date", period.AddDate(0, 0, 9), 0},
				{"end_of_due_date", period.AddDate(0, 0, 10).Add(-time.Nanosecond), 0},
				{"first_overdue_day", period.AddDate(0, 0, 10), mode.firstLateInterest},
			} {
				t.Run(tc.name, func(t *testing.T) {
					got := calculateInterestAmount(contract, period, tc.paymentDate, 733176)
					if math.Abs(got-tc.want) > 0.000001 {
						t.Fatalf("interest = %v, want %v", got, tc.want)
					}
				})
			}
		})
	}
}
