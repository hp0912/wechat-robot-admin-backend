package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/gin-gonic/gin"
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

func (s *MessageService) SendImageMessage(ctx *gin.Context, req dto.SendImageMessageRequest, file multipart.File, header *multipart.FileHeader, robot *model.Robot) error {
	robotURL := fmt.Sprintf("%s/message/send/image", robot.GetBaseURL())
	// 准备转发请求
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// 创建文件表单字段
	part, err := writer.CreateFormFile("image", header.Filename)
	if err != nil {
		return err
	}
	// 复制文件内容
	if _, err = io.Copy(part, file); err != nil {
		return err
	}
	// 添加其他表单字段
	if err := writer.WriteField("to_wxid", req.ToWxid); err != nil {
		return err
	}
	// 关闭multipart writer
	if err = writer.Close(); err != nil {
		return err
	}
	robotRequest, err := http.NewRequest("POST", robotURL, &requestBody)
	if err != nil {
		return err
	}
	// 设置请求头
	robotRequest.Header.Set("Content-Type", writer.FormDataContentType())
	// 发送请求并获取响应
	robotClient := &http.Client{}
	robotResp, err := robotClient.Do(robotRequest)
	if err != nil {
		return err
	}
	defer robotResp.Body.Close()

	return nil
}
