package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type ChatHistoryService struct {
	ctx context.Context
}

func NewChatHistoryService(ctx context.Context) *ChatHistoryService {
	return &ChatHistoryService{
		ctx: ctx,
	}
}

func (c *ChatHistoryService) GetChatHistory(req dto.ChatHistoryRequest, pager appx.Pager, robot *model.Robot) ([]*dto.ChatHistory, int64, error) {
	var result dto.Response[struct {
		Itmes []*dto.ChatHistory `json:"items"`
		Total int64              `json:"total"`
	}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("contact_id", req.ContactID).
		SetQueryParam("keyword", req.Keyword).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", "20").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/chat/history")
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}

func (c *ChatHistoryService) DownloadImageOrVoice(ctx *gin.Context, req dto.AttachDownloadRequest, robot *model.Robot, resp *appx.Response) {
	robotURL := fmt.Sprintf("%s%s?message_id=%d", robot.GetBaseURL(), req.AttachUrl, req.MessageID)
	client := &http.Client{
		Timeout: 300 * time.Second,
	}
	robotReq, err := http.NewRequest("GET", robotURL, nil)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	robotResp, err := client.Do(robotReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "发起下载请求失败: " + err.Error()})
		return
	}
	defer robotResp.Body.Close()
	if robotResp.StatusCode != http.StatusOK {
		ctx.Status(robotResp.StatusCode)
		io.Copy(ctx.Writer, robotResp.Body)
		return
	}
	for key, values := range robotResp.Header {
		for _, value := range values {
			ctx.Header(key, value)
		}
	}
	ctx.Status(robotResp.StatusCode)
	_, err = io.Copy(ctx.Writer, robotResp.Body)
	if err != nil {
		fmt.Printf("下载图片/语音失败: %v\n", err)
	}
}

func (c *ChatHistoryService) DownloadFileOrVideo(ctx *gin.Context, req dto.AttachDownloadRequest, robot *model.Robot) {
	robotURL := fmt.Sprintf("%s%s?message_id=%d", robot.GetBaseURL(), req.AttachUrl, req.MessageID)
	robotReq, err := http.NewRequest("GET", robotURL, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway,
			gin.H{"message": err.Error()})
		return
	}
	robotResp, err := http.DefaultClient.Do(robotReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway,
			gin.H{"message": err.Error()})
		return
	}
	defer robotResp.Body.Close()
	for key, values := range robotResp.Header {
		for _, value := range values {
			ctx.Header(key, value)
		}
	}
	ctx.Status(robotResp.StatusCode)
	_, err = io.Copy(ctx.Writer, robotResp.Body)
	if err != nil {
		fmt.Printf("下载附件/ 视频失败: %v\n", err)
	}
}
