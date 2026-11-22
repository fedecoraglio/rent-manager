package user

import (
	"context"
	"rent-manager-backend/internal/core/port"
)

type DeleteUserUseCase struct {
	userRepository port.UserRepository
}

func NewDeleteUserUseCase(userRepository port.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id int64) error {
	return uc.userRepository.DeleteUser(ctx, id)
}
