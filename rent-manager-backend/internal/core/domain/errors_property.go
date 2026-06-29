package domain

var (
	ErrPropertyNil = NewAppError(
		ErrCodePropertyNil,
		"property is nil",
	)

	ErrPropertyOwnerIDEmpty = NewAppError(
		ErrCodePropertyOwnerIDEmpty,
		"property owner id is empty",
	)

	ErrPropertyTypeIDEmpty = NewAppError(
		ErrCodePropertyTypeIDEmpty,
		"property type id is empty",
	)

	ErrPropertyStatusIDEmpty = NewAppError(
		ErrCodePropertyStatusIDEmpty,
		"property status id is empty",
	)

	ErrPropertyCountryIDEmpty = NewAppError(
		ErrCodePropertyCountryIDEmpty,
		"property country id is empty",
	)

	ErrPropertyStateIDEmpty = NewAppError(
		ErrCodePropertyStateIDEmpty,
		"property state id is empty",
	)

	ErrPropertyCodeEmpty = NewAppError(
		ErrCodePropertyCodeEmpty,
		"property code is empty",
	)

	ErrPropertyTitleEmpty = NewAppError(
		ErrCodePropertyTitleEmpty,
		"property title is empty",
	)

	ErrPropertyStreetEmpty = NewAppError(
		ErrCodePropertyStreetEmpty,
		"property street is empty",
	)

	ErrPropertyCityEmpty = NewAppError(
		ErrCodePropertyCityEmpty,
		"property city is empty",
	)

	ErrPropertyAlreadyExists = NewAppError(
		ErrCodePropertyAlreadyExists,
		"property already exists",
	)

	ErrPropertyCodeLookupError = NewAppError(
		ErrCodePropertyCodeLookupError,
		"error getting property by code",
	)

	ErrPropertyNotFound = NewAppError(
		ErrCodePropertyNotFound,
		"property not found",
	)
)
