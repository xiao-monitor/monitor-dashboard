package auth

import (
	"xiaolfeng/monitor-dashboard/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录
//
// 用于系统用户登录阶段进行登录认证操作使用
//
//	@param ctx *gin.Context 上下文
func (c *AuthController) UserLogin(ctx *gin.Context) {
	utils.SuccessNormal(ctx)
}
