package policy

import (
	"context"
	"errors"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/shared"
	"rent-manager-backend/internal/core/port"
	"strings"
)

type UserUpdatePolicy struct {
	userRepository port.UserRepository
}

func NewUserUpdatePolicy(userRepository port.UserRepository) *UserUpdatePolicy {
	return &UserUpdatePolicy{
		userRepository: userRepository,
	}
}

func (userPolicy *UserUpdatePolicy) Execute(ctx context.Context, user *domain.User) error {
	if user == nil {
		return domain.ErrUserNil
	}

	if user.ID == 0 {
		return domain.ErrUserNotFound
	}

	if user.Name == "" {
		user.Name = strings.TrimSpace(user.Name)
		return domain.ErrUserNameEmpty
	}

	if user.Email != "" {
		user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	}

	if !shared.IsValidEmail(user.Email) {
		return domain.ErrUserEmailInvalid
	}

	if user.PasswordHash == "" {
		user.Email = strings.TrimSpace(strings.ToLower(user.PasswordHash))
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

	if existingUser != nil && existingUser.ID != user.ID {
		return domain.ErrUserAlreadyExists
	}

	return nil
}
