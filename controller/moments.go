package controller

import (
	"encoding/base64"
	"errors"
	"net/http"
	"path/filepath"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type Moments struct{}

func NewMomentsController() *Moments {
	return &Moments{}
}

func (ct *Moments) FriendCircleGetList(c *gin.Context) {
	var req dto.MomentsGetListRequest
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
	data, err := service.NewMomentsService(c).FriendCircleGetList(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) GetFriendCircleSettings(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	data, err := service.NewMomentsService(c).GetFriendCircleSettings(robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) SaveFriendCircleSettings(c *gin.Context) {
	var req dto.MomentSettings
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
	err = service.NewMomentsService(c).SaveFriendCircleSettings(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Moments) FriendCircleGetDetail(c *gin.Context) {
	var req dto.FriendCircleGetDetailRequest
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
	data, err := service.NewMomentsService(c).FriendCircleGetDetail(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCircleGetIdDetail(c *gin.Context) {
	var req dto.FriendCircleGetIdDetailRequest
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
	data, err := service.NewMomentsService(c).FriendCircleGetIdDetail(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCircleComment(c *gin.Context) {
	var req dto.FriendCircleCommentRequest
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
	data, err := service.NewMomentsService(c).FriendCircleComment(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCircleDownFriendCircleMedia(c *gin.Context) {
	var req dto.MomentsDownFriendCircleMediaRequest
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
	data, err := service.NewMomentsService(c).FriendCircleDownFriendCircleMedia(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	videoBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode video")
		return
	}
	c.Data(http.StatusOK, "video/mp4", videoBytes)
}

func (ct *Moments) FriendCircleUpload(c *gin.Context) {
	resp := appx.NewResponse(c)
	robot, err := appx.GetRobot(c)
	if err != nil {
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	// 获取表单文件
	file, fileHeader, err := c.Request.FormFile("media")
	if err != nil {
		resp.ToErrorResponse(errors.New("获取上传文件失败"))
		return
	}
	defer file.Close()

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
		allowedExts = map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".webp": true,
		}
		if !allowedExts[ext] {
			resp.ToErrorResponse(errors.New("不支持的图片/视频格式"))
			return
		} else {
			// 检查文件大小
			if fileHeader.Size > 50*1024*1024 { // 限制为50MB
				resp.ToErrorResponse(errors.New("图片大小不能超过50MB"))
				return
			}
		}
	} else {
		// 检查文件大小
		if fileHeader.Size > 100*1024*1024 { // 限制为100MB
			resp.ToErrorResponse(errors.New("视频大小不能超过100MB"))
			return
		}
	}

	data, err := service.NewMomentsService(c).FriendCircleUpload(file, fileHeader, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCirclePost(c *gin.Context) {
	var req dto.MomentPostRequest
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
	data, err := service.NewMomentsService(c).FriendCirclePost(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCircleOperation(c *gin.Context) {
	var req dto.MomentOpRequest
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
	data, err := service.NewMomentsService(c).FriendCircleOperation(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}

func (ct *Moments) FriendCirclePrivacySettings(c *gin.Context) {
	var req dto.MomentPrivacySettingsRequest
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
	data, err := service.NewMomentsService(c).FriendCirclePrivacySettings(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(data)
}
