package router

import (
	"xiaolfeng/monitor-dashboard/internal/controller"

	"github.com/gin-gonic/gin"
)

// dashboardRoute 仪表盘路由
func (r *Router) dashboardRoute(api *gin.RouterGroup) {
	dashboardController := controller.NewDashboardController()
	
	dashboard := api.Group("/dashboard")
	dashboard.GET("/", dashboardController.GetDashboard)
}
