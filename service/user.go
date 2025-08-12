package service

import (
	"context"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (sv *UserService) SetupLogin(ctx *gin.Context, user *model.User) error {
	session := sessions.Default(ctx)
	session.Set("id", user.ID)
	session.Set("wechat_id", user.WeChatId)
	session.Set("role", user.Role)
	session.Set("status", user.Status)
	return session.Save()
}

func (sv *UserService) Logout(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	return session.Save()
}

func (sv *UserService) LoginUser(id int64) (*model.User, error) {
	return repository.NewUserRepo(sv.ctx, vars.DB).GetUserByID(id)
}

func (sv *UserService) RefreshUserApiToken(id int64) (string, error) {
	apiToken := uuid.New().String()
	err := repository.NewUserRepo(sv.ctx, vars.DB).Update(&model.User{
		ID:       id,
		ApiToken: &apiToken,
	})
	if err != nil {
		return "", err
	}
	return apiToken, nil
}
