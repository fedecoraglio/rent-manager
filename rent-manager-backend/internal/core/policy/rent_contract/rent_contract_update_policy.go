package policy

import (
	"context"
	"errors"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type RentalContractUpdatePolicy struct {
	rentalContractRepository port.RentalContractRepository
}

func NewRentalContractUpdatePolicy(
	rentalContractRepository port.RentalContractRepository,
) *RentalContractUpdatePolicy {
	return &RentalContractUpdatePolicy{
		rentalContractRepository: rentalContractRepository,
	}
}

func (policy *RentalContractUpdatePolicy) Execute(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) error {
	if rentalContract == nil {
		return domain.ErrRentalContractNil
	}

	if rentalContract.ID <= 0 {
		return domain.ErrRentalContractNotFound
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

	if activeContract != nil && activeContract.ID != rentalContract.ID {
		return domain.ErrRentalContractAlreadyExists
	}

	return nil
}
