package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type Contact struct {
}

func NewContactController() *Contact {
	return &Contact{}
}

func (ct *Contact) SyncContacts(c *gin.Context) {
	var req dto.SyncContactsRequest
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	service.NewContactService(c).SyncContact(robot)
	resp.ToResponse(nil)
}

func (ct *Contact) GetContacts(c *gin.Context) {
	var req dto.GetContactsRequest
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
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
