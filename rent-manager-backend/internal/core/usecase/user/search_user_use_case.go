package user

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type SearchUserUseCase struct {
	userRepository port.UserRepository
}

func NewSearchUsersUseCase(
	userRepository port.UserRepository,
) *SearchUserUseCase {
	return &SearchUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *SearchUserUseCase) SearchUsersByNameOrEmail(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.User, error) {
	users, err := uc.userRepository.SearchUsersByNameOrEmail(ctx, value, page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
