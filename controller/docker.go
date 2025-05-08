package controller

import (
	"errors"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type DockerController struct {
}

func NewDockerController() *DockerController {
	return &DockerController{}
}

func (d *DockerController) RobotContainerStats(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	stats, err := service.NewDockerService(c).RobotContainerStats(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(stats)
}

func (d *DockerController) GetRobotContainerLogs(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	logs, err := service.NewDockerService(c).GetRobotContainerLogs(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(logs)
}
