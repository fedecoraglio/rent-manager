package policy

import (
	"context"
	"errors"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type OwnerCreatePolicy struct {
	ownerRepository port.OwnerRepository
}

func NewOwnerCreatePolicy(ownerRepository port.OwnerRepository) *OwnerCreatePolicy {
	return &OwnerCreatePolicy{
		ownerRepository: ownerRepository,
	}
}

func (ownerPolicy *OwnerCreatePolicy) Execute(ctx context.Context, owner *domain.Owner) error {
	if owner == nil {
		return domain.ErrOwnerNil
	}

	owner.Name = strings.TrimSpace(owner.Name)
	owner.Email = strings.TrimSpace(strings.ToLower(owner.Email))
	owner.Phone = strings.TrimSpace(owner.Phone)
	owner.DocumentNumber = strings.TrimSpace(owner.DocumentNumber)

	if owner.Name == "" {
		return domain.ErrOwnerNameEmpty
	}

	if owner.DocumentNumber == "" {
		return domain.ErrOwnerDocumentNumberEmpty
	}

	if owner.Email != "" && !isValidEmail(owner.Email) {
		return domain.ErrOwnerInvalidEmail
	}

	existingOwner, err := ownerPolicy.ownerRepository.GetOwnerByDocumentNumber(ctx, owner.DocumentNumber)
	if err != nil {
		var appErr *domain.AppError
		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return domain.WrapAppError(
			domain.ErrCodeOwnerDocumentLookupError,
			"error getting owner by document number",
			err,
		)
	}

	if existingOwner != nil {
		return domain.ErrOwnerAlreadyExists
	}

	return nil
}

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	atIndex := strings.Index(email, "@")
	dotIndex := strings.LastIndex(email, ".")

	return atIndex > 0 &&
		dotIndex > atIndex+1 &&
		dotIndex < len(email)-1
}
