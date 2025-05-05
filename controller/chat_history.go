package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type ChatHistory struct {
}

func NewChatHistoryController() *ChatHistory {
	return &ChatHistory{}
}

func (ch *ChatHistory) GetChatRoomMembers(c *gin.Context) {
	var req dto.ChatHistoryRequest
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
	data, total, err := service.NewChatHistoryService(c).GetChatHistory(req, pager, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(data, total)
}

func (ch *ChatHistory) DownloadImage(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/image/download"
	service.NewChatHistoryService(c).DownloadImageOrVoice(c, req, robot, resp)
}

func (ch *ChatHistory) DownloadVoice(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/voice/download"
	service.NewChatHistoryService(c).DownloadImageOrVoice(c, req, robot, resp)
}

func (ch *ChatHistory) DownloadFile(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/file/download"
	service.NewChatHistoryService(c).DownloadFileOrVideo(c, req, robot)
}

func (ch *ChatHistory) DownloadVideo(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/video/download"
	service.NewChatHistoryService(c).DownloadFileOrVideo(c, req, robot)
}
