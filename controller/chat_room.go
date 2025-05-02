package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type ChatRoom struct {
}

func NewChatRoomController() *ChatRoom {
	return &ChatRoom{}
}

func (crm *ChatRoom) SyncChatRoomMembers(c *gin.Context) {
	var req dto.SyncChatRoomMemberRequest
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
	service.NewChatRoomService(c).SyncChatRoomMembers(robot, req.ChatRoomID)
	resp.ToResponse(nil)
}

func (crm *ChatRoom) GetChatRoomMembers(c *gin.Context) {
	var req dto.ChatRoomMemberRequest
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
	data, total, err := service.NewChatRoomService(c).GetChatRoomMembers(req, pager, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(data, total)
}
