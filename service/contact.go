package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"

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

func (c *ContaceService) GetContacts(req dto.GetContactsRequest, pager appx.Pager, robot *model.Robot) ([]*dto.GetContactsResponse, int64, error) {
	var result dto.Response[struct {
		Itmes []*dto.GetContactsResponse `json:"items"`
		Total int64                      `json:"total"`
	}]
	// 获取联系人之前，先手动同步一次
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Post(fmt.Sprintf("http://%s:%d/api/v1/robot/sync-contact", robot.RobotCode, 9002)) // TODO
	if err = result.CheckError(err); err != nil {
		log.Printf("同步联系人发生错误: %v", err)
	}
	// 获取联系人列表
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("keyword", req.Keyword).
		SetQueryParam("type", req.Type).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", strconv.Itoa(pager.PageSize)).
		SetResult(&result).
		Get(fmt.Sprintf("http://%s:%d/api/v1/robot/contacts", robot.RobotCode, 9002)) // TODO
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}
