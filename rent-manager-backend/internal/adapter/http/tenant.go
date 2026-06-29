package http

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	tenantUseCase "rent-manager-backend/internal/core/usecase/tenant"
)

type TenantHandler struct {
	createTenantUseCase              *tenantUseCase.CreateTenantUseCase
	getTenantByIDUseCase             *tenantUseCase.GetTenantByIDUseCase
	getTenantByDocumentNumberUseCase *tenantUseCase.GetTenantByDocumentNumberUseCase
	listTenantUseCase                *tenantUseCase.ListTenantUseCase
	searchTenantUseCase              *tenantUseCase.SearchTenantsUseCase
	updateTenantUseCase              *tenantUseCase.UpdateTenantUseCase
}

func NewTenantHandler(
	createTenantUseCase *tenantUseCase.CreateTenantUseCase,
	getTenantByIDUseCase *tenantUseCase.GetTenantByIDUseCase,
	getTenantByDocumentNumberUseCase *tenantUseCase.GetTenantByDocumentNumberUseCase,
	listTenantUseCase *tenantUseCase.ListTenantUseCase,
	searchTenantUseCase *tenantUseCase.SearchTenantsUseCase,
	updateTenantUseCase *tenantUseCase.UpdateTenantUseCase,
) *TenantHandler {
	return &TenantHandler{
		createTenantUseCase:              createTenantUseCase,
		getTenantByIDUseCase:             getTenantByIDUseCase,
		getTenantByDocumentNumberUseCase: getTenantByDocumentNumberUseCase,
		listTenantUseCase:                listTenantUseCase,
		searchTenantUseCase:              searchTenantUseCase,
		updateTenantUseCase:              updateTenantUseCase,
	}
}

func (h *TenantHandler) CreateTenant(ctx *gin.Context) {
	var req createTenantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	createdTenant, err := h.createTenantUseCase.CreateTenant(ctx, newTenantFromCreateRequest(req))
	if err != nil {
		slog.Error("[CreateTenant] failed to create tenant", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newTenantResponse(createdTenant))
}

func (h *TenantHandler) GetTenantByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	tenant, err := h.getTenantByIDUseCase.GetTenantByID(ctx, id)
	if err != nil {
		slog.Error("[GetTenantByID] failed to get tenant by id", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newTenantResponse(tenant))
}

func (h *TenantHandler) ListTenants(ctx *gin.Context) {
	var req listTenantsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	tenants, err := h.listTenantUseCase.ListTenants(ctx, req.Page, req.Limit)
	if err != nil {
		slog.Error("[ListTenants] failed to list tenants", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newTenantsResponse(tenants))
}

func (h *TenantHandler) SearchTenants(ctx *gin.Context) {
	var req searchTenantsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	tenants, err := h.searchTenantUseCase.SearchTenants(ctx, req.Value, req.Page, req.Limit)
	if err != nil {
		slog.Error("[SearchTenants] failed to search tenants", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newTenantsResponse(tenants))
}

func (h *TenantHandler) UpdateTenant(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	var req updateTenantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	updatedTenant, err := h.updateTenantUseCase.UpdateTenant(ctx, newTenantFromUpdateRequest(id, req))
	if err != nil {
		slog.Error("[UpdateTenant] failed to update tenant", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newTenantResponse(updatedTenant))
}
