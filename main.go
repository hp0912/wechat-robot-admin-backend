package main

import (
	"log"
	"os"
	"wechat-robot-admin-backend/router"
	"wechat-robot-admin-backend/startup"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// @title WeChat Robot Backend API
// @version 1.0
// @description wechat-robot-admin-backend 服务接口文档
// @BasePath /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey ApiTokenAuth
// @in header
// @name X-API-Token

func main() {
	// 加载配置
	if err := startup.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	if err := startup.SetupVars(); err != nil {
		log.Fatalf("初始化环境失败: %v", err)
	}
	if err := startup.AutoMigrate(); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}
	// 启动HTTP服务
	gin.SetMode(os.Getenv("GIN_MODE"))
	app := gin.Default()
	store := cookie.NewStore([]byte(vars.SessionSecret))
	app.Use(sessions.Sessions("session", store))
	// 注册路由
	if err := router.RegisterRouter(app); err != nil {
		log.Fatalf("注册路由失败: %v", err)
	}
	// 启动服务
	if err := app.Run(":9000"); err != nil {
		log.Panicf("服务启动失败：%v", err)
	}
}
