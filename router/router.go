package router

import (
	"wechat-robot-admin-backend/controller"
	"wechat-robot-admin-backend/middleware"

	"github.com/gin-gonic/gin"
)

var wechatCtl *controller.WeChat
var userCtl *controller.User
var robotLoginCtl *controller.RobotLogin
var robotManageCtl *controller.RobotManage
var contactCtl *controller.Contact
var chatRoomCtl *controller.ChatRoom
var chatHistoryCtl *controller.ChatHistory
var dockerCtl *controller.DockerController
var globalSettingsCtl *controller.GlobalSettings
var friendSettingsCtl *controller.FriendSettings
var chatRoomSettingsCtl *controller.ChatRoomSettings
var messageCtl *controller.Message
var momentsCtl *controller.Moments
var aiCallbackCtl *controller.AICallback

func initController() {
	wechatCtl = controller.NewWeChatAuthController()
	userCtl = controller.NewUserController()
	robotLoginCtl = controller.NewRobotLoginController()
	robotManageCtl = controller.NewRobotManageController()
	contactCtl = controller.NewContactController()
	dockerCtl = controller.NewDockerController()
	chatRoomCtl = controller.NewChatRoomController()
	chatHistoryCtl = controller.NewChatHistoryController()
	messageCtl = controller.NewMessageController()
	globalSettingsCtl = controller.NewGlobalSettingsController()
	friendSettingsCtl = controller.NewFriendSettingsController()
	chatRoomSettingsCtl = controller.NewChatRoomSettingsController()
	aiCallbackCtl = controller.NewAICallbackController()
	momentsCtl = controller.NewMomentsController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)
	r.Use(middleware.Cors)

	initController()

	api := r.Group("/api/v1")
	{
		oauth := api.Group("/oauth")
		oauth.GET("/official-account/url", wechatCtl.WeChatOfficialAccountAuthURL)
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
		system.GET("/robot-container-stats", middleware.UserOwnerAuth(), dockerCtl.RobotContainerStats)
		system.GET("/robot-container-logs", middleware.UserOwnerAuth(), dockerCtl.GetRobotContainerLogs)
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
		chatRoom.POST("/name", middleware.UserOwnerAuth(), chatRoomCtl.GroupSetChatRoomName)
		chatRoom.POST("/remark", middleware.UserOwnerAuth(), chatRoomCtl.GroupSetChatRoomRemarks)
		chatRoom.POST("/announcement", middleware.UserOwnerAuth(), chatRoomCtl.GroupSetChatRoomAnnouncement)
		chatRoom.DELETE("/members", middleware.UserOwnerAuth(), chatRoomCtl.GroupDelChatRoomMember)
		chatRoom.DELETE("/quit", middleware.UserOwnerAuth(), chatRoomCtl.GroupQuit)
	}

	{
		chat := api.Group("/chat")
		chat.Use(middleware.UserAuth())
		chat.GET("/history", middleware.UserOwnerAuth(), chatHistoryCtl.GetChatRoomMembers)
		chat.GET("/image/download", middleware.UserOwnerAuth(), chatHistoryCtl.DownloadImage)
		chat.GET("/voice/download", middleware.UserOwnerAuth(), chatHistoryCtl.DownloadVoice)
		chat.GET("/file/download", middleware.UserOwnerAuth(), chatHistoryCtl.DownloadFile)
		chat.GET("/video/download", middleware.UserOwnerAuth(), chatHistoryCtl.DownloadVideo)
	}

	{
		message := api.Group("/message")
		message.Use(middleware.UserAuth())
		message.POST("/revoke", middleware.UserOwnerAuth(), messageCtl.MessageRevoke)
		message.POST("/send/text", middleware.UserOwnerAuth(), messageCtl.SendTextMessage)
		message.POST("/send/image", middleware.UserOwnerAuth(), messageCtl.SendImageMessage)
		message.POST("/send/voice", middleware.UserOwnerAuth(), messageCtl.SendVoiceMessage)
		message.POST("/send/video", middleware.UserOwnerAuth(), messageCtl.SendVideoMessage)
		message.GET("/timbre", middleware.UserOwnerAuth(), messageCtl.GetTimbre)
		message.POST("/send/ai/tts", middleware.UserOwnerAuth(), messageCtl.SendAITTSMessage)
	}

	{
		moments := api.Group("/moments")
		moments.Use(middleware.UserAuth())
		moments.GET("/list", middleware.UserOwnerAuth(), momentsCtl.FriendCircleGetList)
		moments.GET("/down-media", middleware.UserOwnerAuth(), momentsCtl.FriendCircleDownFriendCircleMedia)
	}

	{
		globalSettings := api.Group("/global-settings")
		globalSettings.Use(middleware.UserAuth())
		globalSettings.GET("", middleware.UserOwnerAuth(), globalSettingsCtl.GetGlobalSettings)
		globalSettings.POST("", middleware.UserOwnerAuth(), globalSettingsCtl.SaveGlobalSettings)
	}

	{
		friendSettings := api.Group("/friend-settings")
		friendSettings.Use(middleware.UserAuth())
		friendSettings.GET("", middleware.UserOwnerAuth(), friendSettingsCtl.GetFriendSettings)
		friendSettings.POST("", middleware.UserOwnerAuth(), friendSettingsCtl.SaveFriendSettings)
	}

	{
		chatRoomSettings := api.Group("/chat-room-settings")
		chatRoomSettings.Use(middleware.UserAuth())
		chatRoomSettings.GET("", middleware.UserOwnerAuth(), chatRoomSettingsCtl.GetChatRoomSettings)
		chatRoomSettings.POST("", middleware.UserOwnerAuth(), chatRoomSettingsCtl.SaveChatRoomSettings)
	}

	{
		aiCallback := api.Group("/ai-callback")
		aiCallback.POST("/voice/doubao-tts", aiCallbackCtl.DoubaoTTS)
	}

	{
		robot := api.Group("/robot")
		robot.Use(middleware.UserAuth())

		robot.GET("/list", robotManageCtl.RobotList)
		robot.POST("/create", robotManageCtl.RobotCreate)
		robot.GET("/view", middleware.UserOwnerAuth(), robotManageCtl.RobotView)
		robot.POST("/restart-client", middleware.UserOwnerAuth(), robotManageCtl.RobotRestartClient)
		robot.POST("/restart-server", middleware.UserOwnerAuth(), robotManageCtl.RobotRestartServer)
		// 拉取最新镜像
		robot.GET("/docker/image/pull", middleware.UserOwnerAuth(), robotManageCtl.RobotDockerImagePull)
		// 删除客户端和服务端容器
		robot.DELETE("/docker/container/client/remove", middleware.UserOwnerAuth(), robotManageCtl.RobotStopAndRemoveClient)
		robot.DELETE("/docker/container/server/remove", middleware.UserOwnerAuth(), robotManageCtl.RobotStopAndRemoveServer)
		// 启动客户端和服务端容器
		robot.POST("/docker/container/client/start", middleware.UserOwnerAuth(), robotManageCtl.RobotStartClient)
		robot.POST("/docker/container/server/start", middleware.UserOwnerAuth(), robotManageCtl.RobotStartServer)
		// 删除机器人
		robot.DELETE("/remove", middleware.UserOwnerAuth(), robotManageCtl.RobotRemove)
		// 机器人登陆、登出
		robot.GET("/state", middleware.UserOwnerAuth(), robotLoginCtl.RobotState)
		robot.POST("/login", middleware.UserOwnerAuth(), robotLoginCtl.RobotLogin)
		robot.POST("/login-check", middleware.UserOwnerAuth(), robotLoginCtl.RobotLoginCheck)
		robot.DELETE("/logout", middleware.UserOwnerAuth(), robotLoginCtl.RobotLogout)
	}

	return nil
}
