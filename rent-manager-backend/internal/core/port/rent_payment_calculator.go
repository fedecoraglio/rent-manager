package port

import (
	"context"
	"time"

	"rent-manager-backend/internal/core/domain"
)

type RentPaymentCalculator interface {
	Calculate(
		ctx context.Context,
		rentalContract *domain.RentalContract,
		period time.Time,
		paymentDate time.Time,
	) (*domain.RentPaymentCalculation, error)
}
