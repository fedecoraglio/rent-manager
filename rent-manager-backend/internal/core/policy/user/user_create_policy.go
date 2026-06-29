package policy

import (
	"context"
	"errors"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/shared"
	"rent-manager-backend/internal/core/port"
	"strings"
)

type UserCreatePolicy struct {
	userRepository port.UserRepository
}

func NewUserCreatePolicy(userRepository port.UserRepository) *UserCreatePolicy {
	return &UserCreatePolicy{
		userRepository: userRepository,
	}
}

func (userPolicy *UserCreatePolicy) Execute(ctx context.Context, user *domain.User) error {
	if user == nil {
		return domain.ErrUserNil
	}

	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Name == "" {
		return domain.ErrUserNameEmpty
	}

	if user.Email == "" {
		return domain.ErrUserEmailEmpty
	}

	if !shared.IsValidEmail(user.Email) {
		return domain.ErrUserEmailInvalid
	}

	if user.PasswordHash == "" {
		return domain.ErrUserPasswordEmpty
	}

	existingUser, err := userPolicy.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		var appErr *domain.AppError
		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return domain.WrapAppError(
			domain.ErrCodeUserEmailLookupError,
			"error getting user by email",
			err,
		)
	}

	if existingUser != nil {
		return domain.ErrUserAlreadyExists
	}

	return nil
}
