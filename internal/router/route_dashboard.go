package router

import "github.com/gin-gonic/gin"

// dashboardRoute 仪表盘路由
func (r *Router) dashboardRoute(api *gin.RouterGroup) {
	api.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}