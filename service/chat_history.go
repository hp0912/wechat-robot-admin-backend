package service

import (
	"context"
	"fmt"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"

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
		Get(fmt.Sprintf("http://%s:%d/api/v1/robot/chat/history", robot.RobotCode, 9002)) // TODO
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}
