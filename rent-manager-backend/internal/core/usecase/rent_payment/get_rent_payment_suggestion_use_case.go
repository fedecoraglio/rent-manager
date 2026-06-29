package rent_payment

import (
	"context"
	"log/slog"
	"time"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

type GetRentPaymentSuggestionUseCase struct {
	rentalContractRepository port.RentalContractRepository
	rentPaymentRepository    port.RentPaymentRepository
	rentPaymentCalculator    port.RentPaymentCalculator
}

func NewGetRentPaymentSuggestionUseCase(
	rentalContractRepository port.RentalContractRepository,
	rentPaymentRepository port.RentPaymentRepository,
	rentPaymentCalculator port.RentPaymentCalculator,
) *GetRentPaymentSuggestionUseCase {
	return &GetRentPaymentSuggestionUseCase{
		rentalContractRepository: rentalContractRepository,
		rentPaymentRepository:    rentPaymentRepository,
		rentPaymentCalculator:    rentPaymentCalculator,
	}
}

func (uc *GetRentPaymentSuggestionUseCase) GetRentPaymentSuggestion(
	ctx context.Context,
	rentalContractID int64,
	period time.Time,
	paymentDate time.Time,
) (*domain.RentPaymentSuggestion, error) {
	rentalContract, err := uc.rentalContractRepository.GetRentalContractByID(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentPaymentSuggestionUseCase] repository get rental contract failed", "err", err)
		return nil, err
	}

	previousPeriod := period.AddDate(0, -1, 0)
	previousRent, errPeriod := uc.rentPaymentRepository.GetRentPaymentByContractIDAndPeriod(ctx, rentalContractID, previousPeriod)
	if errPeriod != nil {
		return nil, errPeriod
	}

	if previousRent != nil && previousRent.TotalAmount != 0 {
		rentalContract.MonthlyAmount = previousRent.TotalAmount
	}

	calculation, err := uc.rentPaymentCalculator.Calculate(
		ctx,
		rentalContract,
		period,
		paymentDate,
	)
	if err != nil {
		slog.Error("[GetRentPaymentSuggestionUseCase] calculator failed", "err", err)
		return nil, err
	}

	dueDate := CalculateDueDate(period, rentalContract.DueDay)

	return &domain.RentPaymentSuggestion{
		RentalContractID:              rentalContract.ID,
		Period:                        period,
		DueDate:                       dueDate,
		PaymentDate:                   paymentDate,
		BaseAmount:                    calculation.BaseAmount,
		SuggestedAdjustmentPercentage: calculation.AdjustmentPercentage,
		SuggestedAdjustmentAmount:     calculation.AdjustmentAmount,
		SuggestedInterestAmount:       calculation.InterestAmount,
		SuggestedTotalAmount:          calculation.TotalAmount,
	}, nil
}

func CalculateDueDate(period time.Time, dueDay int64) time.Time {
	return time.Date(
		period.Year(),
		period.Month(),
		int(dueDay),
		0,
		0,
		0,
		0,
		time.UTC,
	)
}
