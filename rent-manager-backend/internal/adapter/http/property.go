package http

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	propertyUseCase "rent-manager-backend/internal/core/usecase/property"
)

type PropertyHandler struct {
	createPropertyUseCase          *propertyUseCase.CreatePropertyUseCase
	getPropertyByIDUseCase         *propertyUseCase.GetPropertyByIDUseCase
	getPropertyByCodeUseCase       *propertyUseCase.GetPropertyByCodeUseCase
	listPropertiesUseCase          *propertyUseCase.ListPropertiesUseCase
	listPropertiesSummariesUseCase *propertyUseCase.ListPropertiesSummariesUseCase
	searchPropertiesUseCase        *propertyUseCase.SearchPropertiesUseCase
	updatePropertyUseCase          *propertyUseCase.UpdatePropertyUseCase
}

func NewPropertyHandler(
	createPropertyUseCase *propertyUseCase.CreatePropertyUseCase,
	getPropertyByIDUseCase *propertyUseCase.GetPropertyByIDUseCase,
	getPropertyByCodeUseCase *propertyUseCase.GetPropertyByCodeUseCase,
	listPropertiesUseCase *propertyUseCase.ListPropertiesUseCase,
	listPropertiesSummariesUseCase *propertyUseCase.ListPropertiesSummariesUseCase,
	searchPropertiesUseCase *propertyUseCase.SearchPropertiesUseCase,
	updatePropertyUseCase *propertyUseCase.UpdatePropertyUseCase,
) *PropertyHandler {
	return &PropertyHandler{
		createPropertyUseCase:          createPropertyUseCase,
		getPropertyByIDUseCase:         getPropertyByIDUseCase,
		getPropertyByCodeUseCase:       getPropertyByCodeUseCase,
		listPropertiesUseCase:          listPropertiesUseCase,
		listPropertiesSummariesUseCase: listPropertiesSummariesUseCase,
		searchPropertiesUseCase:        searchPropertiesUseCase,
		updatePropertyUseCase:          updatePropertyUseCase,
	}
}

func (h *PropertyHandler) CreateProperty(ctx *gin.Context) {
	var req createPropertyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	createdProperty, err := h.createPropertyUseCase.CreateProperty(
		ctx,
		newPropertyFromCreateRequest(req),
	)
	if err != nil {
		slog.Error("[CreateProperty] failed to create property", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertyResponse(createdProperty))
}

func (h *PropertyHandler) GetPropertyByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
		return
	}

	property, err := h.getPropertyByIDUseCase.GetPropertyByID(ctx, id)
	if err != nil {
		slog.Error("[GetPropertyByID] failed to get property by id", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertyResponse(property))
}

func (h *PropertyHandler) ListProperties(ctx *gin.Context) {
	var req listPropertiesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	properties, err := h.listPropertiesUseCase.ListProperties(ctx, req.Page, req.Limit)
	if err != nil {
		slog.Error("[ListProperties] failed to list properties", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertiesResponse(properties))
}

func (h *PropertyHandler) ListPropertiesSummary(ctx *gin.Context) {
	var req listPropertiesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	properties, err := h.listPropertiesSummariesUseCase.Execute(ctx, req.Page, req.Limit)
	if err != nil {
		slog.Error("[ListProperties] failed to list properties", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertiesSummaryResponse(properties))
}

func (h *PropertyHandler) SearchProperties(ctx *gin.Context) {
	var req searchPropertiesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	properties, err := h.searchPropertiesUseCase.SearchProperties(ctx, req.Value, req.Page, req.Limit)
	if err != nil {
		slog.Error("[SearchProperties] failed to search properties", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertiesResponse(properties))
}

func (h *PropertyHandler) UpdateProperty(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
		return
	}

	var req updatePropertyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	updatedProperty, err := h.updatePropertyUseCase.UpdateProperty(
		ctx,
		newPropertyFromUpdateRequest(id, req),
	)
	if err != nil {
		slog.Error("[UpdateProperty] failed to update property", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertyResponse(updatedProperty))
}
