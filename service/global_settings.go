package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type GlobalSettingsService struct {
	ctx context.Context
}

func NewGlobalSettingsService(ctx context.Context) *GlobalSettingsService {
	return &GlobalSettingsService{
		ctx: ctx,
	}
}

func (sv *GlobalSettingsService) GetGlobalSettings(robotID int64, robot *model.Robot) (resp dto.GetGlobalSettingsResponse, err error) {
	var result dto.Response[dto.GetGlobalSettingsResponse]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/global-settings")
	if err = result.CheckError(err); err != nil {
		return
	}
	resp = result.Data
	return
}

func (sv *GlobalSettingsService) SaveGlobalSettings(req dto.SaveGlobalSettingsRequest, robot *model.Robot) error {
	var result dto.Response[struct{}]
	req.ID = req.ConfigID
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/global-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
