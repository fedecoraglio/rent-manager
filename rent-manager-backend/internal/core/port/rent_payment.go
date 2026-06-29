package port

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
)

type RentPaymentRepository interface {
	CreateRentPayment(ctx context.Context, rentPayment *domain.RentPayment) (*domain.RentPayment, error)
	UpdateRentPayment(ctx context.Context, rentPayment *domain.RentPayment) (*domain.RentPayment, error)
	GetRentPaymentByID(ctx context.Context, id int64) (*domain.RentPayment, error)
	GetRentPaymentByContractIDAndPeriod(
		ctx context.Context,
		rentalContractID int64,
		period time.Time,
	) (*domain.RentPayment, error)
	ListRentPayments(
		ctx context.Context,
		rentalContractID int64,
		page uint64,
		limit uint64,
	) ([]domain.RentPayment, error)
	ListRentPaymentsByContractID(ctx context.Context, rentalContractID int64) ([]domain.RentPayment, error)
}
