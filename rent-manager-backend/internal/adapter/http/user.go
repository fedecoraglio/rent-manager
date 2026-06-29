package http

import (
	"net/http"
	userUseCase "rent-manager-backend/internal/core/usecase/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase     *userUseCase.CreateUserUseCase
	getUserByIDUseCase    *userUseCase.GetUserByIDUseCase
	getUserByEmailUseCase *userUseCase.GetUserByEmailUseCase
	listUsersUseCase      *userUseCase.ListUserUseCase
	searchUsersUseCase    *userUseCase.SearchUserUseCase
	updateUserUseCase     *userUseCase.UpdateUserUseCase
	deleteUserUseCase     *userUseCase.DeleteUserUseCase
}

func NewUserHandler(
	createUserUseCase *userUseCase.CreateUserUseCase,
	getUserByIDUseCase *userUseCase.GetUserByIDUseCase,
	getUserByEmailUseCase *userUseCase.GetUserByEmailUseCase,
	listUsersUseCase *userUseCase.ListUserUseCase,
	searchUsersUseCase *userUseCase.SearchUserUseCase,
	updateUserUseCase *userUseCase.UpdateUserUseCase,
	deleteUserUseCase *userUseCase.DeleteUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUserUseCase:     createUserUseCase,
		getUserByIDUseCase:    getUserByIDUseCase,
		getUserByEmailUseCase: getUserByEmailUseCase,
		listUsersUseCase:      listUsersUseCase,
		searchUsersUseCase:    searchUsersUseCase,
		updateUserUseCase:     updateUserUseCase,
		deleteUserUseCase:     deleteUserUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := newUserFromCreateRequest(req)

	createdUser, err := h.createUserUseCase.CreateUser(ctx, user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUserResponse(createdUser))
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.getUserByIDUseCase.GetUserByID(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUserResponse(user))
}

func (h *UserHandler) ListUsers(ctx *gin.Context) {
	var req paginationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	users, err := h.listUsersUseCase.ListUsers(ctx, req.Page, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUsersResponse(users))
}

func (h *UserHandler) SearchUsers(ctx *gin.Context) {
	var req searchUsersRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	users, err := h.searchUsersUseCase.SearchUsersByNameOrEmail(ctx, req.Value, req.Page, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUsersResponse(users))
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req updateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := newUserFromUpdateRequest(id, req)

	updatedUser, err := h.updateUserUseCase.UpdateUser(ctx, user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUserResponse(updatedUser))
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.deleteUserUseCase.DeleteUser(ctx, id); err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
