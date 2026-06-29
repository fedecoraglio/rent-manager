package http

import (
	"rent-manager-backend/internal/core/domain"
	"time"
)

type userResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func newUsersResponse(users []domain.User) []userResponse {
	response := make([]userResponse, 0, len(users))

	for _, user := range users {
		userCopy := user

		response = append(response, newUserResponse(&userCopy))
	}

	return response
}
