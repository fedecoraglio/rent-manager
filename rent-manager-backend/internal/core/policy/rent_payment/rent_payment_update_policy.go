package policy

import (
	"context"
	"errors"
	"rent-manager-backend/internal/core/policy/shared"
	"strings"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type RentPaymentUpdatePolicy struct {
	rentPaymentRepository    port.RentPaymentRepository
	rentalContractRepository port.RentalContractRepository
}

func NewRentPaymentUpdatePolicy(
	rentPaymentRepository port.RentPaymentRepository,
	rentalContractRepository port.RentalContractRepository,
) *RentPaymentUpdatePolicy {
	return &RentPaymentUpdatePolicy{
		rentPaymentRepository:    rentPaymentRepository,
		rentalContractRepository: rentalContractRepository,
	}
}

func (policy *RentPaymentUpdatePolicy) Execute(ctx context.Context, rentPayment *domain.RentPayment) error {
	if rentPayment == nil {
		return domain.ErrRentPaymentNil
	}

	if rentPayment.ID <= 0 {
		return domain.ErrRentPaymentNotFound
	}

	if err := policy.validate(rentPayment); err != nil {
		return err
	}

	rentalContract, err := policy.rentalContractRepository.GetRentalContractByID(
		ctx,
		rentPayment.RentalContractID,
	)
	if err != nil {
		return domain.WrapAppError(
			domain.ErrCodeRentPaymentContractLookupError,
			"error getting rental contract for rent payment",
			err,
		)
	}

	if !shared.IsPeriodWithinContract(rentPayment.Period, rentalContract) {
		return domain.ErrRentPaymentPeriodOutsideContract
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

	if existingPayment != nil && existingPayment.ID != rentPayment.ID {
		return domain.ErrRentPaymentAlreadyExists
	}

	return nil
}

func (policy *RentPaymentUpdatePolicy) validate(rentPayment *domain.RentPayment) error {
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
