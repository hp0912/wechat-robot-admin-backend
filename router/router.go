package router

import (
	"wechat-robot-admin-backend/controller"
	"wechat-robot-admin-backend/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat
var userCtl *controller.User
var robotCtl *controller.Robot

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
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
		robot := api.Group("/robot")
		robot.Use(middleware.UserAuth())

		robot.GET("/list", robotCtl.RobotList)
		robot.POST("/create", robotCtl.RobotCreate)
		robot.GET("/view", middleware.UserOwnerAuth(), robotCtl.RobotView)
		robot.POST("/restart-client", middleware.UserOwnerAuth(), robotCtl.RobotRestartClient)
		robot.POST("/restart-server", middleware.UserOwnerAuth(), robotCtl.RobotRestartServer)
		robot.GET("/state", middleware.UserOwnerAuth(), robotCtl.RobotView)
		robot.DELETE("/remove", middleware.UserOwnerAuth(), robotCtl.RobotRemove)
	}

	return nil
}
