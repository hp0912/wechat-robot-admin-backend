package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/vars"

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

func (sv *MessageService) MessageRevoke(req dto.MessageRevokeRequest, robot *model.Robot) error {
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

func (sv *MessageService) SendTextMessage(req dto.SendTextMessageRequest, robot *model.Robot) error {
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

func (sv *MessageService) SendImageMessage(ctx *gin.Context, req dto.SendImageMessageRequest, file io.Reader, header *multipart.FileHeader, robot *model.Robot) error {
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

func (sv *MessageService) SendVideoMessage(ctx *gin.Context, req dto.SendVideoMessageRequest, file io.Reader, header *multipart.FileHeader, robot *model.Robot) error {
	robotURL := fmt.Sprintf("%s/message/send/video", robot.GetBaseURL())
	// 准备转发请求
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// 创建文件表单字段
	part, err := writer.CreateFormFile("video", header.Filename)
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

func (sv *MessageService) SendVoiceMessage(ctx *gin.Context, req dto.SendVoiceMessageRequest, file io.Reader, header *multipart.FileHeader, robot *model.Robot) error {
	robotURL := fmt.Sprintf("%s/message/send/voice", robot.GetBaseURL())
	// 准备转发请求
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// 创建文件表单字段
	part, err := writer.CreateFormFile("voice", header.Filename)
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

func (sv *MessageService) SendFileMessage(ctx *gin.Context, req dto.SendFileMessageRequest, chunk io.Reader, robot *model.Robot) error {
	robotURL := fmt.Sprintf("%s/message/send/file", robot.GetBaseURL())
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 分片文件字段名与前端一致: chunk
	part, err := writer.CreateFormFile("chunk", req.Filename)
	if err != nil {
		return err
	}
	if _, err = io.Copy(part, chunk); err != nil {
		return err
	}
	// 追加其他字段
	if err = writer.WriteField("to_wxid", req.ToWxid); err != nil {
		return err
	}
	if err = writer.WriteField("filename", req.Filename); err != nil {
		return err
	}
	if err = writer.WriteField("file_hash", req.FileHash); err != nil {
		return err
	}
	if err = writer.WriteField("file_size", strconv.FormatInt(req.FileSize, 10)); err != nil {
		return err
	}
	if err = writer.WriteField("chunk_index", strconv.Itoa(req.ChunkIndex)); err != nil {
		return err
	}
	if err = writer.WriteField("total_chunks", strconv.Itoa(req.TotalChunks)); err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}
	robotRequest, err := http.NewRequest("POST", robotURL, &requestBody)
	if err != nil {
		return err
	}
	robotRequest.Header.Set("Content-Type", writer.FormDataContentType())
	robotClient := &http.Client{}
	robotResp, err := robotClient.Do(robotRequest)
	if err != nil {
		return err
	}
	defer robotResp.Body.Close()
	if robotResp.StatusCode != http.StatusOK {
		return fmt.Errorf("robot service returned status %d", robotResp.StatusCode)
	}
	return nil
}

func (sv *MessageService) GetTimbre() ([]string, error) {
	var result dto.TimbreResponse
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("key", vars.ThirdPartyApiKey).
		SetQueryParam("type", "list").
		SetResult(&result).
		Get("https://api.pearktrue.cn/api/dub")
	if err != nil {
		return nil, err
	}
	if result.Code != 200 {
		return nil, fmt.Errorf("failed to get timbre: %s", result.Msg)
	}
	return result.Speakers, nil
}

func (sv *MessageService) SendAITTSMessage(ctx *gin.Context, req dto.RobotSendAITTSMessageRequest, robot *model.Robot) error {
	var result dto.RobotSendAITTSMessageResponse
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("key", vars.ThirdPartyApiKey).
		SetQueryParam("type", "mp3").
		SetQueryParam("dub", req.Speaker).
		SetQueryParam("text", req.Content).
		SetResult(&result).
		Get("https://api.pearktrue.cn/api/dub")
	if err != nil {
		return err
	}
	if result.Code != 200 {
		return fmt.Errorf("failed to send AI TTS message: %s", result.Msg)
	}
	// 下载生成的音频文件
	audioResp, err := http.Get(result.Audiopath)
	if err != nil {
		return fmt.Errorf("failed to download audio file: %v", err)
	}
	defer audioResp.Body.Close()

	if audioResp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download audio file, status: %d", audioResp.StatusCode)
	}
	voiceReq := dto.SendVoiceMessageRequest{
		ToWxid: req.ToWxid,
	}
	header := &multipart.FileHeader{
		Filename: "ai_tts_audio.mp3",
	}
	return sv.SendVoiceMessage(ctx, voiceReq, audioResp.Body, header, robot)
}
