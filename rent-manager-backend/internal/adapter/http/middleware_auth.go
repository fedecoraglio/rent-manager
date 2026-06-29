package http

import (
	"strings"

	"github.com/gin-gonic/gin"

	"rent-manager-backend/internal/core/domain"
	"rent-manager-backend/internal/core/port"
)

const (
	authorizationHeader = "Authorization"
	bearerPrefix        = "Bearer "
	authUserIDKey       = "auth_user_id"
	authUserEmailKey    = "auth_user_email"
)

func AuthMiddleware(tokenProvider port.TokenProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authorizationHeader)
		if header == "" {
			handleError(ctx, domain.ErrUnauthorized)
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(header, bearerPrefix) {
			handleError(ctx, domain.ErrUnauthorized)
			ctx.Abort()
			return
		}

		tokenValue := strings.TrimSpace(strings.TrimPrefix(header, bearerPrefix))
		if tokenValue == "" {
			handleError(ctx, domain.ErrUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := tokenProvider.ValidateToken(tokenValue)
		if err != nil {
			handleError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(authUserIDKey, claims.UserID)
		ctx.Set(authUserEmailKey, claims.Email)

		ctx.Next()
	}
}
