package http

import "rent-manager-backend/internal/core/domain"

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RoleID   int64  `json:"role_id" binding:"required,min=1"`
}

type updateUserRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	PasswordHash string
	RoleID       int64 `json:"role_id" binding:"required,min=1"`
}

type searchUsersRequest struct {
	Value string `form:"value" binding:"required"`
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newUserFromCreateRequest(req createUserRequest) *domain.User {
	return &domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		RoleID:       req.RoleID,
	}
}

func newUserFromUpdateRequest(id int64, req updateUserRequest) *domain.User {
	return &domain.User{
		ID:           id,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		RoleID:       req.RoleID,
	}
}
