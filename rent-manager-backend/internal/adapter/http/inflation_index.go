package http

import (
	"log/slog"
	_ "net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/usecase/inflation_index"
)

type InflationIndexHandler struct {
	createInflationIndexUseCase  *inflation_index.CreateInflationIndexUseCase
	updateInflationIndexUseCase  *inflation_index.UpdateInflationIndexUseCase
	getInflationIndexByIDUseCase *inflation_index.GetInflationIndexByIDUseCase
	listInflationIndexesUseCase  *inflation_index.ListInflationIndexesUseCase
}

func NewInflationIndexHandler(
	createInflationIndexUseCase *inflation_index.CreateInflationIndexUseCase,
	updateInflationIndexUseCase *inflation_index.UpdateInflationIndexUseCase,
	getInflationIndexByIDUseCase *inflation_index.GetInflationIndexByIDUseCase,
	listInflationIndexesUseCase *inflation_index.ListInflationIndexesUseCase,
) *InflationIndexHandler {
	return &InflationIndexHandler{
		createInflationIndexUseCase:  createInflationIndexUseCase,
		updateInflationIndexUseCase:  updateInflationIndexUseCase,
		getInflationIndexByIDUseCase: getInflationIndexByIDUseCase,
		listInflationIndexesUseCase:  listInflationIndexesUseCase,
	}
}

func (h *InflationIndexHandler) Create(ctx *gin.Context) {
	var request InflationIndexCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("[Create] invalid binding data", "err", err)
		handleError(ctx, err)
		return
	}

	period, err := time.Parse("2006-01", request.Period)
	if err != nil {
		slog.Error("[Create] invalid period format, expected yyyy-mm", "err", err)
		handleError(ctx, err)
		return
	}

	inflationIndex := &domain.InflationIndex{
		Period:     period,
		Percentage: request.Percentage,
		Source:     request.Source,
		Notes:      request.Notes,
	}

	created, err := h.createInflationIndexUseCase.Execute(ctx.Request.Context(), inflationIndex)
	if err != nil {
		slog.Error("[Create] error creation inflation index", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newInflationIndexResponse(created))

}

func (h *InflationIndexHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		slog.Error("[Update] invalid id inflation index", "err", err)
		handleError(ctx, err)
		return
	}

	var request InflationIndexUpdateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("[Update] invalid binding json data", "err", err)
		handleError(ctx, err)
		return
	}

	period, err := time.Parse("2006-01", request.Period)
	if err != nil {
		slog.Error("[Update] invalid period format, expected yyyy-mm", "err", err)
		handleError(ctx, err)
		return
	}

	inflationIndex := &domain.InflationIndex{
		ID:         id,
		Period:     period,
		Percentage: request.Percentage,
		Source:     request.Source,
		Notes:      request.Notes,
	}

	updated, err := h.updateInflationIndexUseCase.Execute(ctx.Request.Context(), inflationIndex)
	if err != nil {
		slog.Error("[Update] error saving inflation index", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newInflationIndexResponse(updated))
}

func (h *InflationIndexHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		slog.Error("[GetByID] failed to parse inflation index", "err", err)
		handleError(ctx, err)
		return
	}

	inflationIndex, err := h.getInflationIndexByIDUseCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		slog.Error("[GetByID] failed to get inflation index", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newInflationIndexResponse(inflationIndex))
}

func (h *InflationIndexHandler) List(ctx *gin.Context) {
	page := uint64(1)
	limit := uint64(10)

	if value := ctx.Query("page"); value != "" {
		parsed, err := strconv.ParseUint(value, 10, 64)
		if err == nil && parsed > 0 {
			page = parsed
		}
	}

	if value := ctx.Query("limit"); value != "" {
		parsed, err := strconv.ParseUint(value, 10, 64)
		if err == nil && parsed > 0 {
			limit = parsed
		}
	}

	inflationIndexes, err := h.listInflationIndexesUseCase.Execute(
		ctx.Request.Context(),
		page,
		limit,
	)
	if err != nil {
		slog.Error("[ListInflationIndex] failed to list inflation indexes", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newInflationIndexesResponse(inflationIndexes))
}
