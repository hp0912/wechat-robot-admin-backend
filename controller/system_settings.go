package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type SystemSettings struct{}

func NewSystemSettingsController() *SystemSettings {
	return &SystemSettings{}
}

func (ct *SystemSettings) GetSystemSettings(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewSystemSettingService(c).GetSystemSettings(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *SystemSettings) SaveSystemSettings(c *gin.Context) {
	var req dto.SystemSettings
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err = service.NewSystemSettingService(c).SaveSystemSettings(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
