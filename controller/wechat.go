package controller

import (
	"encoding/json"
	"net/http"
	"wechat-robot-client/service"
	"wechat-robot-client/vars"

	"github.com/gin-gonic/gin"
)

type WeChat struct {
}

type WeChatRequest struct {
	Code string `json:"code"`
}

func NewWeChatAuthController() *WeChat {
	return &WeChat{}
}

func (w *WeChat) WechatAuth(c *gin.Context) {
	var req WeChatRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	ctx := c.Request.Context()
	user, err := service.NewWechatService(ctx).WechatAuth(ctx, req.Code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if user.Status != vars.UserStatusEnabled {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": false,
			"message": "用户已被封禁",
		})
		return
	}
	err = service.NewUserService(ctx).SetupLogin(c, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": "",
	})
}
