package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type ChatRoomSettingsService struct {
	ctx context.Context
}

func NewChatRoomSettingsService(ctx context.Context) *ChatRoomSettingsService {
	return &ChatRoomSettingsService{
		ctx: ctx,
	}
}

func (sv *ChatRoomSettingsService) GetChatRoomSettings(chatRoomID string, robot *model.Robot) (resp dto.GetChatRoomSettingsResponse, err error) {
	var result dto.Response[dto.GetChatRoomSettingsResponse]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("chat_room_id", chatRoomID).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/chat-room-settings")
	if err = result.CheckError(err); err != nil {
		return
	}
	resp = result.Data
	return
}

func (sv *ChatRoomSettingsService) SaveChatRoomSettings(req dto.SaveChatRoomSettingsRequest, robot *model.Robot) error {
	var result dto.Response[struct{}]
	req.ID = req.ConfigID
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
