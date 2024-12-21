package main

import (
	"github.com/spf13/viper"
	"ticket_sys/configs"
	"ticket_sys/pkg/logger"
)

func initLogger() {
	// 从配置中读取日志级别
	logLevel := viper.GetString("logger.level")
	logger.InitLogger(logLevel)
}

func initServices() {
	// 这里初始化其他服务，例如数据库连接、微服务客户端等
	logger.Info("Initializing services...")

	// example: Initialize a database connection
	// db.InitConnection(viper.GetString("database.url"))
}

func main() {
	// 初始化配置
	configs.InitConfig()
	// 初始化日志
	initLogger()
	defer logger.Sync()

	// 初始化其他服务
	initServices()

	// 使用配置
	appName := viper.GetString("app.name")
	logger.Info("Application started", "name", appName)

	// 剩余程序逻辑...
}
