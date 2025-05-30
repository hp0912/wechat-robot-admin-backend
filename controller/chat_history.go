package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type ChatHistory struct {
}

func NewChatHistoryController() *ChatHistory {
	return &ChatHistory{}
}

func (ct *ChatHistory) GetChatRoomMembers(c *gin.Context) {
	var req dto.ChatHistoryRequest
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
	data, total, err := service.NewChatHistoryService(c).GetChatHistory(req, pager, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(data, total)
}

func (ct *ChatHistory) DownloadImage(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/image/download"
	service.NewChatHistoryService(c).DownloadImageOrVoice(c, req, robot, resp)
}

func (ct *ChatHistory) DownloadVoice(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/voice/download"
	service.NewChatHistoryService(c).DownloadImageOrVoice(c, req, robot, resp)
}

func (ct *ChatHistory) DownloadFile(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/file/download"
	service.NewChatHistoryService(c).DownloadFileOrVideo(c, req, robot)
}

func (ct *ChatHistory) DownloadVideo(c *gin.Context) {
	var req dto.AttachDownloadRequest
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
	req.AttachUrl = "/chat/video/download"
	service.NewChatHistoryService(c).DownloadFileOrVideo(c, req, robot)
}
