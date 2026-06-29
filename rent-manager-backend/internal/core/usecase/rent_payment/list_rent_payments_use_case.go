package rent_payment

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type ListRentPaymentsUseCase struct {
	rentPaymentRepository port.RentPaymentRepository
}

func NewListRentPaymentsUseCase(
	rentPaymentRepository port.RentPaymentRepository,
) *ListRentPaymentsUseCase {
	return &ListRentPaymentsUseCase{
		rentPaymentRepository: rentPaymentRepository,
	}
}

func (uc *ListRentPaymentsUseCase) ListRentPayments(
	ctx context.Context,
	rentalContractID int64,
	page uint64,
	limit uint64,
) ([]domain.RentPayment, error) {
	rentPayments, err := uc.rentPaymentRepository.ListRentPayments(
		ctx,
		rentalContractID,
		page,
		limit,
	)
	if err != nil {
		slog.Error("[ListRentPaymentsUseCase] repository list failed", "err", err)
		return nil, err
	}

	return rentPayments, nil
}
