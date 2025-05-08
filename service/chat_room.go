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

type ChatRoomService struct {
	ctx context.Context
}

func NewChatRoomService(ctx context.Context) *ChatRoomService {
	return &ChatRoomService{
		ctx: ctx,
	}
}

func (c *ChatRoomService) SyncChatRoomMembers(robot *model.Robot, chatRoomID string) {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"chat_room_id": chatRoomID,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/members/sync")
	if err = result.CheckError(err); err != nil {
		log.Printf("同步群成员发生错误: %v", err)
	}
}

func (c *ChatRoomService) GetChatRoomMembers(req dto.ChatRoomMemberRequest, pager appx.Pager, robot *model.Robot) ([]*dto.ChatRoomMember, int64, error) {
	var result dto.Response[struct {
		Itmes []*dto.ChatRoomMember `json:"items"`
		Total int64                 `json:"total"`
	}]
	// 获取联系人列表
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("chat_room_id", req.ChatRoomID).
		SetQueryParam("keyword", req.Keyword).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", "20").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/chat-room/members")
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}
