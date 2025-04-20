package service

import (
	"context"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

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

func (u *UserService) Logout(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	return session.Save()
}

func (u *UserService) LoginUser(ctx *gin.Context, id int64) *model.User {
	return repository.NewUserRepo(u.ctx, vars.DB).GetUserByID(id)
}
