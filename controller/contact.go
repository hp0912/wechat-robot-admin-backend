package controller

import (
	"errors"
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
	data, err := service.NewContactService(c).GetContacts(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}
