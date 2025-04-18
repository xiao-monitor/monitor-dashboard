package config

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 初始化日志
func (c *Config) InitLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.DebugLevel)

	// GIN 日志记录输出
	gin.ForceConsoleColor()
	// 创建日志文件(logs/gin/logger.log)
	if _, err := os.Stat(".logs/gin"); os.IsNotExist(err) {
		os.MkdirAll(".logs/gin", 0755)
	}
	// 检查日志是否存在
	if file, _ := os.Open(".logs/gin/latest.log"); file != nil {
		dirName := ".logs/gin/" + time.Now().Format("2006-01-02")
		// 创建一个日期文件夹
		os.MkdirAll(dirName, 0755)
		// 将 latest.log 重命名
		os.Rename(".logs/gin/latest.log", dirName+"/"+strconv.Itoa(int(time.Now().UnixMilli()))+".log")
	}

	f, _ := os.Create(".logs/gin/latest.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
