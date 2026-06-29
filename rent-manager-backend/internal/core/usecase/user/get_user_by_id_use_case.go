package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetUserByIDUseCase struct {
	userRepository port.UserRepository
}

func NewGetUserByIDUseCase(
	userRepository port.UserRepository,
) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserByIDUseCase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
