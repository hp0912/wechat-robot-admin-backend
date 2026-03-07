package service

import (
	"context"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type SkillsService struct {
	ctx context.Context
}

func NewSkillsService(ctx context.Context) *SkillsService {
	return &SkillsService{
		ctx: ctx,
	}
}

func (s *SkillsService) GetSkills(robot *model.Robot) ([]*dto.Skill, error) {
	var result dto.Response[[]*dto.Skill]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/skills")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *SkillsService) InstallSkill(robot *model.Robot, req *dto.InstallSkillRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/skill/install")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *SkillsService) UpdateSkill(robot *model.Robot, req *dto.SkillRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Put(robot.GetBaseURL() + "/skill/update")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *SkillsService) EnableSkill(robot *model.Robot, req *dto.SkillRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/skill/enable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *SkillsService) DisableSkill(robot *model.Robot, req *dto.SkillRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/skill/disable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *SkillsService) UninstallSkill(robot *model.Robot, req *dto.SkillRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/skill/uninstall")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
