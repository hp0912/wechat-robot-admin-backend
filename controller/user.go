package controller

import (
	"errors"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUserController() *User {
	return &User{}
}

func (ct *User) LoginUser(c *gin.Context) {
	resp := appx.NewResponse(c)
	session := sessions.Default(c)
	id := session.Get("id")
	if value, ok := id.(int64); ok {
		user, err := service.NewUserService(c.Request.Context()).LoginUser(c, value)
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
		resp.To401Response(errors.New("登陆信息已失效"))
		return
	}
}

func (ct *User) Logout(c *gin.Context) {
	resp := appx.NewResponse(c)
	err := service.NewUserService(c.Request.Context()).Logout(c)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
	return
}
