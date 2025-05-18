package service

import (
	"context"
	"log"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"

	"github.com/go-resty/resty/v2"
)

type ContactService struct {
	ctx context.Context
}

func NewContactService(ctx context.Context) *ContactService {
	return &ContactService{
		ctx: ctx,
	}
}

func (sv *ContactService) SyncContact(robot *model.Robot) {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contacts/sync")
	if err = result.CheckError(err); err != nil {
		log.Printf("同步联系人发生错误: %v", err)
	}
}

func (sv *ContactService) GetContacts(req dto.GetContactsRequest, pager appx.Pager, robot *model.Robot) ([]*dto.GetContactsResponse, int64, error) {
	var result dto.Response[struct {
		Itmes []*dto.GetContactsResponse `json:"items"`
		Total int64                      `json:"total"`
	}]
	// 获取联系人列表
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("keyword", req.Keyword).
		SetQueryParam("type", req.Type).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", "20").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/contacts")
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}
