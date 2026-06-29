package http

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rentPaymentUseCase "rent-manager-backend/internal/core/usecase/rent_payment"
)

type RentPaymentHandler struct {
	getRentPaymentScheduleUseCase   *rentPaymentUseCase.GetRentPaymentScheduleUseCase
	createRentPaymentUseCase        *rentPaymentUseCase.CreateRentPaymentUseCase
	updateRentPaymentUseCase        *rentPaymentUseCase.UpdateRentPaymentUseCase
	getRentPaymentByIDUseCase       *rentPaymentUseCase.GetRentPaymentByIDUseCase
	listRentPaymentsUseCase         *rentPaymentUseCase.ListRentPaymentsUseCase
	getRentPaymentSuggestionUseCase *rentPaymentUseCase.GetRentPaymentSuggestionUseCase
	getRentalContractSummaryUseCase *rentPaymentUseCase.GetRentalContractSummaryUseCase
}

func NewRentPaymentHandler(
	getRentPaymentScheduleUseCase *rentPaymentUseCase.GetRentPaymentScheduleUseCase,
	createRentPaymentUseCase *rentPaymentUseCase.CreateRentPaymentUseCase,
	updateRentPaymentUseCase *rentPaymentUseCase.UpdateRentPaymentUseCase,
	getRentPaymentByIDUseCase *rentPaymentUseCase.GetRentPaymentByIDUseCase,
	listRentPaymentsUseCase *rentPaymentUseCase.ListRentPaymentsUseCase,
	getRentPaymentSuggestionUseCase *rentPaymentUseCase.GetRentPaymentSuggestionUseCase,
	getRentalContractSummaryUseCase *rentPaymentUseCase.GetRentalContractSummaryUseCase,
) *RentPaymentHandler {
	return &RentPaymentHandler{
		getRentPaymentScheduleUseCase:   getRentPaymentScheduleUseCase,
		getRentPaymentSuggestionUseCase: getRentPaymentSuggestionUseCase,
		createRentPaymentUseCase:        createRentPaymentUseCase,
		updateRentPaymentUseCase:        updateRentPaymentUseCase,
		getRentPaymentByIDUseCase:       getRentPaymentByIDUseCase,
		listRentPaymentsUseCase:         listRentPaymentsUseCase,
		getRentalContractSummaryUseCase: getRentalContractSummaryUseCase,
	}
}

func (h *RentPaymentHandler) GetRentPaymentSchedule(ctx *gin.Context) {
	rentalContractID, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental contract id"})
		return
	}

	schedule, err := h.getRentPaymentScheduleUseCase.GetRentPaymentSchedule(ctx, rentalContractID)
	if err != nil {
		slog.Error("[GetRentPaymentSchedule] failed to get rent payment schedule", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentScheduleResponse(schedule))
}

func (h *RentPaymentHandler) CreateRentPayment(ctx *gin.Context) {
	var req createRentPaymentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentPayment, err := newRentPaymentFromCreateRequest(req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdRentPayment, err := h.createRentPaymentUseCase.CreateRentPayment(ctx, rentPayment)
	if err != nil {
		slog.Error("[CreateRentPayment] failed to create rent payment", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentResponse(createdRentPayment))
}

func (h *RentPaymentHandler) UpdateRentPayment(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rent payment id"})
		return
	}

	var req updateRentPaymentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentPayment, err := newRentPaymentFromUpdateRequest(id, req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	updatedRentPayment, err := h.updateRentPaymentUseCase.UpdateRentPayment(ctx, rentPayment)
	if err != nil {
		slog.Error("[UpdateRentPayment] failed to update rent payment", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentResponse(updatedRentPayment))
}

func (h *RentPaymentHandler) GetRentPaymentByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rent payment id"})
		return
	}

	rentPayment, err := h.getRentPaymentByIDUseCase.GetRentPaymentByID(ctx, id)
	if err != nil {
		slog.Error("[GetRentPaymentByID] failed to get rent payment by id", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentResponse(rentPayment))
}

func (h *RentPaymentHandler) ListRentPayments(ctx *gin.Context) {
	var req listRentPaymentsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentPayments, err := h.listRentPaymentsUseCase.ListRentPayments(
		ctx,
		req.RentalContractID,
		req.Page,
		req.Limit,
	)
	if err != nil {
		slog.Error("[ListRentPayments] failed to list rent payments", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentsResponse(rentPayments))
}

func (h *RentPaymentHandler) GetRentPaymentSuggestion(ctx *gin.Context) {
	rentalContractID, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental contract id"})
		return
	}

	var req getRentPaymentSuggestionRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	period, err := parsePeriod(req.Period)
	if err != nil {
		validationError(ctx, err)
		return
	}

	paymentDate, err := parseDate(req.PaymentDate)
	if err != nil {
		validationError(ctx, err)
		return
	}

	suggestion, err := h.getRentPaymentSuggestionUseCase.GetRentPaymentSuggestion(
		ctx,
		rentalContractID,
		period,
		paymentDate,
	)
	if err != nil {
		slog.Error("[GetRentPaymentSuggestion] failed to get rent payment suggestion", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentPaymentSuggestionResponse(suggestion))
}

func (h *RentPaymentHandler) GetRentalContractSummary(ctx *gin.Context) {
	rentalContractID, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental contract id"})
		return
	}

	summary, err := h.getRentalContractSummaryUseCase.GetRentalContractSummary(
		ctx,
		rentalContractID,
	)
	if err != nil {
		slog.Error("[GetRentalContractSummary] failed to get rental contract summary", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentalContractSummaryResponse(summary))
}
