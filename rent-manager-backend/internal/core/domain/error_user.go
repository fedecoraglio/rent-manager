package domain

var (
	ErrDataNotFound = NewAppError(
		ErrCodeDataNotFound,
		"data not found",
	)

	ErrConflictingData = NewAppError(
		ErrCodeConflictingData,
		"conflicting data",
	)

	ErrInvalidData = NewAppError(
		ErrCodeInvalidData,
		"invalid data",
	)

	ErrUnauthorized = NewAppError(
		ErrCodeUnauthorized,
		"unauthorized",
	)

	ErrForbidden = NewAppError(
		ErrCodeForbidden,
		"forbidden",
	)

	ErrUserNil = NewAppError(
		ErrCodeUserNil,
		"user is nil",
	)

	ErrUserNameEmpty = NewAppError(
		ErrCodeUserNameEmpty,
		"user name is empty",
	)

	ErrUserEmailEmpty = NewAppError(
		ErrCodeUserEmailEmpty,
		"user email is empty",
	)

	ErrUserEmailInvalid = NewAppError(
		ErrCodeUserEmailInvalid,
		"user email is invalid",
	)

	ErrUserPasswordEmpty = NewAppError(
		ErrCodeUserPasswordEmpty,
		"user password is empty",
	)

	ErrUserAlreadyExists = NewAppError(
		ErrCodeUserAlreadyExists,
		"user already exists",
	)

	ErrUserEmailLookupError = NewAppError(
		ErrCodeUserEmailLookupError,
		"error getting user by email",
	)

	ErrUserNotFound = NewAppError(
		ErrCodeUserNotFound,
		"user not found",
	)

	ErrUserRoleEmpty = NewAppError(
		ErrCodeUserRoleEmpty,
		"user role is empty",
	)
)
