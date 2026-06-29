package domain

var (
	ErrOwnerNil = NewAppError(
		ErrCodeOwnerNil,
		"owner is nil",
	)

	ErrOwnerNameEmpty = NewAppError(
		ErrCodeOwnerNameEmpty,
		"owner name is empty",
	)

	ErrOwnerDocumentNumberEmpty = NewAppError(
		ErrCodeOwnerDocumentNumberEmpty,
		"owner document number is empty",
	)

	ErrOwnerAlreadyExists = NewAppError(
		ErrCodeOwnerAlreadyExists,
		"owner already exists",
	)

	ErrOwnerDocumentLookupError = NewAppError(
		ErrCodeOwnerDocumentLookupError,
		"error getting owner by document number",
	)

	ErrOwnerNotFound = NewAppError(
		ErrCodeOwnerNotFound,
		"owner not found",
	)

	ErrOwnerInvalidEmail = NewAppError(
		ErrCodeOwnerInvalidEmail,
		"owner email is invalid",
	)
)
