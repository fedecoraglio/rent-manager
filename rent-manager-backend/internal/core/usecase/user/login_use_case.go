package user

import (
	"context"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type LoginUseCase struct {
	userRepository port.UserRepository
	passwordHasher port.PasswordHasher
	tokenProvider  port.TokenProvider
}

func NewLoginUseCase(
	userRepository port.UserRepository,
	passwordHasher port.PasswordHasher,
	tokenProvider port.TokenProvider,
) *LoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenProvider:  tokenProvider,
	}
}

func (uc *LoginUseCase) Login(ctx context.Context, login *domain.Login) (string, error) {
	if login == nil {
		return "", domain.ErrUnauthorized
	}

	user, err := uc.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return "", domain.ErrUnauthorized
	}

	if err := uc.passwordHasher.Compare(user.PasswordHash, login.Password); err != nil {
		return "", domain.ErrUnauthorized
	}

	token, err := uc.tokenProvider.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
