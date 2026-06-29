package rent_payment

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	rentPaymentPolicy "rent-manager-backend/internal/core/policy/rent_payment"
	"rent-manager-backend/internal/core/port"
)

type CreateRentPaymentUseCase struct {
	rentPaymentRepository   port.RentPaymentRepository
	rentPaymentCreatePolicy *rentPaymentPolicy.RentPaymentCreatePolicy
}

func NewCreateRentPaymentUseCase(
	rentPaymentRepository port.RentPaymentRepository,
	rentPaymentCreatePolicy *rentPaymentPolicy.RentPaymentCreatePolicy,
) *CreateRentPaymentUseCase {
	return &CreateRentPaymentUseCase{
		rentPaymentRepository:   rentPaymentRepository,
		rentPaymentCreatePolicy: rentPaymentCreatePolicy,
	}
}

func (uc *CreateRentPaymentUseCase) CreateRentPayment(
	ctx context.Context,
	rentPayment *domain.RentPayment,
) (*domain.RentPayment, error) {
	if err := uc.rentPaymentCreatePolicy.Execute(ctx, rentPayment); err != nil {
		slog.Error("[CreateRentPaymentUseCase] policy validation failed", "err", err)
		return nil, err
	}

	createdRentPayment, err := uc.rentPaymentRepository.CreateRentPayment(ctx, rentPayment)
	if err != nil {
		slog.Error("[CreateRentPaymentUseCase] repository create failed", "err", err)
		return nil, err
	}

	return createdRentPayment, nil
}
