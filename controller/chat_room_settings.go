package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type ChatRoomSettings struct {
}

func NewChatRoomSettingsController() *ChatRoomSettings {
	return &ChatRoomSettings{}
}

func (ct *ChatRoomSettings) GetChatRoomSettings(c *gin.Context) {
	var req dto.GetChatRoomSettingsRequest
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
	chatRoomSettings, err := service.NewChatRoomSettingsService(c).GetChatRoomSettings(req.ChatRoomID, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(chatRoomSettings)
}

func (ct *ChatRoomSettings) SaveChatRoomSettings(c *gin.Context) {
	var req dto.SaveChatRoomSettingsRequest
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
	err = service.NewChatRoomSettingsService(c).SaveChatRoomSettings(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
