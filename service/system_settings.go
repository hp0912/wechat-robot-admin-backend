package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/go-resty/resty/v2"
)

type SystemSettingService struct {
	ctx      context.Context
	userRepo *repository.User
}

func NewSystemSettingService(ctx context.Context) *SystemSettingService {
	return &SystemSettingService{
		ctx:      ctx,
		userRepo: repository.NewUserRepo(ctx, vars.DB),
	}
}

func (s *SystemSettingService) GetSystemSettings(user *model.User, robot *model.Robot) (dto.SystemSettings, error) {
	var result dto.Response[dto.SystemSettings]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/system-settings")
	if err = result.CheckError(err); err != nil {
		return dto.SystemSettings{}, err
	}
	if result.Data.APITokenEnabled != nil && *result.Data.APITokenEnabled {
		user, err := s.userRepo.GetUserByID(user.ID)
		if err != nil {
			return dto.SystemSettings{}, err
		}
		result.Data.APIToken = user.ApiToken
	}
	return result.Data, nil
}

func (s *SystemSettingService) SaveSystemSettings(req dto.SystemSettingsRequest, user *model.User, robot *model.Robot) error {
	currentUser, err := s.userRepo.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	if req.APITokenEnabled != nil {
		if !*req.APITokenEnabled {
			if currentUser.ApiToken != nil && *currentUser.ApiToken != "" {
				apiToken := ""
				err := s.userRepo.Update(&model.User{
					ID:       user.ID,
					ApiToken: &apiToken,
				})
				if err != nil {
					return err
				}
			}
		}
		if *req.APITokenEnabled {
			if currentUser.ApiToken == nil || *currentUser.ApiToken == "" {
				_, err := NewUserService(s.ctx).RefreshUserApiToken(user.ID)
				if err != nil {
					return err
				}
			}
		}
	}
	var result dto.Response[dto.SystemSettings]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(dto.SystemSettings{
			ID:                         req.SystemSettingsID,
			APITokenEnabled:            req.APITokenEnabled,
			OfflineNotificationEnabled: req.OfflineNotificationEnabled,
			NotificationType:           req.NotificationType,
			PushPlusURL:                req.PushPlusURL,
			PushPlusToken:              req.PushPlusToken,
			AutoVerifyUser:             req.AutoVerifyUser,
			VerifyUserDelay:            req.VerifyUserDelay,
			AutoChatroomInvite:         req.AutoChatroomInvite,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/system-settings")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
