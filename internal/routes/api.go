package routes

import (
	"net/http"
	"rate_limiter/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GetApiRoutes(r *gin.Engine) {
	h := r.Group("health-check")
	h.Use(middleware.InterceptRequest())
	{
		h.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		})
	}
}
