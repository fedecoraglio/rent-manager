package policy

import (
	"context"
	"errors"
	"strings"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type RentalContractCreatePolicy struct {
	rentalContractRepository port.RentalContractRepository
}

func NewRentalContractCreatePolicy(
	rentalContractRepository port.RentalContractRepository,
) *RentalContractCreatePolicy {
	return &RentalContractCreatePolicy{
		rentalContractRepository: rentalContractRepository,
	}
}

func (policy *RentalContractCreatePolicy) Execute(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) error {
	if err := policy.validate(rentalContract); err != nil {
		return err
	}

	activeContract, err := policy.rentalContractRepository.GetActiveRentalContractByPropertyID(
		ctx,
		rentalContract.PropertyID,
	)
	if err != nil {
		var appErr *domain.AppError

		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return domain.WrapAppError(
			domain.ErrCodeRentalContractActiveLookupError,
			"error getting active rental contract by property",
			err,
		)
	}

	if activeContract != nil {
		return domain.ErrRentalContractAlreadyExists
	}

	return nil
}

func (policy *RentalContractCreatePolicy) validate(
	rentalContract *domain.RentalContract,
) error {
	if rentalContract == nil {
		return domain.ErrRentalContractNil
	}

	rentalContract.Currency = strings.TrimSpace(strings.ToUpper(rentalContract.Currency))
	rentalContract.Notes = strings.TrimSpace(rentalContract.Notes)

	if rentalContract.PropertyID <= 0 {
		return domain.ErrRentalContractPropertyIDEmpty
	}

	if rentalContract.TenantID <= 0 {
		return domain.ErrRentalContractTenantIDEmpty
	}

	if rentalContract.StatusID <= 0 {
		return domain.ErrRentalContractStatusIDEmpty
	}

	if rentalContract.InterestCalculationTypeID <= 0 {
		return domain.ErrRentalContractInterestCalculationTypeIDEmpty
	}

	if rentalContract.AdjustmentTypeID <= 0 {
		return domain.ErrRentalContractAdjustmentTypeIDEmpty
	}

	if rentalContract.StartDate.IsZero() {
		return domain.ErrRentalContractStartDateEmpty
	}

	if rentalContract.EndDate.IsZero() {
		return domain.ErrRentalContractEndDateEmpty
	}

	if !rentalContract.EndDate.After(rentalContract.StartDate) {
		return domain.ErrRentalContractInvalidDateRange
	}

	if rentalContract.MonthlyAmount <= 0 {
		return domain.ErrRentalContractMonthlyAmountInvalid
	}

	if rentalContract.DepositAmount < 0 {
		return domain.ErrRentalContractDepositAmountInvalid
	}

	if rentalContract.Currency == "" {
		return domain.ErrRentalContractCurrencyEmpty
	}

	if rentalContract.DueDay < 1 || rentalContract.DueDay > 31 {
		return domain.ErrRentalContractDueDayInvalid
	}

	if rentalContract.DailyInterestPercentage < 0 {
		return domain.ErrRentalContractDailyInterestPercentageInvalid
	}

	if rentalContract.AdjustmentFrequencyMonths <= 0 {
		return domain.ErrRentalContractAdjustmentFrequencyInvalid
	}

	if rentalContract.TotalPayments <= 0 {
		return domain.ErrRentalContractTotalPaymentsInvalid
	}

	return nil
}

func normalizeDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}
