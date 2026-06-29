package rent_payment

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/policy/rent_payment"
	"rent-manager-backend/internal/core/port"
)

type UpdateRentPaymentUseCase struct {
	rentPaymentRepository   port.RentPaymentRepository
	rentPaymentUpdatePolicy *policy.RentPaymentUpdatePolicy
}

func NewUpdateRentPaymentUseCase(
	rentPaymentRepository port.RentPaymentRepository,
	rentPaymentUpdatePolicy *policy.RentPaymentUpdatePolicy,
) *UpdateRentPaymentUseCase {
	return &UpdateRentPaymentUseCase{
		rentPaymentRepository:   rentPaymentRepository,
		rentPaymentUpdatePolicy: rentPaymentUpdatePolicy,
	}
}

func (uc *UpdateRentPaymentUseCase) UpdateRentPayment(
	ctx context.Context,
	rentPayment *domain.RentPayment,
) (*domain.RentPayment, error) {
	if err := uc.rentPaymentUpdatePolicy.Execute(ctx, rentPayment); err != nil {
		slog.Error("[UpdateRentPaymentUseCase] policy validation failed", "err", err)
		return nil, err
	}

	updatedRentPayment, err := uc.rentPaymentRepository.UpdateRentPayment(ctx, rentPayment)
	if err != nil {
		slog.Error("[UpdateRentPaymentUseCase] repository update failed", "err", err)
		return nil, err
	}

	return updatedRentPayment, nil
}
