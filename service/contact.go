package service

import (
	"context"
	"fmt"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type ContaceService struct {
	ctx context.Context
}

func NewContactService(ctx context.Context) *ContaceService {
	return &ContaceService{
		ctx: ctx,
	}
}

func (c *ContaceService) GetContacts(robot *model.Robot) ([]*dto.GetContactsResponse, error) {
	var result dto.Response[[]*dto.GetContactsResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(fmt.Sprintf("http://%s:%d/api/v1/robot/contacts", robot.RobotCode, 9002)) // TODO
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return result.Data, nil
}
