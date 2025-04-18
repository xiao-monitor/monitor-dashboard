package controller

import (
	"xiaolfeng/monitor-dashboard/internal/utils"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

// NewDashboardController 创建DashboardController实例
//
//	@return *DashboardController
func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

// GetDashboard 获取仪表盘数据
//
//	@param ctx *gin.Context 上下文
func (c *DashboardController) GetDashboard(ctx *gin.Context) {
	utils.SuccessNormal(ctx)
}
