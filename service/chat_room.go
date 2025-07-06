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

func (sv *ChatRoomService) SyncChatRoomMembers(robot *model.Robot, chatRoomID string) {
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

func (sv *ChatRoomService) GetChatRoomMembers(req dto.ChatRoomMemberRequest, pager appx.Pager, robot *model.Robot) ([]*dto.ChatRoomMember, int64, error) {
	var result dto.Response[struct {
		Itmes []*dto.ChatRoomMember `json:"items"`
		Total int64                 `json:"total"`
	}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("chat_room_id", req.ChatRoomID).
		SetQueryParam("keyword", req.Keyword).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", strconv.Itoa(pager.PageSize)).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/chat-room/members")
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}

func (sv *ChatRoomService) GetNotLeftMembers(req dto.ChatRoomMemberRequest, robot *model.Robot) ([]*dto.ChatRoomMember, error) {
	var result dto.Response[[]*dto.ChatRoomMember]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("chat_room_id", req.ChatRoomID).
		SetQueryParam("keyword", req.Keyword).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/chat-room/not-left-members")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (sv *ChatRoomService) CreateChatRoom(contactIDs []string, robot *model.Robot) error {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string][]string{
			"contact_ids": contactIDs,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/create")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) InviteChatRoomMember(chatRoomID string, contactIDs []string, robot *model.Robot) error {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"chat_room_id": chatRoomID,
			"contact_ids":  contactIDs,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/invite")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupConsentToJoin(id int64, robot *model.Robot) error {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{
			"system_message_id": id,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/join")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupSetChatRoomName(chatRoomID, content string, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"chat_room_id": chatRoomID,
			"content":      content,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/name")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupSetChatRoomRemarks(chatRoomID, content string, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"chat_room_id": chatRoomID,
			"content":      content,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/remark")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupSetChatRoomAnnouncement(chatRoomID, content string, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"chat_room_id": chatRoomID,
			"content":      content,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/chat-room/announcement")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupDelChatRoomMember(chatRoomID string, memberIDs []string, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"chat_room_id": chatRoomID,
			"member_ids":   memberIDs,
		}).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/chat-room/members")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ChatRoomService) GroupQuit(chatRoomID string, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"chat_room_id": chatRoomID,
		}).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/chat-room/quit")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
