package role

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListRolesUseCase struct {
	roleRepository port.RoleRepository
}

func NewListRolesUseCase(roleRepository port.RoleRepository) *ListRolesUseCase {
	return &ListRolesUseCase{roleRepository: roleRepository}
}

func (lr *ListRolesUseCase) ListRoles(ctx context.Context, page uint64, limit uint64) ([]domain.Role, error) {
	roles, err := lr.roleRepository.ListRoles(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
