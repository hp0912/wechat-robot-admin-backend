package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type Contact struct {
}

func NewContactController() *Contact {
	return &Contact{}
}

func (ct *Contact) GetContacts(c *gin.Context) {
	var req dto.GetContactsRequest
	resp := appx.NewResponse(c)
	_robot, exists := c.Get("robot")
	if !exists {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, ok := _robot.(*model.Robot)
	if !ok {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	pager := appx.InitPager(c)
	data, total, err := service.NewContactService(c).GetContacts(req, pager, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(data, total)
}
