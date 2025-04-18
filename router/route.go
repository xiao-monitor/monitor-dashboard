package router

import (
	"github.com/gin-gonic/gin"
)

// Router 路由
type Router struct {
	engine *gin.Engine
}

// NewRouter 创建路由
func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

// SetRouter 设置路由
func (r *Router) SetRouter() {
	r.dashboardRoute()
	r.userRoute()
}

// Run 启动路由
func (r *Router) Run() {
	r.engine.Run(":8080")
}
