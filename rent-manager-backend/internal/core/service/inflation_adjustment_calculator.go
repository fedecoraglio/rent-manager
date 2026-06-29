package service

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

const rentAdjustmentTypeIPCArgentina = "ipc_argentina"

type IPCInflationAdjustmentCalculator struct {
	inflationIndexRepository port.InflationIndexRepository
}

func NewIPCInflationAdjustmentCalculator(
	inflationIndexRepository port.InflationIndexRepository,
) *IPCInflationAdjustmentCalculator {
	return &IPCInflationAdjustmentCalculator{
		inflationIndexRepository: inflationIndexRepository,
	}
}

func (calculator *IPCInflationAdjustmentCalculator) CalculateSuggestedAdjustment(
	ctx context.Context,
	rentalContract *domain.RentalContract,
	period time.Time,
) (*domain.AdjustmentSuggestion, error) {
	var inflationNotApplicable = &domain.AdjustmentSuggestion{
		Percentage: 0,
		Status:     domain.AdjustmentStatusNotApplicable,
		Message:    "",
	}
	if rentalContract == nil {
		return inflationNotApplicable, nil
	}

	if rentalContract.AdjustmentFrequencyMonths <= 0 {
		return inflationNotApplicable, nil
	}

	if rentalContract.AdjustmentType == nil {
		return inflationNotApplicable, nil
	}

	if rentalContract.AdjustmentType.Code != rentAdjustmentTypeIPCArgentina {
		return inflationNotApplicable, nil
	}

	normalizedPeriod := firstDayOfMonth(period)
	normalizedStartDate := firstDayOfMonth(rentalContract.StartDate)

	if normalizedPeriod.Before(normalizedStartDate) {
		return inflationNotApplicable, nil
	}

	monthsFromStart := monthDiff(normalizedStartDate, normalizedPeriod)

	if monthsFromStart == 0 {
		return inflationNotApplicable, nil
	}

	if monthsFromStart%int(rentalContract.AdjustmentFrequencyMonths) != 0 {
		return inflationNotApplicable, nil
	}

	from := normalizedPeriod.AddDate(
		0,
		-int(rentalContract.AdjustmentFrequencyMonths),
		0,
	)

	to := normalizedPeriod.AddDate(0, -1, 0)

	indexes, err := calculator.inflationIndexRepository.ListByPeriodRange(
		ctx,
		from,
		to,
	)
	if err != nil {
		return inflationNotApplicable, err
	}

	expectedIndexes := int(rentalContract.AdjustmentFrequencyMonths)

	// Fall back to the latest available indexes if the expected range is incomplete.
	if len(indexes) != expectedIndexes {
		indexes, err = calculator.inflationIndexRepository.ListLatestBeforePeriod(
			ctx,
			normalizedPeriod,
			uint64(expectedIndexes),
		)
		if err != nil {
			return nil, err
		}

		if len(indexes) != expectedIndexes {
			return &domain.AdjustmentSuggestion{
				Percentage: 0,
				Status:     domain.AdjustmentStatusMissingIndexes,
				Message:    "Inflation indexes are missing for the adjustment period.",
			}, nil
		}
	}

	accumulated := 1.0

	for _, index := range indexes {
		accumulated *= 1 + index.Percentage/100
	}

	return &domain.AdjustmentSuggestion{
		Percentage: (accumulated - 1) * 100,
		Status:     domain.AdjustmentStatusCalculated,
		Message:    "",
	}, nil
}

func monthDiff(from time.Time, to time.Time) int {
	yearDiff := to.Year() - from.Year()
	monthDiff := int(to.Month()) - int(from.Month())

	return yearDiff*12 + monthDiff
}
