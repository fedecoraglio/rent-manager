package http

import (
	"rent-manager-backend/internal/core/domain"
)

type roleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func newRoleResponse(role *domain.Role) roleResponse {
	return roleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

func newRolesResponse(roles []domain.Role) []roleResponse {
	response := make([]roleResponse, 0, len(roles))

	for _, role := range roles {
		roleCopy := role

		response = append(response, newRoleResponse(&roleCopy))
	}

	return response
}
