package controller

import (
	"errors"
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
	service.NewWeChatService(ctx).WeChatOfficialAccountAuthURL(ctx)
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
