package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	propertyCatalogUseCase "rent-manager-backend/internal/core/usecase/property_catalog"
)

type PropertyCatalogHandler struct {
	listPropertyTypesUseCase    *propertyCatalogUseCase.ListPropertyTypesUseCase
	listPropertyStatusesUseCase *propertyCatalogUseCase.ListPropertyStatusesUseCase
}

func NewPropertyCatalogHandler(
	listPropertyTypesUseCase *propertyCatalogUseCase.ListPropertyTypesUseCase,
	listPropertyStatusesUseCase *propertyCatalogUseCase.ListPropertyStatusesUseCase,
) *PropertyCatalogHandler {
	return &PropertyCatalogHandler{
		listPropertyTypesUseCase:    listPropertyTypesUseCase,
		listPropertyStatusesUseCase: listPropertyStatusesUseCase,
	}
}

func (h *PropertyCatalogHandler) ListPropertyTypes(ctx *gin.Context) {
	propertyTypes, err := h.listPropertyTypesUseCase.ListPropertyTypes(ctx)
	if err != nil {
		slog.Error("[ListPropertyTypes] failed to list property types", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertyTypesResponse(propertyTypes))
}

func (h *PropertyCatalogHandler) ListPropertyStatuses(ctx *gin.Context) {
	propertyStatuses, err := h.listPropertyStatusesUseCase.ListPropertyStatuses(ctx)
	if err != nil {
		slog.Error("[ListPropertyStatuses] failed to list property statuses", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newPropertyStatusesResponse(propertyStatuses))
}
