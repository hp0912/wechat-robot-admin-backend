package controller

import (
	"errors"
	"fmt"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type RobotLogin struct {
}

func NewRobotLoginController() *RobotLogin {
	return &RobotLogin{}
}

func (ct *RobotLogin) RobotLogin(c *gin.Context) {
	var req struct {
		LoginType string `form:"login_type" json:"login_type" binding:"required"`
	}
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
	data, err := service.NewRobotLoginService(c).RobotLogin(robot, req.LoginType)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) RobotLoginCheck(c *gin.Context) {
	var req dto.RobotLoginCheckRequest
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
	data, err := service.NewRobotLoginService(c).RobotLoginCheck(robot, req.Uuid)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) RobotLogin2FA(c *gin.Context) {
	var req dto.RobotLogin2FARequest
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
	err = service.NewRobotLoginService(c).RobotLogin2FA(robot, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotLogin) LoginSliderVerify(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")

	htmlErrorContent := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>获取滑块失败</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 600px; margin: 0 auto; }
        .result { padding: 20px; background: #f5f5f5; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        %s
    </div>
</body>
</html>`

	var req dto.SliderVerifyRequest
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		c.String(200, fmt.Sprintf(htmlErrorContent, err.Error()))
		return
	}
	robot, err := appx.GetRobot(c)
	if err != nil {
		c.String(200, fmt.Sprintf(htmlErrorContent, err.Error()))
		return
	}
	data, err := service.NewRobotLoginService(c).LoginSliderVerify(robot, req)
	if err != nil {
		c.String(200, fmt.Sprintf(htmlErrorContent, err.Error()))
		return
	}

	c.String(200, *data)
}

func (ct *RobotLogin) LoginData62Login(c *gin.Context) {
	var req dto.LoginRequest
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
	data, err := service.NewRobotLoginService(c).LoginData62Login(robot, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) LoginData62SMSAgain(c *gin.Context) {
	var req dto.LoginData62SMSAgainRequest
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
	data, err := service.NewRobotLoginService(c).LoginData62SMSAgain(robot, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) LoginData62SMSVerify(c *gin.Context) {
	var req dto.LoginData62SMSVerifyRequest
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
	data, err := service.NewRobotLoginService(c).LoginData62SMSVerify(robot, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) LoginA16Data1(c *gin.Context) {
	var req dto.LoginRequest
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
	data, err := service.NewRobotLoginService(c).LoginA16Data1(robot, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotLogin) RobotLogout(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err = service.NewRobotLoginService(c).RobotLogout(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotLogin) RobotState(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err = service.NewRobotLoginService(c).RobotState(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
