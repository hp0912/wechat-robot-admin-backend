package controller

import (
	"errors"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/service"

	"github.com/gin-gonic/gin"
)

type Contact struct {
}

func NewContactController() *Contact {
	return &Contact{}
}

func (ct *Contact) SyncContacts(c *gin.Context) {
	var req dto.SyncContactsRequest
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
	service.NewContactService(c).SyncContact(robot)
	resp.ToResponse(nil)
}

func (ct *Contact) GetContacts(c *gin.Context) {
	var req dto.GetContactsRequest
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
	pager := appx.InitPager(c)
	data, total, err := service.NewContactService(c).GetContacts(req, pager, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponseList(data, total)
}

func (ct *Contact) FriendSearch(c *gin.Context) {
	var req dto.FriendSearchRequest
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
	friend, err := service.NewContactService(c).FriendSearch(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(friend)
}

func (ct *Contact) FriendSendRequest(c *gin.Context) {
	var req dto.FriendSendRequestRequest
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
	err = service.NewContactService(c).FriendSendRequest(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Contact) FriendSendRequestFromChatRoom(c *gin.Context) {
	var req dto.FriendSendRequestFromChatRoomRequest
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
	err = service.NewContactService(c).FriendSendRequestFromChatRoom(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Contact) FriendSetRemarks(c *gin.Context) {
	var req dto.FriendSetRemarksRequest
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
	err = service.NewContactService(c).FriendSetRemarks(req, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Contact) FriendPassVerify(c *gin.Context) {
	var req dto.FriendPassVerifyRequest
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
	err = service.NewContactService(c).FriendPassVerify(req.SystemMessageID, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}

func (ct *Contact) FriendDelete(c *gin.Context) {
	var req dto.FriendDeleteRequest
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
	err = service.NewContactService(c).FriendDelete(req.ContactID, robot)
	if err != nil {
		resp.ToErrorResponse(err)
		return
	}
	resp.ToResponse(nil)
}
