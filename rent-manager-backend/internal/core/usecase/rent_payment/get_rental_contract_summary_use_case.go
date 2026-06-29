package rent_payment

import (
	"context"
	"log/slog"
	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
	"time"
)

type GetRentalContractSummaryUseCase struct {
	rentalContractRepository port.RentalContractRepository
	rentPaymentRepository    port.RentPaymentRepository
	rentPaymentCalculator    port.RentPaymentCalculator
}

func NewGetRentalContractSummaryUseCase(
	rentalContractRepository port.RentalContractRepository,
	rentPaymentRepository port.RentPaymentRepository,
	rentPaymentCalculator port.RentPaymentCalculator,
) *GetRentalContractSummaryUseCase {
	return &GetRentalContractSummaryUseCase{
		rentalContractRepository: rentalContractRepository,
		rentPaymentRepository:    rentPaymentRepository,
		rentPaymentCalculator:    rentPaymentCalculator,
	}
}

func (uc *GetRentalContractSummaryUseCase) GetRentalContractSummary(
	ctx context.Context,
	rentalContractID int64,
) (*domain.RentalContractSummary, error) {
	rentalContract, err := uc.rentalContractRepository.GetRentalContractByID(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentalContractSummaryUseCase] repository get rental contract failed", "err", err)
		return nil, err
	}

	rentPayments, err := uc.rentPaymentRepository.ListRentPaymentsByContractID(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentalContractSummaryUseCase] repository list rent payments failed", "err", err)
		return nil, err
	}

	paidByPeriod := make(map[string]domain.RentPayment)

	for _, rentPayment := range rentPayments {
		if rentPayment.IsPaid {
			paidByPeriod[periodKey(rentPayment.Period)] = rentPayment
		}
	}

	paidPayments := int64(len(paidByPeriod))
	remainingPayments := rentalContract.TotalPayments - paidPayments

	if remainingPayments < 0 {
		remainingPayments = 0
	}

	var nextPendingPeriod *time.Time
	currentPeriod := firstDayOfMonth(rentalContract.StartDate)

	for i := int64(0); i < rentalContract.TotalPayments; i++ {
		if _, exists := paidByPeriod[periodKey(currentPeriod)]; !exists {
			periodCopy := currentPeriod
			nextPendingPeriod = &periodCopy
			break
		}

		currentPeriod = currentPeriod.AddDate(0, 1, 0)
	}

	previousPeriod := currentPeriod.AddDate(0, -1, 0)
	payment, exists := paidByPeriod[periodKey(previousPeriod)]
	if exists {
		rentalContract.MonthlyAmount = payment.TotalAmount
	}

	var nextSuggestedAmount = rentalContract.MonthlyAmount
	if nextPendingPeriod != nil {
		calculation, err := uc.rentPaymentCalculator.Calculate(
			ctx,
			rentalContract,
			*nextPendingPeriod,
			time.Now().UTC(),
		)
		if err != nil {
			slog.Error("[GetRentalContractSummaryUseCase] calculator failed", "err", err)
			return nil, err
		}

		nextSuggestedAmount = calculation.TotalAmount
	}

	nextAdjustmentPeriod := calculateNextAdjustmentPeriod(rentalContract)

	return &domain.RentalContractSummary{
		RentalContractID:     rentalContract.ID,
		TotalPayments:        rentalContract.TotalPayments,
		PaidPayments:         paidPayments,
		RemainingPayments:    remainingPayments,
		CurrentAmount:        rentalContract.MonthlyAmount,
		NextSuggestedAmount:  nextSuggestedAmount,
		NextPendingPeriod:    nextPendingPeriod,
		NextAdjustmentPeriod: nextAdjustmentPeriod,
	}, nil
}

func calculateNextAdjustmentPeriod(rentalContract *domain.RentalContract) *time.Time {
	if rentalContract.AdjustmentFrequencyMonths <= 0 {
		return nil
	}

	currentPeriod := firstDayOfMonth(rentalContract.StartDate)
	nowPeriod := firstDayOfMonth(time.Now().UTC())
	endPeriod := firstDayOfMonth(rentalContract.EndDate)

	for !currentPeriod.After(firstDayOfMonth(rentalContract.EndDate)) {
		currentPeriod = currentPeriod.AddDate(
			0,
			int(rentalContract.AdjustmentFrequencyMonths),
			0,
		)

		if currentPeriod.After(endPeriod) {
			return nil
		}

		if currentPeriod.After(nowPeriod) || currentPeriod.Equal(nowPeriod) {
			periodCopy := currentPeriod
			return &periodCopy
		}
	}

	return nil
}
