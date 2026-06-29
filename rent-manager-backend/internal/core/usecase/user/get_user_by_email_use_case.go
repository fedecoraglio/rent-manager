package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetUserByEmailUseCase struct {
	userRepository port.UserRepository
}

func NewGetUserByEmailUseCase(
	userRepository port.UserRepository,
) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserByEmailUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
