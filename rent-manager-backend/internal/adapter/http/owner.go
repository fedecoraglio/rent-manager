package http

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ownerUseCase "rent-manager-backend/internal/core/usecase/owner"
)

type OwnerHandler struct {
	createOwnerUseCase              *ownerUseCase.CreateOwnerUseCase
	getOwnerByIDUseCase             *ownerUseCase.GetOwnerByIDUseCase
	getOwnerByDocumentNumberUseCase *ownerUseCase.GetOwnerByDocumentNumberUseCase
	listOwnersUseCase               *ownerUseCase.ListOwnersUseCase
	searchOwnersUseCase             *ownerUseCase.SearchOwnersUseCase
	updateOwnerUseCase              *ownerUseCase.UpdateOwnerUseCase
}

func NewOwnerHandler(
	createOwnerUseCase *ownerUseCase.CreateOwnerUseCase,
	getOwnerByIDUseCase *ownerUseCase.GetOwnerByIDUseCase,
	getOwnerByDocumentNumberUseCase *ownerUseCase.GetOwnerByDocumentNumberUseCase,
	listOwnersUseCase *ownerUseCase.ListOwnersUseCase,
	searchOwnersUseCase *ownerUseCase.SearchOwnersUseCase,
	updateOwnerUseCase *ownerUseCase.UpdateOwnerUseCase,
) *OwnerHandler {
	return &OwnerHandler{
		createOwnerUseCase:              createOwnerUseCase,
		getOwnerByIDUseCase:             getOwnerByIDUseCase,
		getOwnerByDocumentNumberUseCase: getOwnerByDocumentNumberUseCase,
		listOwnersUseCase:               listOwnersUseCase,
		searchOwnersUseCase:             searchOwnersUseCase,
		updateOwnerUseCase:              updateOwnerUseCase,
	}
}

func (h *OwnerHandler) CreateOwner(ctx *gin.Context) {
	var req createOwnerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	createdOwner, err := h.createOwnerUseCase.CreateOwner(
		ctx,
		newOwnerFromCreateRequest(req),
	)
	if err != nil {
		slog.Error("[CreateOwner] failed to create owner", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newOwnerResponse(createdOwner))
}

func (h *OwnerHandler) GetOwnerByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	owner, err := h.getOwnerByIDUseCase.GetOwnerByID(ctx, id)
	if err != nil {
		slog.Error("[GetOwnerByID] failed to get owner by id", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newOwnerResponse(owner))
}

func (h *OwnerHandler) ListOwners(ctx *gin.Context) {
	var req listOwnersRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	owners, err := h.listOwnersUseCase.ListOwners(ctx, req.Page, req.Limit)
	if err != nil {
		slog.Error("[ListOwners] failed to list owners", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newOwnersResponse(owners))
}

func (h *OwnerHandler) SearchOwners(ctx *gin.Context) {
	var req searchOwnersRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	owners, err := h.searchOwnersUseCase.SearchOwners(ctx, req.Value, req.Page, req.Limit)
	if err != nil {
		slog.Error("[SearchOwners] failed to search owners", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newOwnersResponse(owners))
}

func (h *OwnerHandler) UpdateOwner(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
		return
	}

	var req updateOwnerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	updatedOwner, err := h.updateOwnerUseCase.UpdateOwner(
		ctx,
		newOwnerFromUpdateRequest(id, req),
	)
	if err != nil {
		slog.Error("[UpdateOwner] failed to update owner", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newOwnerResponse(updatedOwner))
}
