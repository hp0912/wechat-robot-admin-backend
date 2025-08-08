package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type SystemSettingService struct {
	ctx context.Context
}

func NewSystemSettingService(ctx context.Context) *SystemSettingService {
	return &SystemSettingService{ctx: ctx}
}

func (s *SystemSettingService) GetSystemSettings(robot *model.Robot) (dto.SystemSettings, error) {
	var result dto.Response[dto.SystemSettings]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/system-settings")
	if err = result.CheckError(err); err != nil {
		return dto.SystemSettings{}, err
	}
	return result.Data, nil
}

func (s *SystemSettingService) SaveSystemSettings(req dto.SystemSettingsRequest, robot *model.Robot) error {
	var result dto.Response[dto.SystemSettings]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/system-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
