package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	contractCatalogUseCase "rent-manager-backend/internal/core/usecase/contract_catalog"
)

type ContractCatalogHandler struct {
	listContractStatusesUseCase         *contractCatalogUseCase.ListContractStatusesUseCase
	listInterestCalculationTypesUseCase *contractCatalogUseCase.ListInterestCalculationTypesUseCase
	listRentAdjustmentTypesUseCase      *contractCatalogUseCase.ListRentAdjustmentTypesUseCase
}

func NewContractCatalogHandler(
	listContractStatusesUseCase *contractCatalogUseCase.ListContractStatusesUseCase,
	listInterestCalculationTypesUseCase *contractCatalogUseCase.ListInterestCalculationTypesUseCase,
	listRentAdjustmentTypesUseCase *contractCatalogUseCase.ListRentAdjustmentTypesUseCase,
) *ContractCatalogHandler {
	return &ContractCatalogHandler{
		listContractStatusesUseCase:         listContractStatusesUseCase,
		listInterestCalculationTypesUseCase: listInterestCalculationTypesUseCase,
		listRentAdjustmentTypesUseCase:      listRentAdjustmentTypesUseCase,
	}
}

func (h *ContractCatalogHandler) ListContractStatuses(ctx *gin.Context) {
	statuses, err := h.listContractStatusesUseCase.ListContractStatuses(ctx)
	if err != nil {
		slog.Error("[ListContractStatuses] failed to list contract statuses", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newContractStatusesResponse(statuses))
}

func (h *ContractCatalogHandler) ListInterestCalculationTypes(ctx *gin.Context) {
	interestTypes, err := h.listInterestCalculationTypesUseCase.ListInterestCalculationTypes(ctx)
	if err != nil {
		slog.Error("[ListInterestCalculationTypes] failed to list interest calculation types", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newInterestCalculationTypesResponse(interestTypes))
}

func (h *ContractCatalogHandler) ListRentAdjustmentTypes(ctx *gin.Context) {
	adjustmentTypes, err := h.listRentAdjustmentTypesUseCase.ListRentAdjustmentTypes(ctx)
	if err != nil {
		slog.Error("[ListRentAdjustmentTypes] failed to list rent adjustment types", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newRentAdjustmentTypesResponse(adjustmentTypes))
}
