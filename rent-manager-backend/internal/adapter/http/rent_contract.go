package http

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rentalContractUseCase "rent-manager-backend/internal/core/usecase/rental_contract"
)

type RentalContractHandler struct {
	createRentalContractUseCase  *rentalContractUseCase.CreateRentalContractUseCase
	updateRentalContractUseCase  *rentalContractUseCase.UpdateRentalContractUseCase
	getRentalContractByIDUseCase *rentalContractUseCase.GetRentalContractByIDUseCase
	listRentalContractsUseCase   *rentalContractUseCase.ListRentalContractsUseCase
}

func NewRentalContractHandler(
	createRentalContractUseCase *rentalContractUseCase.CreateRentalContractUseCase,
	updateRentalContractUseCase *rentalContractUseCase.UpdateRentalContractUseCase,
	getRentalContractByIDUseCase *rentalContractUseCase.GetRentalContractByIDUseCase,
	listRentalContractsUseCase *rentalContractUseCase.ListRentalContractsUseCase,
) *RentalContractHandler {
	return &RentalContractHandler{
		createRentalContractUseCase:  createRentalContractUseCase,
		updateRentalContractUseCase:  updateRentalContractUseCase,
		getRentalContractByIDUseCase: getRentalContractByIDUseCase,
		listRentalContractsUseCase:   listRentalContractsUseCase,
	}
}

func (h *RentalContractHandler) CreateRentalContract(ctx *gin.Context) {
	var req createRentalContractRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentalContract, err := newRentalContractFromCreateRequest(req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdRentalContract, err := h.createRentalContractUseCase.CreateRentalContract(ctx, rentalContract)
	if err != nil {
		slog.Error("[CreateRentalContract] failed to create rental contract", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentalContractResponse(createdRentalContract))
}

func (h *RentalContractHandler) UpdateRentalContract(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental contract id"})
		return
	}

	var req updateRentalContractRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentalContract, err := newRentalContractFromUpdateRequest(id, req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	updatedRentalContract, err := h.updateRentalContractUseCase.UpdateRentalContract(ctx, rentalContract)
	if err != nil {
		slog.Error("[UpdateRentalContract] failed to update rental contract", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentalContractResponse(updatedRentalContract))
}

func (h *RentalContractHandler) GetRentalContractByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental contract id"})
		return
	}

	rentalContract, err := h.getRentalContractByIDUseCase.GetRentalContractByID(ctx, id)
	if err != nil {
		slog.Error("[GetRentalContractByID] failed to get rental contract by id", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentalContractResponse(rentalContract))
}

func (h *RentalContractHandler) ListRentalContracts(ctx *gin.Context) {
	var req listRentalContractsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rentalContracts, err := h.listRentalContractsUseCase.ListRentalContracts(
		ctx,
		req.PropertyID,
		req.Page,
		req.Limit,
	)
	if err != nil {
		slog.Error("[ListRentalContracts] failed to list rental contracts", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentalContractsResponse(rentalContracts))
}
