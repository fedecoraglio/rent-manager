package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	policy "rent-manager-backend/internal/core/policy/user"
	"rent-manager-backend/internal/core/port"
)

type CreateUserUseCase struct {
	userRepository   port.UserRepository
	userCreatePolicy *policy.UserCreatePolicy
	passwordHasher   port.PasswordHasher
}

func NewCreateUserUseCase(
	userRepository port.UserRepository,
	userCreatePolicy *policy.UserCreatePolicy,
	passwordHasher port.PasswordHasher,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository:   userRepository,
		userCreatePolicy: userCreatePolicy,
		passwordHasher:   passwordHasher,
	}
}

func (uc *CreateUserUseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.userCreatePolicy.Execute(ctx, user); err != nil {
		return nil, err
	}

	passwordHash, err := uc.passwordHasher.Hash(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = passwordHash

	createdUser, err := uc.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
