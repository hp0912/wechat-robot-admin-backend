package router

import (
	"wechat-robot-admin-backend/controller"
	"wechat-robot-admin-backend/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat
var userCtl *controller.User
var robotCtl *controller.Robot
var systemCtl *controller.System

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
	userCtl = controller.NewUserController()
	robotCtl = controller.NewRobotController()
	systemCtl = controller.NewSystemController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)

	initController()

	api := r.Group("/api/v1")
	{
		oauth := api.Group("/oauth")
		oauth.POST("/wechat", wechatCtl.WeChatAuth)
	}

	{
		user := api.Group("/user")
		user.GET("/self", userCtl.LoginUser)
		user.DELETE("/logout", userCtl.Logout)
	}

	{
		system := api.Group("/system")
		system.Use(middleware.UserAuth())
		system.GET("/robot-container-stats", middleware.UserOwnerAuth(), systemCtl.RobotContainerStats)
	}

	{
		robot := api.Group("/robot")
		robot.Use(middleware.UserAuth())

		robot.GET("/list", robotCtl.RobotList)
		robot.POST("/create", robotCtl.RobotCreate)
		robot.GET("/view", middleware.UserOwnerAuth(), robotCtl.RobotView)
		robot.POST("/restart-client", middleware.UserOwnerAuth(), robotCtl.RobotRestartClient)
		robot.POST("/restart-server", middleware.UserOwnerAuth(), robotCtl.RobotRestartServer)
		robot.DELETE("/remove", middleware.UserOwnerAuth(), robotCtl.RobotRemove)
		// 机器人登陆、登出
		robot.GET("/state", middleware.UserOwnerAuth(), robotCtl.RobotState)
		robot.POST("/login", middleware.UserOwnerAuth(), robotCtl.RobotLogin)
		robot.POST("/login-check", middleware.UserOwnerAuth(), robotCtl.RobotLoginCheck)
		robot.DELETE("/logout", middleware.UserOwnerAuth(), robotCtl.RobotLogout)
	}

	return nil
}
