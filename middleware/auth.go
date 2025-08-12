package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authHelper(c *gin.Context, minRole int) {
	user, ok := resolveUserFromSessionOrToken(c)
	if !ok || user == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "请先登陆或提供有效API Token",
			"data":    nil,
		})
		c.Abort()
		return
	}
	if user.Status == vars.UserStatusDisabled {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "用户已被封禁",
			"data":    nil,
		})
		c.Abort()
		return
	}
	if user.Role < minRole {
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

// resolveUserFromSessionOrToken 按顺序从 Session、Authorization Header、X-API-Token Header、api_token Query 中解析用户
func resolveUserFromSessionOrToken(c *gin.Context) (*model.User, bool) {
	session := sessions.Default(c)
	if id := session.Get("id"); id != nil {
		user := &model.User{}
		if v, ok := id.(int64); ok {
			user.ID = v
		}
		if v := session.Get("wechat_id"); v != nil {
			if s, ok := v.(string); ok {
				user.WeChatId = s
			}
		}
		if v := session.Get("role"); v != nil {
			if i, ok := v.(int); ok {
				user.Role = i
			}
		}
		if v := session.Get("status"); v != nil {
			if i, ok := v.(int); ok {
				user.Status = i
			}
		}
		// 将解析到的用户写入上下文，便于后续复用
		c.Set("login_user", user)
		return user, true
	}

	token := c.GetHeader("Authorization")
	if token != "" {
		lower := strings.ToLower(token)
		if strings.HasPrefix(lower, "bearer ") {
			token = token[7:]
		}
	}
	if token == "" {
		token = c.GetHeader("X-API-Token")
	}
	if token == "" {
		token = c.Query("api_token")
	}
	if token == "" {
		return nil, false
	}
	user, err := repository.NewUserRepo(c.Request.Context(), vars.DB).GetUserByApiToken(token)
	if err != nil || user == nil {
		return nil, false
	}
	c.Set("login_user", user)
	return user, true
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
		user, ok := resolveUserFromSessionOrToken(c)
		if !ok || user == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请先登陆或提供有效API Token",
				"data":    nil,
			})
			c.Abort()
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
		c.Set("robot", robot)
		if user.Role == vars.RoleRootUser {
			c.Next()
			return
		}
		if robot.Owner != user.WeChatId {
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "无权进行此操作，权限不足",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
