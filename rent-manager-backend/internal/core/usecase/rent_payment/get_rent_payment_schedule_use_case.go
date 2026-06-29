package rent_payment

import (
	"context"
	"log/slog"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetRentPaymentScheduleUseCase struct {
	rentalContractRepository port.RentalContractRepository
	rentPaymentRepository    port.RentPaymentRepository
	rentPaymentCalculator    port.RentPaymentCalculator
}

func NewGetRentPaymentScheduleUseCase(
	rentalContractRepository port.RentalContractRepository,
	rentPaymentRepository port.RentPaymentRepository,
	rentPaymentCalculator port.RentPaymentCalculator,
) *GetRentPaymentScheduleUseCase {
	return &GetRentPaymentScheduleUseCase{
		rentalContractRepository: rentalContractRepository,
		rentPaymentRepository:    rentPaymentRepository,
		rentPaymentCalculator:    rentPaymentCalculator,
	}
}

func (uc *GetRentPaymentScheduleUseCase) GetRentPaymentSchedule(
	ctx context.Context,
	rentalContractID int64,
) ([]domain.RentPaymentScheduleItem, error) {
	rentalContract, err := uc.rentalContractRepository.GetRentalContractByID(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentPaymentScheduleUseCase] repository get rental contract failed", "err", err)
		return nil, err
	}

	rentPayments, err := uc.rentPaymentRepository.ListRentPaymentsByContractID(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentPaymentScheduleUseCase] repository list rent payments failed", "err", err)
		return nil, err
	}

	paymentsByPeriod := make(map[string]domain.RentPayment)

	for _, rentPayment := range rentPayments {
		paymentsByPeriod[periodKey(rentPayment.Period)] = rentPayment
	}

	return uc.buildPaymentSchedule(ctx, rentalContract, paymentsByPeriod)
}

func (uc *GetRentPaymentScheduleUseCase) buildPaymentSchedule(
	ctx context.Context,
	rentalContract *domain.RentalContract,
	paymentsByPeriod map[string]domain.RentPayment,
) ([]domain.RentPaymentScheduleItem, error) {
	items := make([]domain.RentPaymentScheduleItem, 0, rentalContract.TotalPayments)

	currentPeriod := firstDayOfMonth(rentalContract.StartDate)

	nextAdjustmentPeriod := calculateNextAdjustmentPeriod(rentalContract)

	for i := int64(0); i < rentalContract.TotalPayments; i++ {

		item := domain.RentPaymentScheduleItem{
			RentalContractID: rentalContract.ID,
		}

		if rentPayment, exists := paymentsByPeriod[periodKey(currentPeriod)]; exists {
			item.RentPaymentID = rentPayment.ID
			item.DueDate = rentPayment.DueDate
			item.Period = rentPayment.Period
			item.AppliedAdjustmentPercentage = rentPayment.AppliedAdjustmentPercentage
			item.AppliedInterestAmount = rentPayment.AppliedInterestAmount
			item.TotalAmount = rentPayment.TotalAmount
			item.PaidAmount = rentPayment.PaidAmount
			item.PaymentDate = rentPayment.PaymentDate
			item.IsPaid = rentPayment.IsPaid
		} else {
			if *nextAdjustmentPeriod == currentPeriod {
				calculation, _ := uc.rentPaymentCalculator.Calculate(
					ctx,
					rentalContract,
					currentPeriod,
					time.Now().UTC(),
				)
				item.AppliedAdjustmentPercentage = calculation.AdjustmentPercentage
				item.AppliedInterestAmount = calculation.InterestAmount
				item.SuggestedTotalAmount = calculation.TotalAmount
			}
			previousPeriod := currentPeriod.AddDate(0, -1, 0)
			previousRent, _ := uc.rentPaymentRepository.GetRentPaymentByContractIDAndPeriod(ctx, rentalContract.ID, previousPeriod)
			monthlyAmount := rentalContract.MonthlyAmount
			if previousRent != nil {
				monthlyAmount = previousRent.TotalAmount
			}
			item.DueDate = CalculateDueDate(currentPeriod, rentalContract.DueDay)
			item.Period = currentPeriod
			item.TotalAmount = monthlyAmount
			item.PaidAmount = monthlyAmount
			item.IsPaid = false
		}

		items = append(items, item)
		currentPeriod = currentPeriod.AddDate(0, 1, 0)
	}

	return items, nil
}

func firstDayOfMonth(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func periodKey(value time.Time) string {
	return firstDayOfMonth(value).Format("2006-01")
}
