package service

import (
	"context"
	"strconv"

	"github.com/go-resty/resty/v2"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
)

type SystemPromptService struct {
	ctx context.Context
}

func NewSystemPromptService(ctx context.Context) *SystemPromptService {
	return &SystemPromptService{ctx: ctx}
}

func (s *SystemPromptService) ListSystemPrompts(robot *model.Robot, req *dto.ListSystemPromptRequest) ([]*dto.SystemPrompt, error) {
	var result dto.Response[[]*dto.SystemPrompt]
	request := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result)
	if req != nil && req.Keyword != "" {
		request.SetQueryParam("keyword", req.Keyword)
	}
	_, err := request.Get(robot.GetBaseURL() + "/system-prompts")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *SystemPromptService) GetSystemPrompt(robot *model.Robot, req *dto.GetSystemPromptRequest) (*dto.SystemPrompt, error) {
	var result dto.Response[*dto.SystemPrompt]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("id", strconv.FormatInt(req.PromptID, 10)).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/system-prompt")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *SystemPromptService) CreateSystemPrompt(robot *model.Robot, req *dto.CreateSystemPromptRequest) (*dto.SystemPrompt, error) {
	var result dto.Response[*dto.SystemPrompt]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/system-prompt")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *SystemPromptService) UpdateSystemPrompt(robot *model.Robot, req *dto.UpdateSystemPromptRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Put(robot.GetBaseURL() + "/system-prompt")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *SystemPromptService) DeleteSystemPrompt(robot *model.Robot, req *dto.DeleteSystemPromptRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/system-prompt")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
