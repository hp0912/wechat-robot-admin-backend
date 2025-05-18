package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type FriendSettingsService struct {
	ctx context.Context
}

func NewFriendSettingsService(ctx context.Context) *FriendSettingsService {
	return &FriendSettingsService{
		ctx: ctx,
	}
}

func (sv *FriendSettingsService) GetFriendSettings(contactID string, robot *model.Robot) (resp dto.GetFriendSettingsResponse, err error) {
	var result dto.Response[dto.GetFriendSettingsResponse]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("contact_id", contactID).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/friend-settings")
	if err = result.CheckError(err); err != nil {
		return
	}
	resp = result.Data
	return
}

func (sv *FriendSettingsService) SaveFriendSettings(req dto.SaveFriendSettingsRequest, robot *model.Robot) error {
	var result dto.Response[struct{}]
	req.ID = req.ConfigID
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/friend-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
