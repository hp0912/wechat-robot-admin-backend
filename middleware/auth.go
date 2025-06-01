package middleware

import (
	"net/http"
	"strconv"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authHelper(c *gin.Context, minRole int) {
	session := sessions.Default(c)
	role := session.Get("role")
	id := session.Get("id")
	status := session.Get("status")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "请先登陆",
			"data":    nil,
		})
		c.Abort()
		return
	}
	if status.(int) == vars.UserStatusDisabled {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "用户已被封禁",
			"data":    nil,
		})
		c.Abort()
		return
	}
	if role.(int) < minRole {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "无权进行此操作，权限不足",
			"data":    nil,
		})
		c.Abort()
		return
	}
	c.Next()
}

func UserAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, vars.RoleCommonUser)
	}
}

func AdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, vars.RoleAdminUser)
	}
}

func RootAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelper(c, vars.RoleRootUser)
	}
}

func UserOwnerAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("id")
		role := session.Get("role")
		weChatId := session.Get("wechat_id")
		if id == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请先登陆",
				"data":    nil,
			})
			c.Abort()
			return
		}
		if role.(int) == vars.RoleRootUser {
			c.Next()
			return
		}
		idStr := c.Query("id") // 获取字符串
		robotId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "无效的机器人ID",
				"data":    nil,
			})
			c.Abort()
			return
		}
		robot, err := repository.NewRobotRepo(c.Request.Context(), vars.DB).GetByID(robotId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取机器人信息失败",
				"data":    nil,
			})
			c.Abort()
			return
		}
		if robot == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "机器人不存在",
				"data":    nil,
			})
			c.Abort()
			return
		}
		if robot.Owner != weChatId.(string) {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "无权进行此操作，权限不足",
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("robot", robot)
		c.Next()
	}
}
