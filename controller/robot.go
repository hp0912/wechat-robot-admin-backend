package controller

import (
	"errors"
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
