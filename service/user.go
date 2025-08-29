package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
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

func (sv *UserService) Login(ctx context.Context, token string) (*model.User, error) {
	if vars.LoginMethod != string(model.LoginMethodToken) {
		return nil, errors.New("登录方式不合法")
	}
	if vars.LoginToken == "" {
		return nil, errors.New("未配置登录密钥，请联系管理员")
	}
	if token != vars.LoginToken {
		return nil, errors.New("登录密钥不正确")
	}
	userRespo := repository.NewUserRepo(ctx, vars.DB)
	user, err := userRespo.GetUser()
	if err != nil {
		return nil, err
	}
	if user == nil {
		wechatId := uuid.New().String()
		user = &model.User{
			WeChatId:    strings.ReplaceAll(wechatId, "-", "")[:28],
			DisplayName: "超级管理员",
			Role:        vars.RoleRootUser,
			Status:      vars.UserStatusEnabled,
			AvatarUrl:   vars.UserDefaultAvatar,
			LastLoginAt: time.Now().Unix(),
			CreatedAt:   time.Now().Unix(),
		}
		err = userRespo.Create(user)
		if err != nil {
			return nil, fmt.Errorf("登录失败，请联系管理员: %w", err)
		}
	} else {
		user.LastLoginAt = time.Now().Unix()
		err = userRespo.Update(user)
		if err != nil {
			return nil, fmt.Errorf("登录失败，请联系管理员: %w", err)
		}
	}
	return user, nil
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
