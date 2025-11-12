package controller

import (
	"errors"
	"io"
	"regexp"
	"time"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RobotManage struct {
}

func NewRobotManageController() *RobotManage {
	return &RobotManage{}
}

func (ct *RobotManage) RobotList(c *gin.Context) {
	var req dto.RobotListRequest
	resp := appx.NewResponse(c)
	session := sessions.Default(c)
	role := session.Get("role")
	wxid := session.Get("wechat_id")
	if role.(int) != vars.RoleRootUser {
		req.Owner = wxid.(string)
	}
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	pager := appx.InitPager(c)
	list, total, err := service.NewRobotManageService(c).RobotList(req, pager)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(list, total)
}

func (ct *RobotManage) RobotCreate(c *gin.Context) {
	var req dto.RobotCreateRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 编译正则表达式
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+$`)
	// 使用正则表达式匹配字符串
	if !re.MatchString(req.RobotCode) {
		resp.ToErrorResponse(errors.New("机器人编码只能包含字母、数字和下划线，并且必须以字母开头"))
		return
	}
	// 验证代理IP格式（如果启用了代理）
	if req.ProxyEnabled && req.ProxyIP != "" {
		// 简单的IP:PORT格式验证
		proxyRe := regexp.MustCompile(`^[^:]+:\d+$`)
		if !proxyRe.MatchString(req.ProxyIP) {
			resp.ToErrorResponse(errors.New("代理地址格式错误，应为 IP:端口 格式"))
			return
		}
	}
	err := service.NewRobotManageService(c).RobotCreate(c, req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotView(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robot, err := service.NewRobotManageService(c).RobotView(req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	if robot == nil {
		resp.ToErrorResponse(errors.New("机器人不存在"))
		return
	}
	resp.ToResponse(robot)
}

func (ct *RobotManage) RobotUpdate(c *gin.Context) {
	var req dto.RobotUpdateRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 验证代理IP格式（如果启用了代理）
	if req.ProxyEnabled && req.ProxyIP != "" {
		// 简单的IP:PORT格式验证
		re := regexp.MustCompile(`^[^:]+:\d+$`)
		if !re.MatchString(req.ProxyIP) {
			resp.ToErrorResponse(errors.New("代理地址格式错误，应为 IP:端口 格式"))
			return
		}
	}
	err := service.NewRobotManageService(c).RobotUpdate(req)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotRemove(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotRemove(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotRestartClient(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotRestartClient(req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotRestartServer(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotRestartServer(req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotDockerImagePull(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 设置SSE响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	// 创建进度通道
	progressChan := make(chan dto.PullProgress, 100)
	// 启动goroutine执行拉取任务
	go func() {
		err := service.NewRobotManageService(c).RobotDockerImagePull(c, progressChan)
		if err != nil {
			// 错误已经通过progressChan发送
			return
		}
	}()
	// 发送进度数据
	c.Stream(func(w io.Writer) bool {
		select {
		case progress, ok := <-progressChan:
			if !ok {
				// 通道已关闭，发送完成事件
				c.SSEvent("complete", "拉取完成")
				return false
			}
			// 发送进度数据
			c.SSEvent("progress", progress)
			return true
		case <-c.Request.Context().Done():
			// 客户端断开连接
			return false
		case <-time.After(300 * time.Second):
			// 超时处理
			c.SSEvent("error", "操作超时")
			return false
		}
	})
}

func (ct *RobotManage) RobotStopAndRemoveClient(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStopAndRemoveWeChatClient(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotStopAndRemoveServer(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStopAndRemoveWeChatServer(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotStartClient(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStartWeChatClient(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotStartServer(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStartWeChatServer(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) ExportRobotLoginData(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewRobotManageService(c).ExportRobotLoginData(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *RobotManage) ImportRobotLoginData(c *gin.Context) {
	var req struct {
		Data string `json:"data" binding:"required"`
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
	err = service.NewRobotManageService(c).ImportRobotLoginData(robot, req.Data)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
