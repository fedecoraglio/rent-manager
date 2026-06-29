package domain

var (
	ErrTenantNil = NewAppError(
		ErrCodeTenantNil,
		"tenant is nil",
	)

	ErrTenantNameEmpty = NewAppError(
		ErrCodeTenantNameEmpty,
		"tenant name is empty",
	)

	ErrTenantDocumentNumberEmpty = NewAppError(
		ErrCodeTenantDocumentNumberEmpty,
		"tenant document number is empty",
	)

	ErrTenantAlreadyExists = NewAppError(
		ErrCodeTenantAlreadyExists,
		"tenant already exists",
	)

	ErrTenantDocumentLookupError = NewAppError(
		ErrCodeTenantDocumentLookupError,
		"error getting tenant by document number",
	)

	ErrTenantNotFound = NewAppError(
		ErrCodeTenantNotFound,
		"tenant not found",
	)

	ErrTenantInvalidEmail = NewAppError(
		ErrCodeTenantInvalidEmail,
		"tenant email is invalid",
	)
)
