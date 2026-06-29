package http

import "rent-manager-backend/internal/core/domain"

type createOwnerRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"omitempty,email"`
	Phone          string `json:"phone"`
	DocumentNumber string `json:"document_number" binding:"required"`
}

type updateOwnerRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"omitempty,email"`
	Phone          string `json:"phone"`
	DocumentNumber string `json:"document_number" binding:"required"`
}

type listOwnersRequest struct {
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

type searchOwnersRequest struct {
	Value string `form:"value" binding:"required"`
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}

func newOwnerFromCreateRequest(req createOwnerRequest) *domain.Owner {
	return &domain.Owner{
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		DocumentNumber: req.DocumentNumber,
	}
}

func newOwnerFromUpdateRequest(id int64, req updateOwnerRequest) *domain.Owner {
	return &domain.Owner{
		ID:             id,
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		DocumentNumber: req.DocumentNumber,
	}
}
