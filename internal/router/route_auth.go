package router

import (
	"xiaolfeng/monitor-dashboard/internal/controller/auth"

	"github.com/gin-gonic/gin"
)

func (r *Router) authRoute(api *gin.RouterGroup) {
	authController := auth.NewAuthController()
	
	auth := api.Group("/auth")
	auth.POST("/login", authController.UserLogin)
}
