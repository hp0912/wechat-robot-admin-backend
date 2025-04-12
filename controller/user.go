package controller

import (
	"net/http"
	"wechat-robot-client/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUserController() *User {
	return &User{}
}

func (w *User) LoginUser(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	if value, ok := id.(int64); ok {
		user := service.NewUserService(c.Request.Context()).LoginUser(c, value)
		if user == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "",
			"data":    user,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "登陆信息已失效",
		})
		return
	}
}

func (w *User) Logout(c *gin.Context) {
	err := service.NewUserService(c.Request.Context()).Logout(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
	})
	return
}
