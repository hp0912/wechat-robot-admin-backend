package controller

import (
	"errors"
	"path/filepath"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type WXApp struct{}

func NewWXAppController() *WXApp {
	return &WXApp{}
}

func (ct *WXApp) WxappQrcodeAuthLogin(c *gin.Context) {
	resp := appx.NewResponse(c)
	// 获取表单文件
	file, fileHeader, err := c.Request.FormFile("qrcode")
	if err != nil {
		resp.ToErrorResponse(errors.New("获取上传文件失败"))
		return
	}
	defer file.Close()

	// 检查文件大小
	if fileHeader.Size > 50*1024*1024 { // 限制为50MB
		resp.ToErrorResponse(errors.New("文件大小不能超过50MB"))
		return
	}

	// 检查文件类型
	ext := filepath.Ext(fileHeader.Filename)
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	if !allowedExts[ext] {
		resp.ToErrorResponse(errors.New("不支持的图片格式"))
		return
	}

	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	err = service.NewWXAppService(c).WxappQrcodeAuthLogin(c, file, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
