package controller

import (
	"errors"
	"regexp"
	"wechat-robot-client/dto"
	"wechat-robot-client/pkg/appx"
	"wechat-robot-client/service"
	"wechat-robot-client/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Robot struct {
}

func NewRobotController() *Robot {
	return &Robot{}
}

func (r *Robot) RobotList(c *gin.Context) {
	var req dto.RobotListRequest
	resp := appx.NewResponse(c)
	session := sessions.Default(c)
	role := session.Get("role")
	wxid := session.Get("wechat_id")
	if role.(int) != vars.RoleRootUser {
		req.Owner = wxid.(string)
	}
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	pager := appx.InitPager(c)
	list, total, err := service.NewRobotService(c.Request.Context()).RobotList(c, req, pager)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(list, total)
}

func (r *Robot) RobotCreate(c *gin.Context) {
	var req dto.RobotCreateRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 编译正则表达式
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+$`)
	// 使用正则表达式匹配字符串
	if !re.MatchString(req.RobotCode) {
		resp.ToErrorResponse(errors.New("机器人编码只能包含字母、数字和下划线，并且必须以字母开头"))
		return
	}
	err := service.NewRobotService(c.Request.Context()).RobotCreate(c, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (r *Robot) RobotView(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot := service.NewRobotService(c.Request.Context()).RobotView(c, req.ID)
	if robot == nil {
		resp.ToErrorResponse(errors.New("机器人不存在"))
		return
	}
	resp.ToResponse(robot)
}

func (r *Robot) RobotRemove(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotService(c.Request.Context()).RobotRemove(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
