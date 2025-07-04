package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type SystemMessageService struct {
	ctx context.Context
}

func NewSystemMessageService(ctx context.Context) *SystemMessageService {
	return &SystemMessageService{ctx: ctx}
}

func (s *SystemMessageService) GetRecentMonthMessages(robot *model.Robot) ([]*dto.SystemMessage, error) {
	var result dto.Response[[]*dto.SystemMessage]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/system-messages")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return result.Data, nil
}
