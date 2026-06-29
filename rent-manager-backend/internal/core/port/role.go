package port

import (
	"context"
	"rent-manager-backend/internal/core/domain"
)

type RoleRepository interface {
	ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error)
}
