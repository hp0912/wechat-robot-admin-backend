package controller

import (
	"net/http"
	"wechat-robot-client/service"
	"wechat-robot-client/vars"

	"github.com/gin-gonic/gin"
)

type WeChat struct {
}

func NewWeChatAuthController() *WeChat {
	return &WeChat{}
}

func (w *WeChat) WechatAuth(c *gin.Context) {
	code := c.Query("code")
	ctx := c.Request.Context()
	user, err := service.NewWechatService(ctx).WechatAuth(ctx, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if user.Status != vars.UserStatusEnabled {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户已被封禁",
		})
		return
	}
	err = service.NewUserService(ctx).SetupLogin(c, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}
