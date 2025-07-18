package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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
		SetQueryParam("max_id", req.MaxID).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/moments/list")
	if err = result.CheckError(err); err != nil {
		return dto.MomentsGetListResponse{}, err
	}
	for _, ObjectItem := range result.Data.ObjectList {
		if ObjectItem == nil {
			continue
		}
		if ObjectItem.Id != nil {
			ObjectItem.IdStr = strconv.FormatUint(*ObjectItem.Id, 10)
			ObjectItem.Id = nil
		}
		if ObjectItem.TimelineObject == nil {
			continue
		}
		medias := ObjectItem.TimelineObject.ContentObject.MediaList.Media
		for i := range medias {
			medias[i].IDStr = strconv.FormatUint(medias[i].ID, 10)
			if medias[i].Type == 6 {
				medias[i].VideoDurationStr = strconv.FormatFloat(medias[i].VideoDuration, 'f', -1, 64)
			}
		}
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

func (s *MomentsService) FriendCircleUpload(file io.Reader, header *multipart.FileHeader, robot *model.Robot) (resp dto.FriendCircleMedia, err error) {
	robotURL := fmt.Sprintf("%s/moments/upload-media", robot.GetBaseURL())
	// 准备转发请求
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	// 创建文件表单字段
	var part io.Writer
	part, err = writer.CreateFormFile("media", header.Filename)
	if err != nil {
		return
	}
	// 复制文件内容
	if _, err = io.Copy(part, file); err != nil {
		return
	}
	// 关闭multipart writer
	if err = writer.Close(); err != nil {
		return
	}
	var robotRequest *http.Request
	robotRequest, err = http.NewRequest("POST", robotURL, &requestBody)
	if err != nil {
		return
	}
	// 设置请求头
	robotRequest.Header.Set("Content-Type", writer.FormDataContentType())
	// 发送请求并获取响应
	robotClient := &http.Client{}
	var robotResp *http.Response
	robotResp, err = robotClient.Do(robotRequest)
	if err != nil {
		return
	}
	defer robotResp.Body.Close()
	// 解析响应体为结构体
	var result dto.Response[dto.FriendCircleMedia]
	if err = json.NewDecoder(robotResp.Body).Decode(&result); err != nil {
		return
	}
	// 检查响应状态和错误
	if err = result.CheckError(nil); err != nil {
		return
	}
	if result.Data.Id != nil {
		idStr := strconv.FormatUint(*result.Data.Id, 10)
		result.Data.IdStr = &idStr
		result.Data.Id = nil
	}
	if result.Data.VideoDuration > 0 {
		videoDurationStr := strconv.FormatFloat(result.Data.VideoDuration, 'f', -1, 64)
		result.Data.VideoDurationStr = videoDurationStr
		result.Data.VideoDuration = 0
	}
	// 返回解析后的数据
	resp = result.Data
	return
}

func (s *MomentsService) FriendCirclePost(req dto.MomentPostRequest, robot *model.Robot) (dto.MomentPostResponse, error) {
	for i := range req.MediaList {
		if req.MediaList[i].IdStr != nil && *req.MediaList[i].IdStr != "" {
			mediaId, err := strconv.ParseUint(*req.MediaList[i].IdStr, 10, 64)
			if err != nil {
				return dto.MomentPostResponse{}, fmt.Errorf("invalid media ID: %v", err)
			}
			req.MediaList[i].Id = &mediaId
		}
	}
	var result dto.Response[dto.MomentPostResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/moments/post")
	if err = result.CheckError(err); err != nil {
		return dto.MomentPostResponse{}, err
	}
	return result.Data, nil
}

func (s *MomentsService) FriendCircleOperation(req dto.MomentOpRequest, robot *model.Robot) (dto.MomentOpResponse, error) {
	var result dto.Response[dto.MomentOpResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]any{
			"Id":        req.MomentID,
			"Type":      req.Type,
			"CommentId": req.CommentId,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/moments/operate")
	if err = result.CheckError(err); err != nil {
		return dto.MomentOpResponse{}, err
	}
	return result.Data, nil
}

func (s *MomentsService) FriendCirclePrivacySettings(req dto.MomentPrivacySettingsRequest, robot *model.Robot) (dto.MomentPrivacySettingsResponse, error) {
	var result dto.Response[dto.MomentPrivacySettingsResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]uint32{
			"Function": req.Function,
			"Value":    req.Value,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/moments/privacy-settings")
	if err = result.CheckError(err); err != nil {
		return dto.MomentPrivacySettingsResponse{}, err
	}
	return result.Data, nil
}
