package controller

import (
	"errors"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUserController() *User {
	return &User{}
}

func (ct *User) Login(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	ctx := c.Request.Context()
	user, err := service.NewUserService(ctx).Login(ctx, req.Token)
	if err != nil {
		resp.ToErrorResponse(err)
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

func (ct *User) LoginUser(c *gin.Context) {
	resp := appx.NewResponse(c)

	defer func() {
		if err := recover(); err != nil {
			user := &model.User{
				LoginMethod: model.LoginMethod(vars.LoginMethod),
			}
			c.SetCookie("session", "", -1, "/", "", false, true)
			resp.To401ResponseWithData(user, errors.New("登陆信息已失效"))
			return
		}
	}()

	session := sessions.Default(c)
	id := session.Get("id")
	if value, ok := id.(int64); ok {
		user, err := service.NewUserService(c.Request.Context()).LoginUser(value)
		if err != nil {
			resp.ToErrorResponse(err)
			return
		}
		if user == nil {
			resp.ToErrorResponse(errors.New("用户不存在"))
			return
		}
		resp.ToResponse(user)
		return
	} else {
		user := &model.User{
			LoginMethod: model.LoginMethod(vars.LoginMethod),
		}
		resp.To401ResponseWithData(user, errors.New("登陆信息已失效"))
		return
	}
}

func (ct *User) RefreshUserApiToken(c *gin.Context) {
	resp := appx.NewResponse(c)
	user, err := appx.GetLoginUser(c)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	apiToken, err := service.NewUserService(c).RefreshUserApiToken(user.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(apiToken)
}

func (ct *User) Logout(c *gin.Context) {
	resp := appx.NewResponse(c)
	err := service.NewUserService(c.Request.Context()).Logout(c)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	user := &model.User{
		LoginMethod: model.LoginMethod(vars.LoginMethod),
	}
	resp.ToResponse(user)
}
