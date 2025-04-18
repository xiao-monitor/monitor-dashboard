package router

import (
	"strconv"

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
	api := r.engine.Group("/api/v1")
	{
		r.authRoute(api)
		r.dashboardRoute(api)
		r.userRoute(api)
	}
}

// Run 启动路由
func (r *Router) Run(port uint16) {
	if port == 0 {
		port = 8888
	}
	r.engine.Run(":" + strconv.FormatUint(uint64(port), 10))
}
