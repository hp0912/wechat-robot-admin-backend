package controller

import (
	"errors"
	"fmt"
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
	robot := service.NewRobotManageService(c).RobotView(req.ID)
	if robot == nil {
		resp.ToErrorResponse(errors.New("机器人不存在"))
		return
	}
	resp.ToResponse(robot)
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
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Accel-Buffering", "no") // 禁用nginx缓冲

	// 立即刷新响应头
	c.Writer.Flush()

	// 创建进度通道
	progressChan := make(chan dto.PullProgress, 100)
	defer close(progressChan) // 确保通道被关闭
	// 启动goroutine执行拉取任务
	go func() {
		err := service.NewRobotManageService(c).RobotDockerImagePull(c, progressChan)
		if err != nil {
			// 错误已经通过progressChan发送
			return
		}
	}()
	// 发送进度数据
	for {
		select {
		case progress, ok := <-progressChan:
			if !ok {
				// 通道已关闭，发送完成事件
				fmt.Fprintf(c.Writer, "event: complete\ndata: 拉取完成\n\n")
				c.Writer.Flush()
				return
			}

			// 如果有错误，发送错误事件
			if progress.Error != "" {
				fmt.Fprintf(c.Writer, "event: error\ndata: %s\n\n", progress.Error)
				c.Writer.Flush()
				return
			}

			// 发送进度数据
			fmt.Fprintf(c.Writer, "event: progress\ndata: %s\n\n", progress)
			c.Writer.Flush()

		case <-c.Request.Context().Done():
			// 客户端断开连接
			return
		case <-time.After(30 * time.Second):
			// 超时处理
			fmt.Fprintf(c.Writer, "event: error\ndata: 操作超时\n\n")
			c.Writer.Flush()
			return
		}
	}
}

func (ct *RobotManage) RobotStopAndRemoveClientAndServer(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStopAndRemoveClientAndServer(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *RobotManage) RobotStartClientAndServer(c *gin.Context) {
	var req dto.RobotCommonRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	err := service.NewRobotManageService(c).RobotStartClientAndServer(c, req.ID)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
