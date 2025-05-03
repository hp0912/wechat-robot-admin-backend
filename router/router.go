package router

import (
	"wechat-robot-admin-backend/controller"
	"wechat-robot-admin-backend/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat
var userCtl *controller.User
var robotCtl *controller.Robot
var contactCtl *controller.Contact
var chatRoomCtl *controller.ChatRoom
var chatHistoryCtl *controller.ChatHistory
var systemCtl *controller.System

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
	userCtl = controller.NewUserController()
	robotCtl = controller.NewRobotController()
	contactCtl = controller.NewContactController()
	systemCtl = controller.NewSystemController()
	chatRoomCtl = controller.NewChatRoomController()
	chatHistoryCtl = controller.NewChatHistoryController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)

	initController()

	api := r.Group("/api/v1")
	{
		oauth := api.Group("/oauth")
		oauth.POST("/wechat", wechatCtl.WeChatAuth)
	}

	{
		user := api.Group("/user")
		user.GET("/self", userCtl.LoginUser)
		user.DELETE("/logout", userCtl.Logout)
	}

	{
		system := api.Group("/system")
		system.Use(middleware.UserAuth())
		system.GET("/robot-container-stats", middleware.UserOwnerAuth(), systemCtl.RobotContainerStats)
		system.GET("/robot-container-logs", middleware.UserOwnerAuth(), systemCtl.GetRobotContainerLogs)
	}

	{
		contact := api.Group("/contact")
		contact.Use(middleware.UserAuth())
		contact.GET("/list", middleware.UserOwnerAuth(), contactCtl.GetContacts)
		contact.POST("/sync", middleware.UserOwnerAuth(), contactCtl.SyncContacts)
	}

	{
		chatRoom := api.Group("/chat-room")
		chatRoom.Use(middleware.UserAuth())
		chatRoom.GET("/members", middleware.UserOwnerAuth(), chatRoomCtl.GetChatRoomMembers)
		chatRoom.POST("/members/sync", middleware.UserOwnerAuth(), chatRoomCtl.SyncChatRoomMembers)
	}

	{
		chat := api.Group("/chat")
		chat.Use(middleware.UserAuth())
		chat.GET("/history", middleware.UserOwnerAuth(), chatHistoryCtl.GetChatRoomMembers)
	}

	{
		robot := api.Group("/robot")
		robot.Use(middleware.UserAuth())

		robot.GET("/list", robotCtl.RobotList)
		robot.POST("/create", robotCtl.RobotCreate)
		robot.GET("/view", middleware.UserOwnerAuth(), robotCtl.RobotView)
		robot.POST("/restart-client", middleware.UserOwnerAuth(), robotCtl.RobotRestartClient)
		robot.POST("/restart-server", middleware.UserOwnerAuth(), robotCtl.RobotRestartServer)
		robot.DELETE("/remove", middleware.UserOwnerAuth(), robotCtl.RobotRemove)
		// 机器人登陆、登出
		robot.GET("/state", middleware.UserOwnerAuth(), robotCtl.RobotState)
		robot.POST("/login", middleware.UserOwnerAuth(), robotCtl.RobotLogin)
		robot.POST("/login-check", middleware.UserOwnerAuth(), robotCtl.RobotLoginCheck)
		robot.DELETE("/logout", middleware.UserOwnerAuth(), robotCtl.RobotLogout)
	}

	return nil
}
