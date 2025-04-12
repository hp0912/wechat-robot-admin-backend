package service

import (
	"context"
)

type RobotService struct {
	ctx context.Context
}

func NewRobotService(ctx context.Context) *RobotService {
	return &RobotService{
		ctx: ctx,
	}
}
