package config

import (
	"os"
	"xiaolfeng/monitor-dashboard/internal/model/other"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	Config *other.Config
}

// NewConfig 初始化配置
func NewConfig() *Config {
	// 从当前目录解析配置文件
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		yamlFile, err = os.ReadFile("config.yaml")
		if err != nil {
			// 如果文件不存在，使用 config.Database 的默认配置
			defaultConfig := other.Config{
				Database: other.Database{
					Host:     "localhost",
					Port:     5432,
					Username: "postgres",
					Password: "postgres",
					DBName:   "monitor_dashboard",
					SSLMode:  "disable",
				},
			}
			yamlFile, err = yaml.Marshal(defaultConfig)
			if err != nil {
				logrus.Panicf("无法生成默认配置文件: %v", err)
			}
			os.WriteFile("config.yaml", yamlFile, 0644)
			// 退出程序
			logrus.Debug("配置文件不存在，已生成默认配置文件")
			os.Exit(1)
		}
	}
	var config other.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logrus.Panicf("无法解析配置文件: %v", err)
	}

	return &Config{
		Config: &config,
	}
}
