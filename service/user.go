package service

import (
	"context"
	"wechat-robot-client/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (u *UserService) SetupLogin(ctx *gin.Context, user *model.User) error {
	session := sessions.Default(ctx)
	session.Set("id", user.Id)
	session.Set("wechat_id", user.WeChatId)
	session.Set("role", user.Role)
	session.Set("status", user.Status)
	return session.Save()
}
