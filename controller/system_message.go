package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type SystemMessage struct{}

func NewSystemMessageController() *SystemMessage {
	return &SystemMessage{}
}

func (ct *SystemMessage) GetRecentMonthMessages(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewSystemMessageService(c).GetRecentMonthMessages(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *SystemMessage) MarkAsReadBatch(c *gin.Context) {
	var req dto.MarkAsReadBatchRequest
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
	data, err := service.NewSystemMessageService(c).MarkAsReadBatch(req.SystemMessageIDs, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}
