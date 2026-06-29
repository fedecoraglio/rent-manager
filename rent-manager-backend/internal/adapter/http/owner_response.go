package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type ownerResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func newOwnerResponse(owner *domain.Owner) ownerResponse {
	return ownerResponse{
		ID:             owner.ID,
		Name:           owner.Name,
		Email:          owner.Email,
		Phone:          owner.Phone,
		DocumentNumber: owner.DocumentNumber,
		CreatedAt:      owner.CreatedAt,
		UpdatedAt:      owner.UpdatedAt,
	}
}

func newOwnersResponse(owners []domain.Owner) []ownerResponse {
	response := make([]ownerResponse, 0, len(owners))

	for _, owner := range owners {
		ownerCopy := owner
		response = append(response, newOwnerResponse(&ownerCopy))
	}

	return response
}
