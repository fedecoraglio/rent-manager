package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/owner"
	"rent-manager-backend/internal/core/port"
)

type CreateOwnerUseCase struct {
	ownerRepository   port.OwnerRepository
	ownerCreatePolicy *policy.OwnerCreatePolicy
}

func NewCreateOwnerUseCase(
	ownerRepository port.OwnerRepository,
	ownerCreatePolicy *policy.OwnerCreatePolicy,
) *CreateOwnerUseCase {
	return &CreateOwnerUseCase{
		ownerRepository:   ownerRepository,
		ownerCreatePolicy: ownerCreatePolicy,
	}
}

func (uc *CreateOwnerUseCase) CreateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	if err := uc.ownerCreatePolicy.Execute(ctx, owner); err != nil {
		slog.Error("[CreateOwnerUseCase] policy validation failed", "err", err)
		return nil, err
	}

	createdOwner, err := uc.ownerRepository.CreateOwner(ctx, owner)
	if err != nil {
		slog.Error("[CreateOwnerUseCase] repository create failed", "err", err)
		return nil, err
	}

	return createdOwner, nil
}
