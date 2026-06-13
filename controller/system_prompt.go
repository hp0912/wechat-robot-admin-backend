package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"
)

type SystemPrompt struct{}

func NewSystemPromptController() *SystemPrompt {
	return &SystemPrompt{}
}

func (p *SystemPrompt) ListSystemPrompts(c *gin.Context) {
	var req dto.ListSystemPromptRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewSystemPromptService(c).ListSystemPrompts(robot, &req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (p *SystemPrompt) GetSystemPrompt(c *gin.Context) {
	var req dto.GetSystemPromptRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewSystemPromptService(c).GetSystemPrompt(robot, &req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (p *SystemPrompt) CreateSystemPrompt(c *gin.Context) {
	var req dto.CreateSystemPromptRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewSystemPromptService(c).CreateSystemPrompt(robot, &req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (p *SystemPrompt) UpdateSystemPrompt(c *gin.Context) {
	var req dto.UpdateSystemPromptRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	if err = service.NewSystemPromptService(c).UpdateSystemPrompt(robot, &req); err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (p *SystemPrompt) DeleteSystemPrompt(c *gin.Context) {
	var req dto.DeleteSystemPromptRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	if err = service.NewSystemPromptService(c).DeleteSystemPrompt(robot, &req); err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
