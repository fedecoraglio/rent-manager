package policy

import (
	"context"
	"errors"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type RentPaymentCreatePolicy struct {
	rentPaymentRepository    port.RentPaymentRepository
	rentalContractRepository port.RentalContractRepository
}

func NewRentPaymentCreatePolicy(
	rentPaymentRepository port.RentPaymentRepository,
	rentalContractRepository port.RentalContractRepository,
) *RentPaymentCreatePolicy {
	return &RentPaymentCreatePolicy{
		rentPaymentRepository:    rentPaymentRepository,
		rentalContractRepository: rentalContractRepository,
	}
}

func (policy *RentPaymentCreatePolicy) Execute(ctx context.Context, rentPayment *domain.RentPayment) error {
	if err := policy.validate(rentPayment); err != nil {
		return err
	}

	existingPayment, err := policy.rentPaymentRepository.GetRentPaymentByContractIDAndPeriod(
		ctx,
		rentPayment.RentalContractID,
		rentPayment.Period,
	)
	if err != nil {
		var appErr *domain.AppError

		if errors.As(err, &appErr) && appErr.Code == domain.ErrCodeDataNotFound {
			return nil
		}

		return err
	}

	if existingPayment != nil {
		return domain.ErrRentPaymentAlreadyExists
	}

	return nil
}

func (policy *RentPaymentCreatePolicy) validate(rentPayment *domain.RentPayment) error {
	if rentPayment == nil {
		return domain.ErrRentPaymentNil
	}

	rentPayment.Notes = strings.TrimSpace(rentPayment.Notes)

	if rentPayment.RentalContractID <= 0 {
		return domain.ErrRentPaymentContractIDEmpty
	}

	if rentPayment.Period.IsZero() {
		return domain.ErrRentPaymentPeriodEmpty
	}

	if rentPayment.DueDate.IsZero() {
		return domain.ErrRentPaymentDueDateEmpty
	}

	if rentPayment.BaseAmount <= 0 {
		return domain.ErrRentPaymentBaseAmountInvalid
	}

	if rentPayment.TotalAmount <= 0 {
		return domain.ErrRentPaymentTotalAmountInvalid
	}

	if rentPayment.PaidAmount < 0 {
		return domain.ErrRentPaymentPaidAmountInvalid
	}

	return nil
}
