package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB 初始化数据库
func (c *Config) InitDB() {
	dbConfig := c.Config.Database

	// 连接数据库(postgres)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("无法连接到数据库: %v", err)
	}

	// db.AutoMigrate(&other.Config{})

	c.DB = db
}
