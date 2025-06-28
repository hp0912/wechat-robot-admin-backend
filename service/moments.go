package service

import (
	"context"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type MomentsService struct {
	ctx context.Context
}

func NewMomentsService(ctx context.Context) *MomentsService {
	return &MomentsService{ctx: ctx}
}

func (s *MomentsService) FriendCircleGetList(req dto.MomentsGetListRequest, robot *model.Robot) (dto.MomentsGetListResponse, error) {
	var result dto.Response[dto.MomentsGetListResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("frist_page_md5", req.FristPageMd5).
		SetQueryParam("max_id", strconv.FormatInt(req.MaxID, 10)).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/moments/list")
	if err = result.CheckError(err); err != nil {
		return dto.MomentsGetListResponse{}, err
	}
	return result.Data, nil
}

func (s *MomentsService) FriendCircleDownFriendCircleMedia(req dto.MomentsDownFriendCircleMediaRequest, robot *model.Robot) (string, error) {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("url", req.Url).
		SetQueryParam("key", req.Key).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/moments/down-media")
	if err = result.CheckError(err); err != nil {
		return "", err
	}
	return result.Data, nil
}
