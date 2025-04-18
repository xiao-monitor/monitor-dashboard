package config

import (
	"errors"
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
	config := prepareConfig()
	if _, err := checkDataIsCorrect(config); err != nil {
		logrus.Panicf("配置文件数据不正确: %v", err.Error())
	}

	return &Config{
		Config: config,
	}
}

// prepareConfig 准备配置
func prepareConfig() *other.Config {
	// 从当前目录解析配置文件
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		yamlFile, err = os.ReadFile("config.yaml")
		if err != nil {
			// 如果文件不存在，使用 config.Database 的默认配置
			defaultConfig := other.Config{
				System: other.System{
					Port: 8888,
					Mode: "debug",
				},
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
	return &config
}

// checkDataIsCorrect 检查配置文件数据是否正确
func checkDataIsCorrect(config *other.Config) (bool, error) {
	if config.System.Port <= 1000 {
		return false, errors.New("系统端口不能小于1000")
	}
	if config.System.Mode != "debug" && config.System.Mode != "release" {
		return false, errors.New("系统模式只能为debug或release")
	}
	return true, nil
}
