package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type GlobalSettings struct {
}

func NewGlobalSettingsController() *GlobalSettings {
	return &GlobalSettings{}
}

func (ct *GlobalSettings) GetGlobalSettings(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	globalSettings, err := service.NewGlobalSettingsService(c).GetGlobalSettings(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(globalSettings)
}

func (ct *GlobalSettings) SaveGlobalSettings(c *gin.Context) {
	var req dto.SaveGlobalSettingsRequest
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
	err = service.NewGlobalSettingsService(c).SaveGlobalSettings(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
