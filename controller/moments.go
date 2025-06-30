package controller

import (
	"encoding/base64"
	"errors"
	"net/http"
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
