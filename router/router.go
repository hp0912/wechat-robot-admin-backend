package router

import (
	"wechat-robot-client/controller"
	"wechat-robot-client/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)

	initController()

	api := r.Group("/api/v1")
	api.GET("/oauth/wechat", wechatCtl.WechatAuth)

	return nil
}
