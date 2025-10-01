package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/go-resty/resty/v2"
)

type OSSSettingService struct {
	ctx      context.Context
	userRepo *repository.User
}

func NewOSSSettingService(ctx context.Context) *OSSSettingService {
	return &OSSSettingService{
		ctx:      ctx,
		userRepo: repository.NewUserRepo(ctx, vars.DB),
	}
}

func (s *OSSSettingService) GetOSSSettings(robot *model.Robot) (dto.OSSSettings, error) {
	var result dto.Response[dto.OSSSettings]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/oss-settings")
	if err = result.CheckError(err); err != nil {
		return dto.OSSSettings{}, err
	}
	return result.Data, nil
}

func (s *OSSSettingService) SaveOSSSettings(robot *model.Robot, req dto.OSSSettingsRequest) error {
	var result dto.Response[dto.OSSSettings]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/oss-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
