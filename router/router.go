package router

import (
	"wechat-robot-client/controller"
	"wechat-robot-client/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat
var userCtl *controller.User

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)

	initController()

	api := r.Group("/api/v1")
	{
		oauth := api.Group("/oauth")
		oauth.POST("/wechat", wechatCtl.WechatAuth)
	}

	{
		user := api.Group("/user")
		user.GET("/self", userCtl.LoginUser)
		user.DELETE("/logout", userCtl.Logout)
	}

	return nil
}
