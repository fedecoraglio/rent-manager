package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListUserUseCase struct {
	userRepository port.UserRepository
}

func NewListUsersUseCase(
	userRepository port.UserRepository,
) *ListUserUseCase {
	return &ListUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *ListUserUseCase) ListUsers(ctx context.Context, page uint64, limit uint64) ([]domain.User, error) {
	users, err := uc.userRepository.ListUsers(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
