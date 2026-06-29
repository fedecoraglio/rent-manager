package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	userPolicy "rent-manager-backend/internal/core/policy/user"
	"rent-manager-backend/internal/core/port"
)

type UpdateUserUseCase struct {
	userRepository   port.UserRepository
	userUpdatePolicy *userPolicy.UserUpdatePolicy
}

func NewUpdateUserUseCase(userRepository port.UserRepository, policy *userPolicy.UserUpdatePolicy) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository, userUpdatePolicy: policy}
}

func (uc *UpdateUserUseCase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := uc.userUpdatePolicy.Execute(ctx, user)
	if err != nil {
		return nil, err
	}
	updatedUser, err := uc.userRepository.PathUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
