package router

import (
	"wechat-robot-client/controller"
	"wechat-robot-client/middleware"

	"github.com/gin-gonic/gin"
)

var robotCtl *controller.Robot

func initController() {
	robotCtl = controller.NewRobotController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)

	initController()

	api := r.Group("/api/v1")
	api.POST("/test", robotCtl.Test)

	return nil
}
