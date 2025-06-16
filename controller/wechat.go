package controller

import (
	"errors"
	"fmt"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-gonic/gin"
)

type WeChat struct {
}

type WeChatRequest struct {
	Code string `form:"code" json:"code" binding:"required"`
}

func NewWeChatAuthController() *WeChat {
	return &WeChat{}
}

func (w *WeChat) WeChatOfficialAccountAuthURL(c *gin.Context) {
	ctx := c.Request.Context()
	imageData, contentType, err := service.NewWeChatService(ctx).WeChatOfficialAccountAuthURL(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 设置响应头
	c.Header("Content-Type", contentType)
	c.Header("Cache-Control", "public, max-age=3600") // 缓存1小时
	c.Header("Content-Length", fmt.Sprintf("%d", len(imageData)))
	// 直接返回图片数据
	c.Data(200, contentType, imageData)
}

func (w *WeChat) WeChatAuth(c *gin.Context) {
	req := &WeChatRequest{}
	resp := appx.NewResponse(c)
	ok, validErrs := appx.BindAndValid(c, req)
	if !ok {
		resp.ToInvalidResponse(validErrs)
		return
	}
	ctx := c.Request.Context()
	user, err := service.NewWeChatService(ctx).WeChatAuth(ctx, req.Code)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	if user.Status != vars.UserStatusEnabled {
		resp.ToErrorResponse(errors.New("用户已被封禁"))
		return
	}
	err = service.NewUserService(ctx).SetupLogin(c, user)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(gin.H{
		"success": true,
	})
}
