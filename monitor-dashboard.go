package main

import (
	"xiaolfeng/monitor-dashboard/config"
	"xiaolfeng/monitor-dashboard/internal/router"
)

func main() {
	// 初始化日志
	config := config.NewConfig()
	config.InitLogger()
	config.InitDB()

	// 创建路由
	router := router.NewRouter()
	router.SetRouter()
	router.Run(config.Config.System.Port)
}
