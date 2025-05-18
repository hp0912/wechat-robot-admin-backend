package service

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/vars"

	"github.com/go-resty/resty/v2"
)

type RobotLoginService struct {
	ctx context.Context
}

func NewRobotLoginService(ctx context.Context) *RobotLoginService {
	return &RobotLoginService{
		ctx: ctx,
	}
}

// RobotLogin 获取机器人登陆二维码
func (sv *RobotLoginService) RobotLogin(robot *model.Robot) (dto.RobotLoginResponse, error) {
	var result dto.Response[dto.RobotLoginResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login")
	if err = result.CheckError(err); err != nil {
		return dto.RobotLoginResponse{}, err
	}
	return result.Data, nil
}

// RobotLoginCheck 检查机器人登陆状态
func (sv *RobotLoginService) RobotLoginCheck(robot *model.Robot, uuid string) (dto.RobotLoginCheckResponse, error) {
	var result dto.Response[dto.RobotLoginCheckResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"uuid": uuid,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login/check")
	if err = result.CheckError(err); err != nil {
		return dto.RobotLoginCheckResponse{}, err
	}
	return result.Data, nil
}

// RobotLogout 机器人登出
func (sv *RobotLoginService) RobotLogout(robot *model.Robot) (err error) {
	var resp dto.Response[struct{}]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&resp).
		Delete(robot.GetBaseURL() + "/logout")
	if err = resp.CheckError(err); err != nil {
		return
	}
	return
}

// RobotState 获取机器人状态
func (sv *RobotLoginService) RobotState(robot *model.Robot) (err error) {
	var isRunningResp dto.Response[bool]
	var isLoggedInResp dto.Response[bool]
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&isRunningResp).
		Get(robot.GetBaseURL() + "/is-running")
	if err = isRunningResp.CheckError(err); err != nil {
		return
	}
	_, err = resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&isLoggedInResp).
		Get(robot.GetBaseURL() + "/is-loggedin")
	if err = isLoggedInResp.CheckError(err); err != nil {
		return
	}
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	if isRunningResp.Data && isLoggedInResp.Data {
		newRobot := model.Robot{
			ID:     robot.ID,
			Status: model.RobotStatusOnline,
		}
		respo.Update(&newRobot)
	} else {
		newRobot := model.Robot{
			ID:     robot.ID,
			Status: model.RobotStatusOffline,
		}
		respo.Update(&newRobot)
	}
	return
}
