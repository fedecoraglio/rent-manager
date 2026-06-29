package port

import (
	"context"

	"rent-manager-backend/internal/core/domain"
)

type RentalContractRepository interface {
	CreateRentalContract(ctx context.Context, rentalContract *domain.RentalContract) (*domain.RentalContract, error)
	UpdateRentalContract(ctx context.Context, rentalContract *domain.RentalContract) (*domain.RentalContract, error)
	GetRentalContractByID(ctx context.Context, id int64) (*domain.RentalContract, error)
	GetActiveRentalContractByPropertyID(ctx context.Context, propertyID int64) (*domain.RentalContract, error)
	ListRentalContracts(ctx context.Context, propertyID int64, page uint64, limit uint64) ([]domain.RentalContract, error)
	ListActiveRentalContracts(ctx context.Context, page uint64, limit uint64) ([]domain.RentalContract, error)
}
