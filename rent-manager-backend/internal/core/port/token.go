package port

import "rent-manager-backend/internal/core/domain"

type TokenProvider interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.TokenClaims, error)
}
