package owner

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/owner"
	"rent-manager-backend/internal/core/port"
)

type UpdateOwnerUseCase struct {
	ownerRepository   port.OwnerRepository
	ownerUpdatePolicy *policy.OwnerUpdatePolicy
}

func NewUpdateOwnerUseCase(
	ownerRepository port.OwnerRepository,
	ownerUpdatePolicy *policy.OwnerUpdatePolicy,
) *UpdateOwnerUseCase {
	return &UpdateOwnerUseCase{
		ownerRepository:   ownerRepository,
		ownerUpdatePolicy: ownerUpdatePolicy,
	}
}

func (uc *UpdateOwnerUseCase) UpdateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	if err := uc.ownerUpdatePolicy.Execute(ctx, owner); err != nil {
		slog.Error("[UpdateOwnerUseCase] policy validation failed", "err", err)
		return nil, err
	}

	updatedOwner, err := uc.ownerRepository.UpdateOwner(ctx, owner)
	if err != nil {
		slog.Error("[UpdateOwnerUseCase] repository update failed", "err", err)
		return nil, err
	}

	return updatedOwner, nil
}
