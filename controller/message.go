package controller

import (
	"errors"
	"path/filepath"
	"unicode/utf8"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-gonic/gin"
)

type Message struct {
}

func NewMessageController() *Message {
	return &Message{}
}

func (ct *Message) MessageRevoke(c *gin.Context) {
	var req dto.MessageRevokeRequest
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
	err = service.NewMessageService(c).MessageRevoke(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Message) SendTextMessage(c *gin.Context) {
	var req dto.SendTextMessageRequest
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
	err = service.NewMessageService(c).SendTextMessage(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Message) SendImageMessage(c *gin.Context) {
	resp := appx.NewResponse(c)
	// 获取表单文件
	file, fileHeader, err := c.Request.FormFile("image")
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
	// 解析表单参数
	var req dto.SendImageMessageRequest
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	err = service.NewMessageService(c).SendImageMessage(c, req, file, fileHeader, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Message) SendVideoMessage(c *gin.Context) {
	resp := appx.NewResponse(c)
	// 获取表单文件
	file, fileHeader, err := c.Request.FormFile("video")
	if err != nil {
		resp.ToErrorResponse(errors.New("获取上传文件失败"))
		return
	}
	defer file.Close()

	// 检查文件大小
	if fileHeader.Size > 100*1024*1024 { // 限制为100MB
		resp.ToErrorResponse(errors.New("文件大小不能超过100MB"))
		return
	}

	// 检查文件类型
	ext := filepath.Ext(fileHeader.Filename)
	allowedExts := map[string]bool{
		".mp4":  true,
		".avi":  true,
		".mov":  true,
		".mkv":  true,
		".flv":  true,
		".webm": true,
	}
	if !allowedExts[ext] {
		resp.ToErrorResponse(errors.New("不支持的视频格式"))
		return
	}

	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 解析表单参数
	var req dto.SendVideoMessageRequest
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	err = service.NewMessageService(c).SendVideoMessage(c, req, file, fileHeader, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Message) SendVoiceMessage(c *gin.Context) {
	resp := appx.NewResponse(c)
	// 获取表单文件
	file, fileHeader, err := c.Request.FormFile("voice")
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
		".amr": true,
		".mp3": true,
		".wav": true,
	}
	if !allowedExts[ext] {
		resp.ToErrorResponse(errors.New("不支持的音频格式"))
		return
	}

	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 解析表单参数
	var req dto.SendVoiceMessageRequest
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	err = service.NewMessageService(c).SendVoiceMessage(c, req, file, fileHeader, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Message) SendFileMessage(c *gin.Context) {
	resp := appx.NewResponse(c)
	// 取得分片内容
	file, fileHeader, err := c.Request.FormFile("chunk")
	if err != nil {
		resp.ToErrorResponse(errors.New("获取上传文件失败"))
		return
	}
	defer file.Close()

	if fileHeader.Size > vars.UploadFileChunkSize {
		resp.ToErrorResponse(errors.New("单个分片大小不能超过50KB"))
		return
	}

	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	var req dto.SendFileMessageRequest
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}

	// 基本参数校验（额外）
	if req.ChunkIndex < 0 || req.TotalChunks <= 0 || req.FileSize <= 0 || req.ChunkIndex >= req.TotalChunks {
		resp.ToErrorResponse(errors.New("分片参数错误"))
		return
	}
	if len(req.FileHash) == 0 || len(req.Filename) == 0 {
		resp.ToErrorResponse(errors.New("缺少文件信息"))
		return
	}

	if err = service.NewMessageService(c).SendFileMessage(c, req, file, robot); err != nil {
		resp.ToErrorResponse(err)
		return
	}

	resp.ToResponse(nil)
}

func (ct *Message) GetTimbre(c *gin.Context) {
	resp := appx.NewResponse(c)
	timbre, err := service.NewMessageService(c).GetTimbre()
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(timbre)
}

func (ct *Message) SendAITTSMessage(c *gin.Context) {
	var req dto.RobotSendAITTSMessageRequest
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
	if utf8.RuneCountInString(req.Content) > 260 {
		resp.ToErrorResponse(errors.New("内容长度不能超过260个字符"))
		return
	}
	err = service.NewMessageService(c).SendAITTSMessage(c, req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
