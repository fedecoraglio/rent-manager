package http

import (
	"github.com/gin-gonic/gin"
	"rent-manager-backend/internal/core/usecase/role"
)

type RoleHandler struct {
	listRolesUseCase *role.ListRolesUseCase
}

func NewRoleHandler(listRolesUseCase *role.ListRolesUseCase) *RoleHandler {
	return &RoleHandler{listRolesUseCase: listRolesUseCase}
}

func (rh *RoleHandler) ListRoles(ctx *gin.Context) {
	var req paginationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	roles, err := rh.listRolesUseCase.ListRoles(ctx, req.Page, req.Limit)
	if err != nil {
		handleError(ctx, err)
	}

	handleSuccess(ctx, newRolesResponse(roles))
}
