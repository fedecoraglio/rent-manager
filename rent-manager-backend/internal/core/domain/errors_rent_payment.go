package domain

var (
	ErrRentPaymentNil = NewAppError(
		ErrCodeRentPaymentNil,
		"rent payment is nil",
	)

	ErrRentPaymentContractIDEmpty = NewAppError(
		ErrCodeRentPaymentContractIDEmpty,
		"rent payment contract id is empty",
	)

	ErrRentPaymentPeriodEmpty = NewAppError(
		ErrCodeRentPaymentPeriodEmpty,
		"rent payment period is empty",
	)

	ErrRentPaymentDueDateEmpty = NewAppError(
		ErrCodeRentPaymentDueDateEmpty,
		"rent payment due date is empty",
	)

	ErrRentPaymentBaseAmountInvalid = NewAppError(
		ErrCodeRentPaymentBaseAmountInvalid,
		"rent payment base amount is invalid",
	)

	ErrRentPaymentTotalAmountInvalid = NewAppError(
		ErrCodeRentPaymentTotalAmountInvalid,
		"rent payment total amount is invalid",
	)

	ErrRentPaymentPaidAmountInvalid = NewAppError(
		ErrCodeRentPaymentPaidAmountInvalid,
		"rent payment paid amount is invalid",
	)

	ErrRentPaymentAlreadyExists = NewAppError(
		ErrCodeRentPaymentAlreadyExists,
		"rent payment already exists for this period",
	)

	ErrRentPaymentNotFound = NewAppError(
		ErrCodeRentPaymentNotFound,
		"rent payment not found",
	)

	ErrRentPaymentContractLookupError = NewAppError(
		ErrCodeRentPaymentContractLookupError,
		"error getting rental contract for rent payment",
	)

	ErrRentPaymentPeriodAlreadyPaid = NewAppError(
		ErrCodeRentPaymentPeriodAlreadyPaid,
		"rent payment period is already paid",
	)

	ErrRentPaymentPeriodOutsideContract = NewAppError(
		ErrCodeRentPaymentPeriodOutsideContract,
		"rent payment period is outside rental contract date range",
	)
)
