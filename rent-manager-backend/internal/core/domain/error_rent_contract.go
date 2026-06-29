package domain

var (
	ErrRentalContractNil = NewAppError(
		ErrCodeRentalContractNil,
		"rental contract is nil",
	)

	ErrRentalContractPropertyIDEmpty = NewAppError(
		ErrCodeRentalContractPropertyIDEmpty,
		"rental contract property id is empty",
	)

	ErrRentalContractTenantIDEmpty = NewAppError(
		ErrCodeRentalContractTenantIDEmpty,
		"rental contract tenant id is empty",
	)

	ErrRentalContractStatusIDEmpty = NewAppError(
		ErrCodeRentalContractStatusIDEmpty,
		"rental contract status id is empty",
	)

	ErrRentalContractStartDateEmpty = NewAppError(
		ErrCodeRentalContractStartDateEmpty,
		"rental contract start date is empty",
	)

	ErrRentalContractEndDateEmpty = NewAppError(
		ErrCodeRentalContractEndDateEmpty,
		"rental contract end date is empty",
	)

	ErrRentalContractInvalidDateRange = NewAppError(
		ErrCodeRentalContractInvalidDateRange,
		"rental contract end date must be after start date",
	)

	ErrRentalContractMonthlyAmountInvalid = NewAppError(
		ErrCodeRentalContractMonthlyAmountInvalid,
		"rental contract monthly amount is invalid",
	)

	ErrRentalContractCurrencyEmpty = NewAppError(
		ErrCodeRentalContractCurrencyEmpty,
		"rental contract currency is empty",
	)

	ErrRentalContractAlreadyExists = NewAppError(
		ErrCodeRentalContractAlreadyExists,
		"property already has an active rental contract",
	)

	ErrRentalContractNotFound = NewAppError(
		ErrCodeRentalContractNotFound,
		"rental contract not found",
	)

	ErrRentalContractActiveLookupError = NewAppError(
		ErrCodeRentalContractActiveLookupError,
		"error getting active rental contract by property",
	)

	ErrRentalContractInterestCalculationTypeIDEmpty = NewAppError(
		ErrCodeRentalContractInterestCalculationTypeIDEmpty,
		"rental contract interest calculation type id is empty",
	)

	ErrRentalContractAdjustmentTypeIDEmpty = NewAppError(
		ErrCodeRentalContractAdjustmentTypeIDEmpty,
		"rental contract adjustment type id is empty",
	)

	ErrRentalContractDepositAmountInvalid = NewAppError(
		ErrCodeRentalContractDepositAmountInvalid,
		"rental contract deposit amount is invalid",
	)

	ErrRentalContractDueDayInvalid = NewAppError(
		ErrCodeRentalContractDueDayInvalid,
		"rental contract due day is invalid",
	)

	ErrRentalContractDailyInterestPercentageInvalid = NewAppError(
		ErrCodeRentalContractDailyInterestPercentageInvalid,
		"rental contract daily interest percentage is invalid",
	)

	ErrRentalContractAdjustmentFrequencyInvalid = NewAppError(
		ErrCodeRentalContractAdjustmentFrequencyInvalid,
		"rental contract adjustment frequency is invalid",
	)

	ErrRentalContractTotalPaymentsInvalid = NewAppError(
		ErrCodeRentalContractTotalPaymentsInvalid,
		"rental contract total payments is invalid",
	)
)
