package router

import (
	"wechat-robot-client/controller"
	"wechat-robot-client/middleware"

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
		robot.POST("/view", middleware.UserOwnerAuth(), robotCtl.RobotView)
		robot.POST("/remove", middleware.UserOwnerAuth(), robotCtl.RobotRemove)
	}

	return nil
}
