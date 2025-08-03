package service

import (
	"context"
	"log"
	"net/url"
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
	queryParams := url.Values{}
	queryParams.Add("keyword", req.Keyword)
	queryParams.Add("type", req.Type)
	queryParams.Add("page_index", strconv.Itoa(pager.PageIndex))
	if pager.PageSize > 0 {
		queryParams.Add("page_size", strconv.Itoa(pager.PageSize))
	} else {
		queryParams.Add("page_size", "20")
	}
	for _, contactID := range req.ContactIDs {
		queryParams.Add("contact_ids", contactID)
	}
	// 获取联系人列表
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParamsFromValues(queryParams).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/contacts")
	if err = result.CheckError(err); err != nil {
		return nil, 0, err
	}
	return result.Data.Itmes, result.Data.Total, nil
}

func (sv *ContactService) FriendSearch(req dto.FriendSearchRequest, robot *model.Robot) (dto.FriendSearchResponse, error) {
	var result dto.Response[dto.FriendSearchResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"to_username":  req.ToUserName,
			"from_scene":   req.FromScene,
			"search_scene": req.SearchScene,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contact/friend/search")
	if err = result.CheckError(err); err != nil {
		return dto.FriendSearchResponse{}, err
	}
	return result.Data, nil
}

func (sv *ContactService) FriendSendRequest(req dto.FriendSendRequestRequest, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"v1":             req.V1,
			"v2":             req.V2,
			"opcode":         req.Opcode,
			"scene":          req.Scene,
			"verify_content": req.VerifyContent,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contact/friend/add")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ContactService) FriendSendRequestFromChatRoom(req dto.FriendSendRequestFromChatRoomRequest, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"chat_room_member_id": req.ChatRoomMemberID,
			"verify_content":      req.VerifyContent,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contact/friend/add-from-chat-room")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ContactService) FriendSetRemarks(req dto.FriendSetRemarksRequest, robot *model.Robot) error {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"to_wxid": req.ToWxid,
			"remarks": req.Remarks,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contact/friend/remark")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ContactService) FriendPassVerify(id int64, robot *model.Robot) error {
	var result dto.Response[any]
	// 通过好友验证
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{
			"system_message_id": id,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/contact/friend/pass-verify")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *ContactService) FriendDelete(contactID string, robot *model.Robot) error {
	var result dto.Response[any]
	// 删除好友
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"contact_id": contactID,
		}).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/contact/friend")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
