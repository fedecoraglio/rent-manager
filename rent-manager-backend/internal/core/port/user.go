package port

import (
	"context"
	"rent-manager-backend/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, page uint64, limit uint64) ([]domain.User, error)
	SearchUsersByNameOrEmail(ctx context.Context, value string, page uint64, limit uint64) ([]domain.User, error)
	PathUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
