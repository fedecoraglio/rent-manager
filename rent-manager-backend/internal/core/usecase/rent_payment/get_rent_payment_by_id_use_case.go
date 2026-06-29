package rent_payment

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetRentPaymentByIDUseCase struct {
	rentPaymentRepository port.RentPaymentRepository
}

func NewGetRentPaymentByIDUseCase(
	rentPaymentRepository port.RentPaymentRepository,
) *GetRentPaymentByIDUseCase {
	return &GetRentPaymentByIDUseCase{
		rentPaymentRepository: rentPaymentRepository,
	}
}

func (uc *GetRentPaymentByIDUseCase) GetRentPaymentByID(
	ctx context.Context,
	id int64,
) (*domain.RentPayment, error) {
	rentPayment, err := uc.rentPaymentRepository.GetRentPaymentByID(ctx, id)
	if err != nil {
		slog.Error("[GetRentPaymentByIDUseCase] repository get by id failed", "err", err)
		return nil, err
	}

	return rentPayment, nil
}
