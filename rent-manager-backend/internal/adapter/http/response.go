package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Data any `json:"data"`
}

func handleSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, successResponse{
		Data: data,
	})
}
