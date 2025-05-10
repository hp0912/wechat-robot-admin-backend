package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) MessageRevoke(req dto.MessageRevokeRequest, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(dto.RobotMessageRevokeRequest{
			MessageID: req.MessageID,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/message/revoke")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *MessageService) SendTextMessage(req dto.SendTextMessageRequest, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(dto.RobotSendTextMessageRequest{
			ToWxid:  req.ToWxid,
			Content: req.Content,
			At:      req.At,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/message/send/text")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
