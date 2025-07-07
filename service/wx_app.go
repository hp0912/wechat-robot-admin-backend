package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type WXAppService struct {
	ctx context.Context
}

func NewWXAppService(ctx context.Context) *WXAppService {
	return &WXAppService{ctx: ctx}
}

func (s *WXAppService) WxappQrcodeAuthLogin(URL string, robot *model.Robot) error {
	var result dto.Response[dto.MomentsGetListResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"url": URL,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/wxapp/qrcode-auth-login")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
