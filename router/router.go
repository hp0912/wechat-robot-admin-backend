package router

import (
	"wechat-robot-admin-backend/controller"
	"wechat-robot-admin-backend/middleware"

	"github.com/gin-gonic/gin"
)

var probeCtl *controller.Probe
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
var systemMessageCtl *controller.SystemMessage
var momentsCtl *controller.Moments
var systemSettingsCtl *controller.SystemSettings
var ossSettingsCtl *controller.OSSSettings
var mcpServerCtl *controller.MCPServer
var wxAppCtl *controller.WXApp
var aiCallbackCtl *controller.AICallback
var pprofProxyCtl *controller.PprofProxy

func initController() {
	probeCtl = controller.NewProbeController()
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
	systemMessageCtl = controller.NewSystemMessageController()
	systemSettingsCtl = controller.NewSystemSettingsController()
	ossSettingsCtl = controller.NewOSSSettingsController()
	mcpServerCtl = controller.NewMCPController()
	wxAppCtl = controller.NewWXAppController()
	pprofProxyCtl = controller.NewPprofProxyController()
}

func RegisterRouter(r *gin.Engine) error {
	r.Use(middleware.ErrorRecover)
	r.Use(middleware.Cors)

	initController()

	api := r.Group("/api/v1")
	{
		probe := api.Group("/probe")
		probe.GET("", probeCtl.Probe)
	}

	{
		user := api.Group("/login")
		user.POST("", userCtl.Login)
	}

	{
		oauth := api.Group("/oauth")
		oauth.GET("/official-account/url", wechatCtl.WeChatOfficialAccountAuthURL)
		oauth.POST("/wechat", wechatCtl.WeChatAuth)
	}

	{
		user := api.Group("/user")
		user.GET("/self", userCtl.LoginUser)
		user.POST("/api-token/refresh", middleware.UserAuth(), userCtl.RefreshUserApiToken)
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
		contact.POST("/friend/search", middleware.UserOwnerAuth(), contactCtl.FriendSearch)
		contact.POST("/friend/add", middleware.UserOwnerAuth(), contactCtl.FriendSendRequest)
		contact.POST("/friend/add-from-chat-room", middleware.UserOwnerAuth(), contactCtl.FriendSendRequestFromChatRoom)
		contact.POST("/friend/remark", middleware.UserOwnerAuth(), contactCtl.FriendSetRemarks)
		contact.POST("/friend/pass-verify", middleware.UserOwnerAuth(), contactCtl.FriendPassVerify)
		contact.POST("/sync", middleware.UserOwnerAuth(), contactCtl.SyncContacts)
		contact.DELETE("/friend", middleware.UserOwnerAuth(), contactCtl.FriendDelete)
	}

	{
		chatRoom := api.Group("/chat-room")
		chatRoom.Use(middleware.UserAuth())
		chatRoom.GET("/members", middleware.UserOwnerAuth(), chatRoomCtl.GetChatRoomMembers)
		chatRoom.GET("/not-left-members", middleware.UserOwnerAuth(), chatRoomCtl.GetNotLeftMembers)
		chatRoom.POST("/create", middleware.UserOwnerAuth(), chatRoomCtl.CreateChatRoom)
		chatRoom.POST("/invite", middleware.UserOwnerAuth(), chatRoomCtl.InviteChatRoomMember)
		chatRoom.POST("/join", middleware.UserOwnerAuth(), chatRoomCtl.GroupConsentToJoin)
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
		message.POST("/send/file", middleware.UserOwnerAuth(), messageCtl.SendFileMessage)
		message.GET("/timbre", middleware.UserOwnerAuth(), messageCtl.GetTimbre)
		message.POST("/send/ai/tts", middleware.UserOwnerAuth(), messageCtl.SendAITTSMessage)
	}

	{
		systemMessage := api.Group("/system-messages")
		systemMessage.Use(middleware.UserAuth())
		systemMessage.GET("", middleware.UserOwnerAuth(), systemMessageCtl.GetRecentMonthMessages)
		systemMessage.POST("/mark-as-read", middleware.UserOwnerAuth(), systemMessageCtl.MarkAsReadBatch)
	}

	{
		systemSettings := api.Group("/system-settings")
		systemSettings.Use(middleware.UserAuth())
		systemSettings.GET("", middleware.UserOwnerAuth(), systemSettingsCtl.GetSystemSettings)
		systemSettings.POST("", middleware.UserOwnerAuth(), systemSettingsCtl.SaveSystemSettings)
	}

	{
		ossSettings := api.Group("/oss-settings")
		ossSettings.Use(middleware.UserAuth())
		ossSettings.GET("", middleware.UserOwnerAuth(), ossSettingsCtl.GetOSSSettings)
		ossSettings.POST("", middleware.UserOwnerAuth(), ossSettingsCtl.SaveOSSSettings)
	}

	{
		mcpServer := api.Group("/mcp-server")
		mcpServer.Use(middleware.UserAuth())
		mcpServer.GET("", middleware.UserOwnerAuth(), mcpServerCtl.GetMCPServer)
		mcpServer.GET("/list", middleware.UserOwnerAuth(), mcpServerCtl.GetMCPServers)
		mcpServer.GET("/tools", middleware.UserOwnerAuth(), mcpServerCtl.GetMCPServerTools)
		mcpServer.POST("", middleware.UserOwnerAuth(), mcpServerCtl.CreateMCPServer)
		mcpServer.POST("/enable", middleware.UserOwnerAuth(), mcpServerCtl.EnableMCPServer)
		mcpServer.POST("/disable", middleware.UserOwnerAuth(), mcpServerCtl.DisableMCPServer)
		mcpServer.PUT("", middleware.UserOwnerAuth(), mcpServerCtl.UpdateMCPServer)
		mcpServer.DELETE("", middleware.UserOwnerAuth(), mcpServerCtl.DeleteMCPServer)
	}

	{
		moments := api.Group("/moments")
		moments.Use(middleware.UserAuth())
		moments.GET("/list", middleware.UserOwnerAuth(), momentsCtl.FriendCircleGetList)
		moments.GET("/settings", middleware.UserOwnerAuth(), momentsCtl.GetFriendCircleSettings)
		moments.GET("/get-detail", middleware.UserOwnerAuth(), momentsCtl.FriendCircleGetDetail)
		moments.GET("/get-id-detail", middleware.UserOwnerAuth(), momentsCtl.FriendCircleGetIdDetail)
		moments.GET("/down-media", middleware.UserOwnerAuth(), momentsCtl.FriendCircleDownFriendCircleMedia)
		moments.POST("/settings", middleware.UserOwnerAuth(), momentsCtl.SaveFriendCircleSettings)
		moments.POST("/comment", middleware.UserOwnerAuth(), momentsCtl.FriendCircleComment)
		moments.POST("/upload-media", middleware.UserOwnerAuth(), momentsCtl.FriendCircleUpload)
		moments.POST("/post", middleware.UserOwnerAuth(), momentsCtl.FriendCirclePost)
		moments.POST("/operate", middleware.UserOwnerAuth(), momentsCtl.FriendCircleOperation)
		moments.POST("/privacy-settings", middleware.UserOwnerAuth(), momentsCtl.FriendCirclePrivacySettings)
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
		robot.GET("/export-login-data", middleware.UserOwnerAuth(), robotManageCtl.ExportRobotLoginData)
		robot.POST("/import-login-data", middleware.UserOwnerAuth(), robotManageCtl.ImportRobotLoginData)
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
		robot.POST("/login/slider-auto", robotLoginCtl.LoginSliderAutoVerify)           // 不需要 middleware.UserOwnerAuth() 鉴权
		robot.GET("/login/slider", robotLoginCtl.LoginSliderVerify)                     // 不需要 middleware.UserOwnerAuth() 鉴权
		robot.GET("/login/slider-verify-submit", robotLoginCtl.LoginSliderVerifySubmit) // 不需要 middleware.UserOwnerAuth() 鉴权
		robot.POST("/login", middleware.UserOwnerAuth(), robotLoginCtl.RobotLogin)
		robot.POST("/login-check", middleware.UserOwnerAuth(), robotLoginCtl.RobotLoginCheck)
		robot.POST("/login/2fa", middleware.UserOwnerAuth(), robotLoginCtl.RobotLogin2FA)
		robot.POST("/login/data62", middleware.UserOwnerAuth(), robotLoginCtl.LoginData62Login)
		robot.POST("/login/data62-sms-again", middleware.UserOwnerAuth(), robotLoginCtl.LoginData62SMSAgain)
		robot.POST("/login/data62-sms-verify", middleware.UserOwnerAuth(), robotLoginCtl.LoginData62SMSVerify)
		robot.POST("/login/a16", middleware.UserOwnerAuth(), robotLoginCtl.LoginA16Data1)
		robot.DELETE("/logout", middleware.UserOwnerAuth(), robotLoginCtl.RobotLogout)
	}

	{
		robot := api.Group("/wxapp")
		robot.Use(middleware.UserAuth())

		robot.POST("/qrcode-auth-login", middleware.UserOwnerAuth(), wxAppCtl.WxappQrcodeAuthLogin)
	}

	{
		// pprof代理路由
		pprof := api.Group("/pprof")
		pprof.Use(middleware.UserAuth())
		pprof.GET("/*pprof_path", middleware.UserOwnerAuth(), pprofProxyCtl.ProxyPprof)
	}

	return nil
}
