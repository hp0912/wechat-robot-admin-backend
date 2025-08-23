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
func (sv *RobotLoginService) RobotLogin(robot *model.Robot, loginType string) (dto.RobotLoginResponse, error) {
	var result dto.Response[dto.RobotLoginResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"login_type": loginType,
		}).
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

// RobotLogin2FA 新设备登陆双重认证
func (sv *RobotLoginService) RobotLogin2FA(robot *model.Robot, req dto.RobotLogin2FARequest) error {
	var result dto.Response[dto.RobotLoginCheckResponse]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"uuid":   req.Uuid,
			"data62": req.Data62,
			"code":   req.Code,
			"ticket": req.Ticket,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login/2fa")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (sv *RobotLoginService) LoginNewDeviceVerify(robot *model.Robot, ticket string) (*dto.SilderOCR, error) {
	var result dto.Response[dto.SilderOCR]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"ticket": ticket,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login/new-device-verify")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

func (sv *RobotLoginService) LoginData62Login(robot *model.Robot, req dto.LoginRequest) (any, error) {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login/data62")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

func (sv *RobotLoginService) LoginA16Data1(robot *model.Robot, req dto.LoginRequest) (any, error) {
	var result dto.Response[any]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/login/a16")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return &result.Data, nil
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
		err = respo.Update(&newRobot)
		if err != nil {
			return
		}
	} else {
		newRobot := model.Robot{
			ID:     robot.ID,
			Status: model.RobotStatusOffline,
		}
		err = respo.Update(&newRobot)
		if err != nil {
			return
		}
	}
	return
}
