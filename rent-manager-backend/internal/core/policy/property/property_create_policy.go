package policy

import (
	"context"
	"errors"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type PropertyCreatePolicy struct {
	propertyRepository port.PropertyRepository
}

func NewPropertyCreatePolicy(propertyRepository port.PropertyRepository) *PropertyCreatePolicy {
	return &PropertyCreatePolicy{
		propertyRepository: propertyRepository,
	}
}

func (propertyPolicy *PropertyCreatePolicy) Execute(
	ctx context.Context,
	property *domain.Property,
) error {
	if property == nil {
		return domain.ErrPropertyNil
	}

	property.Code = strings.TrimSpace(property.Code)
	property.Title = strings.TrimSpace(property.Title)
	property.Description = strings.TrimSpace(property.Description)
	property.Street = strings.TrimSpace(property.Street)
	property.StreetNumber = strings.TrimSpace(property.StreetNumber)
	property.Floor = strings.TrimSpace(property.Floor)
	property.Apartment = strings.TrimSpace(property.Apartment)
	property.City = strings.TrimSpace(property.City)
	property.PostalCode = strings.TrimSpace(property.PostalCode)

	if property.OwnerID <= 0 {
		return domain.ErrPropertyOwnerIDEmpty
	}

	if property.TypeID <= 0 {
		return domain.ErrPropertyTypeIDEmpty
	}

	if property.StatusID <= 0 {
		return domain.ErrPropertyStatusIDEmpty
	}

	if property.CountryID <= 0 {
		return domain.ErrPropertyCountryIDEmpty
	}

	if property.StateID <= 0 {
		return domain.ErrPropertyStateIDEmpty
	}

	if property.Code == "" {
		return domain.ErrPropertyCodeEmpty
	}

	if property.Title == "" {
		return domain.ErrPropertyTitleEmpty
	}

	if property.Street == "" {
		return domain.ErrPropertyStreetEmpty
	}

	if property.City == "" {
		return domain.ErrPropertyCityEmpty
	}

	existingProperty, err := propertyPolicy.propertyRepository.GetPropertyByCode(ctx, property.Code)
	if err != nil {
		var appErr *domain.AppError

		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return domain.WrapAppError(
			domain.ErrCodePropertyCodeLookupError,
			"error getting property by code",
			err,
		)
	}

	if existingProperty != nil {
		return domain.ErrPropertyAlreadyExists
	}

	return nil
}
