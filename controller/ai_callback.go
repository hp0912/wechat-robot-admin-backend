package controller

import (
	"errors"
	"log"
	"strconv"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type AICallback struct {
}

func NewAICallbackController() *AICallback {
	return &AICallback{}
}

func (ct *AICallback) DoubaoTTS(c *gin.Context) {
	var req dto.DoubaoTTSCallbackRequest
	resp := appx.NewResponse(c)
	if ok, err := appx.BindAndValid(c, &req); !ok || err != nil {
		log.Println("DoubaoTTS 回调参数错误:", err)
		resp.ToErrorResponse(errors.New("参数错误"))
		return
	}
	robotIdStr := c.Query("robot_id")
	robotId, err := strconv.ParseInt(robotIdStr, 10, 64)
	if err != nil {
		log.Println("DoubaoTTS 回调参数 robot_id 错误:", err)
		resp.ToErrorResponse(err)
		return
	}
	err = service.NewAICallbackService(c).DoubaoTTS(req, robotId)
	if err != nil {
		log.Println("DoubaoTTS 回调处理失败:", err)
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
