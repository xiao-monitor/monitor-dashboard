package router

import "github.com/gin-gonic/gin"

// userRoute 用户路由
func (r *Router) userRoute() {
	r.engine.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}