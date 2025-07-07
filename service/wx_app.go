package service

import (
	"context"
	"io"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/tuotoo/qrcode"
)

type WXAppService struct {
	ctx context.Context
}

func NewWXAppService(ctx context.Context) *WXAppService {
	return &WXAppService{ctx: ctx}
}

func (s *WXAppService) WxappQrcodeAuthLogin(ctx *gin.Context, file io.Reader, robot *model.Robot) error {
	qrmatrix, err := qrcode.Decode(file)
	if err != nil {
		return err
	}

	var result dto.Response[struct{}]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"url": qrmatrix.Content,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/wxapp/qrcode-auth-login")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
