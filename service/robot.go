package service

import (
	"context"
	"wechat-robot-client/dto"
	"wechat-robot-client/model"
	"wechat-robot-client/pkg/appx"
	"wechat-robot-client/repository"
	"wechat-robot-client/vars"

	"github.com/gin-gonic/gin"
)

type RobotService struct {
	ctx context.Context
}

func NewRobotService(ctx context.Context) *RobotService {
	return &RobotService{
		ctx: ctx,
	}
}

func (r *RobotService) RobotList(ctx *gin.Context, req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	return repository.NewRobotRepo(r.ctx, vars.DB).RobotList(ctx, req, pager)
}
