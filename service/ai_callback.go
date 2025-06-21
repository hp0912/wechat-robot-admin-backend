package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/go-resty/resty/v2"
)

type AICallbackService struct {
	ctx        context.Context
	robotRespo *repository.Robot
}

func NewAICallbackService(ctx context.Context) *AICallbackService {
	return &AICallbackService{ctx: ctx, robotRespo: repository.NewRobotRepo(ctx, vars.DB)}
}

func (s *AICallbackService) DoubaoTTS(req dto.DoubaoTTSCallbackRequest, robotId int64) error {
	robot, err := s.robotRespo.GetByID(robotId)
	if err != nil {
		return err
	}
	var result dto.Response[any]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/ai-callback/voice/doubao-tts")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
