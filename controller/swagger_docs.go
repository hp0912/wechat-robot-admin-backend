package controller

import (
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
)

var (
	_ = dto.Response[dto.SwaggerEmpty]{}
	_ = model.User{}
)

type SwaggerEmpty struct{}

type SwaggerSuccess struct {
	Success bool `json:"success"`
}

type SwaggerLoginRequest struct {
	Token string `json:"token" binding:"required"`
}

type SwaggerRobotLoginRequest struct {
	LoginType   string `json:"login_type" binding:"required"`
	IsPretender *bool  `json:"is_pretender" binding:"required"`
}

type SwaggerImportRobotLoginDataRequest struct {
	Data string `json:"data" binding:"required"`
}

type SwaggerIDRequest struct {
	ID int64 `json:"id" binding:"required"`
}

// ProbeDoc godoc
// @Summary 健康检查
// @Tags Probe
// @Produce json
// @Success 200 {object} SwaggerSuccess
// @Router /probe [get]
func ProbeDoc() {}

// LoginDoc godoc
// @Summary 管理后台登录
// @Tags User
// @Accept json
// @Produce json
// @Param request body SwaggerLoginRequest true "登录参数"
// @Success 200 {object} dto.Response[dto.SwaggerSuccess]
// @Router /login [post]
func LoginDoc() {}

// WeChatOfficialAccountAuthURLDoc godoc
// @Summary 获取微信公众号授权二维码
// @Tags OAuth
// @Produce image/png
// @Success 200 {file} binary
// @Router /oauth/official-account/url [get]
func WeChatOfficialAccountAuthURLDoc() {}

// WeChatAuthDoc godoc
// @Summary 微信 OAuth 登录
// @Tags OAuth
// @Accept json
// @Produce json
// @Param request body WeChatRequest true "授权 code"
// @Success 200 {object} dto.Response[dto.SwaggerSuccess]
// @Router /oauth/wechat [post]
func WeChatAuthDoc() {}

// LoginUserDoc godoc
// @Summary 获取当前登录用户
// @Tags User
// @Produce json
// @Success 200 {object} dto.Response[model.User]
// @Router /user/self [get]
func LoginUserDoc() {}

// RefreshUserApiTokenDoc godoc
// @Summary 刷新当前用户 API Token
// @Tags User
// @Security ApiTokenAuth
// @Produce json
// @Success 200 {object} dto.Response[string]
// @Router /user/api-token/refresh [post]
func RefreshUserApiTokenDoc() {}

// LogoutDoc godoc
// @Summary 退出登录
// @Tags User
// @Produce json
// @Success 200 {object} dto.Response[model.User]
// @Router /user/logout [delete]
func LogoutDoc() {}

// RobotContainerStatsDoc godoc
// @Summary 获取机器人容器资源统计
// @Tags System
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.RobotContainerStatsResponse]
// @Router /system/robot-container-stats [get]
func RobotContainerStatsDoc() {}

// GetRobotContainerLogsDoc godoc
// @Summary 获取机器人容器日志
// @Tags System
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.RobotContainerLogsResponse]
// @Router /system/robot-container-logs [get]
func GetRobotContainerLogsDoc() {}

// GetContactsDoc godoc
// @Summary 获取联系人列表
// @Tags Contact
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param contact_ids query []string false "联系人ID列表"
// @Param type query string false "联系人类型"
// @Param keyword query string false "关键词"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.GetContactsResponse]]
// @Router /contact/list [get]
func GetContactsDoc() {}

// FriendSearchDoc godoc
// @Summary 搜索好友
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendSearchRequest true "搜索参数"
// @Success 200 {object} dto.Response[dto.FriendSearchResponse]
// @Router /contact/friend/search [post]
func FriendSearchDoc() {}

// FriendSendRequestDoc godoc
// @Summary 发送好友申请
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendSendRequestRequest true "好友申请参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/friend/add [post]
func FriendSendRequestDoc() {}

// FriendSendRequestFromChatRoomDoc godoc
// @Summary 从群成员发送好友申请
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendSendRequestFromChatRoomRequest true "好友申请参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/friend/add-from-chat-room [post]
func FriendSendRequestFromChatRoomDoc() {}

// FriendSetRemarksDoc godoc
// @Summary 设置好友备注
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendSetRemarksRequest true "备注参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/friend/remark [post]
func FriendSetRemarksDoc() {}

// FriendPassVerifyDoc godoc
// @Summary 通过好友验证
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendPassVerifyRequest true "验证参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/friend/pass-verify [post]
func FriendPassVerifyDoc() {}

// SyncContactsDoc godoc
// @Summary 同步联系人
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SyncContactsRequest false "同步参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/sync [post]
func SyncContactsDoc() {}

// FriendDeleteDoc godoc
// @Summary 删除好友
// @Tags Contact
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendDeleteRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /contact/friend [delete]
func FriendDeleteDoc() {}

// GetChatRoomMembersDoc godoc
// @Summary 获取群成员列表
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param chat_room_id query string true "群ID"
// @Param keyword query string false "关键词"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.ChatRoomMember]]
// @Router /chat-room/members [get]
func GetChatRoomMembersDoc() {}

// GetChatRoomMemberDoc godoc
// @Summary 获取群成员详情
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param chat_room_id query string true "群ID"
// @Param wechat_id query string true "成员微信ID"
// @Success 200 {object} dto.Response[dto.ChatRoomMember]
// @Router /chat-room/member [get]
func GetChatRoomMemberDoc() {}

// UpdateChatRoomMemberDoc godoc
// @Summary 更新群成员配置
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.UpdateChatRoomMemberRequest true "群成员配置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/member [post]
func UpdateChatRoomMemberDoc() {}

// GetNotLeftMembersDoc godoc
// @Summary 获取未退群成员列表
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param chat_room_id query string true "群ID"
// @Param keyword query string false "关键词"
// @Success 200 {object} dto.Response[[]dto.ChatRoomMember]
// @Router /chat-room/not-left-members [get]
func GetNotLeftMembersDoc() {}

// CreateChatRoomDoc godoc
// @Summary 创建群聊
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.CreateChatRoomRequest true "建群参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/create [post]
func CreateChatRoomDoc() {}

// InviteChatRoomMemberDoc godoc
// @Summary 邀请群成员
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.InviteChatRoomMemberRequest true "邀请参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/invite [post]
func InviteChatRoomMemberDoc() {}

// GroupConsentToJoinDoc godoc
// @Summary 同意入群邀请
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomJoinRequest true "入群参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/join [post]
func GroupConsentToJoinDoc() {}

// SyncChatRoomMembersDoc godoc
// @Summary 同步群成员
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomRequestBase true "同步参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/members/sync [post]
func SyncChatRoomMembersDoc() {}

// GroupSetChatRoomNameDoc godoc
// @Summary 设置群名称
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomOperateRequest true "群名称参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/name [post]
func GroupSetChatRoomNameDoc() {}

// GroupSetChatRoomRemarksDoc godoc
// @Summary 设置群备注
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomOperateRequest true "群备注参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/remark [post]
func GroupSetChatRoomRemarksDoc() {}

// GroupSetChatRoomAnnouncementDoc godoc
// @Summary 设置群公告
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomOperateRequest true "群公告参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/announcement [post]
func GroupSetChatRoomAnnouncementDoc() {}

// GroupDelChatRoomMemberDoc godoc
// @Summary 删除群成员
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DelChatRoomMemberRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/members [delete]
func GroupDelChatRoomMemberDoc() {}

// GroupQuitDoc godoc
// @Summary 退出群聊
// @Tags ChatRoom
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.ChatRoomRequestBase true "退群参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room/quit [delete]
func GroupQuitDoc() {}

// GetChatHistoryDoc godoc
// @Summary 获取聊天记录
// @Tags Chat
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param contact_id query string true "联系人ID"
// @Param keyword query string false "关键词"
// @Param chat_room_member query string false "群成员微信ID"
// @Param time_start query int false "开始时间"
// @Param time_end query int false "结束时间"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.ChatHistory]]
// @Router /chat/history [get]
func GetChatHistoryDoc() {}

// DownloadImageDoc godoc
// @Summary 下载聊天图片
// @Tags Chat
// @Security ApiTokenAuth
// @Produce application/octet-stream
// @Param id query int true "机器人实例ID"
// @Param message_id query int true "消息ID"
// @Param attach_url query string false "附件地址"
// @Success 200 {file} binary
// @Router /chat/image/download [get]
func DownloadImageDoc() {}

// DownloadVoiceDoc godoc
// @Summary 下载聊天语音
// @Tags Chat
// @Security ApiTokenAuth
// @Produce application/octet-stream
// @Param id query int true "机器人实例ID"
// @Param message_id query int true "消息ID"
// @Param attach_url query string false "附件地址"
// @Success 200 {file} binary
// @Router /chat/voice/download [get]
func DownloadVoiceDoc() {}

// DownloadFileDoc godoc
// @Summary 下载聊天文件
// @Tags Chat
// @Security ApiTokenAuth
// @Produce application/octet-stream
// @Param id query int true "机器人实例ID"
// @Param message_id query int true "消息ID"
// @Param attach_url query string false "附件地址"
// @Success 200 {file} binary
// @Router /chat/file/download [get]
func DownloadFileDoc() {}

// DownloadVideoDoc godoc
// @Summary 下载聊天视频
// @Tags Chat
// @Security ApiTokenAuth
// @Produce application/octet-stream
// @Param id query int true "机器人实例ID"
// @Param message_id query int true "消息ID"
// @Param attach_url query string false "附件地址"
// @Success 200 {file} binary
// @Router /chat/video/download [get]
func DownloadVideoDoc() {}

// MessageRevokeDoc godoc
// @Summary 撤回消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MessageRevokeRequest true "撤回参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/revoke [post]
func MessageRevokeDoc() {}

// SendTextMessageDoc godoc
// @Summary 发送文本消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SendTextMessageRequest true "文本消息"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/text [post]
func SendTextMessageDoc() {}

// SendLongTextMessageDoc godoc
// @Summary 发送长文本消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SendLongTextMessageRequest true "长文本消息"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/longtext [post]
func SendLongTextMessageDoc() {}

// SendImageMessageDoc godoc
// @Summary 发送图片消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param to_wxid formData string true "接收方微信ID"
// @Param image formData file true "图片文件"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/image [post]
func SendImageMessageDoc() {}

// SendVoiceMessageDoc godoc
// @Summary 发送语音消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param to_wxid formData string true "接收方微信ID"
// @Param voice formData file true "语音文件"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/voice [post]
func SendVoiceMessageDoc() {}

// SendVideoMessageDoc godoc
// @Summary 发送视频消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param to_wxid formData string true "接收方微信ID"
// @Param video formData file true "视频文件"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/video [post]
func SendVideoMessageDoc() {}

// SendFileMessageDoc godoc
// @Summary 发送文件分片
// @Tags Message
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param to_wxid formData string true "接收方微信ID"
// @Param client_app_data_id formData string true "客户端文件ID"
// @Param filename formData string true "文件名"
// @Param file_hash formData string true "文件哈希"
// @Param file_size formData int true "文件大小"
// @Param chunk_index formData int false "分片序号"
// @Param total_chunks formData int true "分片总数"
// @Param chunk formData file true "文件分片"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/file [post]
func SendFileMessageDoc() {}

// GetTimbreDoc godoc
// @Summary 获取语音音色列表
// @Tags Message
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[[]string]
// @Router /message/timbre [get]
func GetTimbreDoc() {}

// SendAITTSMessageDoc godoc
// @Summary 发送 AI 语音消息
// @Tags Message
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotSendAITTSMessageRequest true "AI 语音消息"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /message/send/ai/tts [post]
func SendAITTSMessageDoc() {}

// GetRecentMonthMessagesDoc godoc
// @Summary 获取近一个月系统消息
// @Tags SystemMessage
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[[]dto.SystemMessage]
// @Router /system-messages [get]
func GetRecentMonthMessagesDoc() {}

// MarkAsReadBatchDoc godoc
// @Summary 批量标记系统消息已读
// @Tags SystemMessage
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MarkAsReadBatchRequest true "系统消息ID列表"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /system-messages/mark-as-read [post]
func MarkAsReadBatchDoc() {}

// GetSystemSettingsDoc godoc
// @Summary 获取系统设置
// @Tags SystemSettings
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SystemSettings]
// @Router /system-settings [get]
func GetSystemSettingsDoc() {}

// SaveSystemSettingsDoc godoc
// @Summary 保存系统设置
// @Tags SystemSettings
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SystemSettingsRequest true "系统设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /system-settings [post]
func SaveSystemSettingsDoc() {}

// GetOSSSettingsDoc godoc
// @Summary 获取 OSS 设置
// @Tags OSSSettings
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.OSSSettings]
// @Router /oss-settings [get]
func GetOSSSettingsDoc() {}

// SaveOSSSettingsDoc godoc
// @Summary 保存 OSS 设置
// @Tags OSSSettings
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.OSSSettingsRequest true "OSS 设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /oss-settings [post]
func SaveOSSSettingsDoc() {}

// GetMCPServerDoc godoc
// @Summary 获取 MCP 服务详情
// @Tags MCPServer
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param mcp_server_id query int true "MCP 服务ID"
// @Success 200 {object} dto.Response[dto.MCPServer]
// @Router /mcp-server [get]
func GetMCPServerDoc() {}

// GetMCPServersDoc godoc
// @Summary 获取 MCP 服务列表
// @Tags MCPServer
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[[]dto.MCPServer]
// @Router /mcp-server/list [get]
func GetMCPServersDoc() {}

// GetMCPServerToolsDoc godoc
// @Summary 获取 MCP 服务工具列表
// @Tags MCPServer
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param mcp_server_id query int true "MCP 服务ID"
// @Success 200 {object} dto.Response[[]dto.MCPServerTool]
// @Router /mcp-server/tools [get]
func GetMCPServerToolsDoc() {}

// CreateMCPServerDoc godoc
// @Summary 创建 MCP 服务
// @Tags MCPServer
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MCPServer true "MCP 服务"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /mcp-server [post]
func CreateMCPServerDoc() {}

// EnableMCPServerDoc godoc
// @Summary 启用 MCP 服务
// @Tags MCPServer
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerIDRequest true "MCP 服务业务ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /mcp-server/enable [post]
func EnableMCPServerDoc() {}

// DisableMCPServerDoc godoc
// @Summary 禁用 MCP 服务
// @Tags MCPServer
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerIDRequest true "MCP 服务业务ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /mcp-server/disable [post]
func DisableMCPServerDoc() {}

// UpdateMCPServerDoc godoc
// @Summary 更新 MCP 服务
// @Tags MCPServer
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MCPServer true "MCP 服务"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /mcp-server [put]
func UpdateMCPServerDoc() {}

// DeleteMCPServerDoc godoc
// @Summary 删除 MCP 服务
// @Tags MCPServer
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MCPServer true "MCP 服务"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /mcp-server [delete]
func DeleteMCPServerDoc() {}

// GetSkillsDoc godoc
// @Summary 获取技能列表
// @Tags Skills
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[[]dto.Skill]
// @Router /skills [get]
func GetSkillsDoc() {}

// InstallSkillDoc godoc
// @Summary 安装技能
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.InstallSkillRequest true "安装参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/install [post]
func InstallSkillDoc() {}

// EnableSkillDoc godoc
// @Summary 启用技能
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SkillRequest true "技能参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/enable [post]
func EnableSkillDoc() {}

// DisableSkillDoc godoc
// @Summary 禁用技能
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SkillRequest true "技能参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/disable [post]
func DisableSkillDoc() {}

// UpdateSkillDoc godoc
// @Summary 更新技能
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SkillRequest true "技能参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/update [put]
func UpdateSkillDoc() {}

// UninstallSkillDoc godoc
// @Summary 卸载技能
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SkillRequest true "技能参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/uninstall [delete]
func UninstallSkillDoc() {}

// SetSkillEnvsDoc godoc
// @Summary 设置技能环境变量
// @Tags Skills
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SetSkillEnvsRequest true "环境变量"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /skills/envs [post]
func SetSkillEnvsDoc() {}

// ListSystemPromptsDoc godoc
// @Summary 获取系统提示词列表
// @Tags SystemPrompt
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param keyword query string false "关键词"
// @Success 200 {object} dto.Response[[]dto.SystemPrompt]
// @Router /system-prompts [get]
func ListSystemPromptsDoc() {}

// GetSystemPromptDoc godoc
// @Summary 获取系统提示词详情
// @Tags SystemPrompt
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param prompt_id query int true "系统提示词业务ID"
// @Success 200 {object} dto.Response[dto.SystemPrompt]
// @Router /system-prompts/detail [get]
func GetSystemPromptDoc() {}

// CreateSystemPromptDoc godoc
// @Summary 创建系统提示词
// @Tags SystemPrompt
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.CreateSystemPromptRequest true "系统提示词"
// @Success 200 {object} dto.Response[dto.SystemPrompt]
// @Router /system-prompts [post]
func CreateSystemPromptDoc() {}

// UpdateSystemPromptDoc godoc
// @Summary 更新系统提示词
// @Tags SystemPrompt
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.UpdateSystemPromptRequest true "系统提示词"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /system-prompts [put]
func UpdateSystemPromptDoc() {}

// DeleteSystemPromptDoc godoc
// @Summary 删除系统提示词
// @Tags SystemPrompt
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DeleteSystemPromptRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /system-prompts [delete]
func DeleteSystemPromptDoc() {}

// AddDocumentDoc godoc
// @Summary 新增知识文档
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.AddKnowledgeDocumentRequest true "知识文档"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/document [post]
func AddDocumentDoc() {}

// UpdateDocumentDoc godoc
// @Summary 更新知识文档
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.UpdateKnowledgeDocumentRequest true "知识文档"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/document [put]
func UpdateDocumentDoc() {}

// DeleteDocumentDoc godoc
// @Summary 删除知识文档
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DeleteKnowledgeRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/document [delete]
func DeleteDocumentDoc() {}

// EnableDocumentDoc godoc
// @Summary 启用知识文档
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerIDRequest true "知识文档业务ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/document/enable [post]
func EnableDocumentDoc() {}

// DisableDocumentDoc godoc
// @Summary 禁用知识文档
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerIDRequest true "知识文档业务ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/document/disable [post]
func DisableDocumentDoc() {}

// ListDocumentsDoc godoc
// @Summary 获取知识文档列表
// @Tags Knowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param category query string true "分类"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.KnowledgeDocument]]
// @Router /knowledge/documents [get]
func ListDocumentsDoc() {}

// GetCategoriesDoc godoc
// @Summary 获取知识分类
// @Tags Knowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param type query string false "分类类型 text/image"
// @Success 200 {object} dto.Response[[]dto.KnowledgeCategory]
// @Router /knowledge/categories [get]
func GetCategoriesDoc() {}

// CreateKnowledgeCategoryDoc godoc
// @Summary 创建知识分类
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.CreateKnowledgeCategoryRequest true "知识分类"
// @Success 200 {object} dto.Response[dto.KnowledgeCategory]
// @Router /knowledge/category [post]
func CreateKnowledgeCategoryDoc() {}

// UpdateKnowledgeCategoryDoc godoc
// @Summary 更新知识分类
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.UpdateKnowledgeCategoryRequest true "知识分类"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/category [put]
func UpdateKnowledgeCategoryDoc() {}

// DeleteKnowledgeCategoryDoc godoc
// @Summary 删除知识分类
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DeleteKnowledgeCategoryRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/category [delete]
func DeleteKnowledgeCategoryDoc() {}

// SearchKnowledgeDoc godoc
// @Summary 搜索知识库
// @Tags Knowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SearchKnowledgeRequest true "搜索参数"
// @Success 200 {object} dto.Response[[]dto.VectorSearchResult]
// @Router /knowledge/search [post]
func SearchKnowledgeDoc() {}

// ReindexAllDoc godoc
// @Summary 重建文本知识库索引
// @Tags Knowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /knowledge/reindex [post]
func ReindexAllDoc() {}

// SaveMemoryDoc godoc
// @Summary 保存记忆
// @Tags Memory
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SaveMemoryRequest true "记忆"
// @Success 200 {object} dto.Response[dto.Memory]
// @Router /memory [post]
func SaveMemoryDoc() {}

// SearchMemoryDoc godoc
// @Summary 搜索记忆
// @Tags Memory
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SearchMemoryRequest true "搜索参数"
// @Success 200 {object} dto.Response[[]dto.Memory]
// @Router /memory/search [post]
func SearchMemoryDoc() {}

// DeleteMemoryDoc godoc
// @Summary 删除记忆
// @Tags Memory
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DeleteMemoryRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /memory [delete]
func DeleteMemoryDoc() {}

// AddImageDocumentDoc godoc
// @Summary 新增图片知识文档
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.AddImageKnowledgeRequest true "图片知识文档"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /image-knowledge/document [post]
func AddImageDocumentDoc() {}

// DeleteImageDocumentDoc godoc
// @Summary 删除图片知识文档
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.DeleteImageKnowledgeRequest true "删除参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /image-knowledge/document [delete]
func DeleteImageDocumentDoc() {}

// ListImageDocumentsDoc godoc
// @Summary 获取图片知识文档列表
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param category query string true "分类"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.ImageKnowledgeDocument]]
// @Router /image-knowledge/documents [get]
func ListImageDocumentsDoc() {}

// GetImageCategoriesDoc godoc
// @Summary 获取图片知识分类
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[[]dto.KnowledgeCategory]
// @Router /image-knowledge/categories [get]
func GetImageCategoriesDoc() {}

// SearchImageByTextDoc godoc
// @Summary 按文本搜索图片知识
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SearchImageKnowledgeByTextRequest true "搜索参数"
// @Success 200 {object} dto.Response[[]dto.VectorSearchResult]
// @Router /image-knowledge/search/text [post]
func SearchImageByTextDoc() {}

// SearchImageByImageDoc godoc
// @Summary 按图片搜索图片知识
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SearchImageKnowledgeByImageRequest true "搜索参数"
// @Success 200 {object} dto.Response[[]dto.VectorSearchResult]
// @Router /image-knowledge/search/image [post]
func SearchImageByImageDoc() {}

// ReindexAllImagesDoc godoc
// @Summary 重建图片知识库索引
// @Tags ImageKnowledge
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /image-knowledge/reindex [post]
func ReindexAllImagesDoc() {}

// ReindexAllVectorsDoc godoc
// @Summary 重建全部向量索引
// @Tags Vector
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /vector/reindex-all [post]
func ReindexAllVectorsDoc() {}

// FriendCircleGetListDoc godoc
// @Summary 获取朋友圈列表
// @Tags Moments
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param frist_page_md5 query string false "首页 MD5"
// @Param max_id query string true "最大朋友圈ID"
// @Success 200 {object} dto.Response[dto.MomentsGetListResponse]
// @Router /moments/list [get]
func FriendCircleGetListDoc() {}

// GetFriendCircleSettingsDoc godoc
// @Summary 获取朋友圈设置
// @Tags Moments
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.MomentSettings]
// @Router /moments/settings [get]
func GetFriendCircleSettingsDoc() {}

// FriendCircleGetDetailDoc godoc
// @Summary 获取朋友圈用户页详情
// @Tags Moments
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param Towxid query string true "目标微信ID"
// @Param Fristpagemd5 query string false "首页 MD5"
// @Param Maxid query string false "最大朋友圈ID"
// @Success 200 {object} dto.Response[dto.SnsUserPageResponse]
// @Router /moments/get-detail [get]
func FriendCircleGetDetailDoc() {}

// FriendCircleGetIdDetailDoc godoc
// @Summary 获取朋友圈单条详情
// @Tags Moments
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param Towxid query string true "目标微信ID"
// @Param MomentId query string true "朋友圈业务ID"
// @Success 200 {object} dto.Response[dto.SnsObjectDetailResponse]
// @Router /moments/get-id-detail [get]
func FriendCircleGetIdDetailDoc() {}

// FriendCircleDownFriendCircleMediaDoc godoc
// @Summary 下载朋友圈媒体
// @Tags Moments
// @Security ApiTokenAuth
// @Produce video/mp4
// @Param id query int true "机器人实例ID"
// @Param url query string true "媒体 URL"
// @Param key query string false "媒体 Key"
// @Success 200 {file} binary
// @Router /moments/down-media [get]
func FriendCircleDownFriendCircleMediaDoc() {}

// SaveFriendCircleSettingsDoc godoc
// @Summary 保存朋友圈设置
// @Tags Moments
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MomentSettings true "朋友圈设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /moments/settings [post]
func SaveFriendCircleSettingsDoc() {}

// FriendCircleCommentDoc godoc
// @Summary 评论朋友圈
// @Tags Moments
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.FriendCircleCommentRequest true "评论参数"
// @Success 200 {object} dto.Response[dto.SnsObject]
// @Router /moments/comment [post]
func FriendCircleCommentDoc() {}

// FriendCircleUploadDoc godoc
// @Summary 上传朋友圈媒体
// @Tags Moments
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param media formData file true "图片或视频文件"
// @Success 200 {object} dto.Response[dto.FriendCircleMedia]
// @Router /moments/upload-media [post]
func FriendCircleUploadDoc() {}

// FriendCirclePostDoc godoc
// @Summary 发布朋友圈
// @Tags Moments
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MomentPostRequest true "发布参数"
// @Success 200 {object} dto.Response[dto.MomentPostResponse]
// @Router /moments/post [post]
func FriendCirclePostDoc() {}

// FriendCircleOperationDoc godoc
// @Summary 操作朋友圈
// @Tags Moments
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MomentOpRequest true "操作参数"
// @Success 200 {object} dto.Response[dto.MomentOpResponse]
// @Router /moments/operate [post]
func FriendCircleOperationDoc() {}

// FriendCirclePrivacySettingsDoc godoc
// @Summary 设置朋友圈隐私
// @Tags Moments
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.MomentPrivacySettingsRequest true "隐私设置"
// @Success 200 {object} dto.Response[dto.MomentPrivacySettingsResponse]
// @Router /moments/privacy-settings [post]
func FriendCirclePrivacySettingsDoc() {}

// GetGlobalSettingsDoc godoc
// @Summary 获取全局设置
// @Tags GlobalSettings
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.GetGlobalSettingsResponse]
// @Router /global-settings [get]
func GetGlobalSettingsDoc() {}

// SaveGlobalSettingsDoc godoc
// @Summary 保存全局设置
// @Tags GlobalSettings
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SaveGlobalSettingsRequest true "全局设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /global-settings [post]
func SaveGlobalSettingsDoc() {}

// GetFriendSettingsDoc godoc
// @Summary 获取好友设置
// @Tags FriendSettings
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param contact_id query string true "联系人业务ID"
// @Success 200 {object} dto.Response[dto.GetFriendSettingsResponse]
// @Router /friend-settings [get]
func GetFriendSettingsDoc() {}

// SaveFriendSettingsDoc godoc
// @Summary 保存好友设置
// @Tags FriendSettings
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SaveFriendSettingsRequest true "好友设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /friend-settings [post]
func SaveFriendSettingsDoc() {}

// GetChatRoomSettingsDoc godoc
// @Summary 获取群设置
// @Tags ChatRoomSettings
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param chat_room_id query string true "群业务ID"
// @Success 200 {object} dto.Response[dto.GetChatRoomSettingsResponse]
// @Router /chat-room-settings [get]
func GetChatRoomSettingsDoc() {}

// SaveChatRoomSettingsDoc godoc
// @Summary 保存群设置
// @Tags ChatRoomSettings
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.SaveChatRoomSettingsRequest true "群设置"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /chat-room-settings [post]
func SaveChatRoomSettingsDoc() {}

// DoubaoTTSDoc godoc
// @Summary 豆包 TTS 回调
// @Tags AICallback
// @Accept json
// @Produce json
// @Param robot_id query int true "机器人实例ID"
// @Param request body dto.DoubaoTTSCallbackRequest true "回调参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /ai-callback/voice/doubao-tts [post]
func DoubaoTTSDoc() {}

// RobotListDoc godoc
// @Summary 获取机器人列表
// @Tags Robot
// @Security ApiTokenAuth
// @Produce json
// @Param owner query string false "拥有者"
// @Param status query string false "状态"
// @Param keyword query string false "关键词"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} dto.Response[dto.ListResponse[dto.RobotResponse]]
// @Router /robot/list [get]
func RobotListDoc() {}

// RobotCreateDoc godoc
// @Summary 创建机器人
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param request body dto.RobotCreateRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/create [post]
func RobotCreateDoc() {}

// RobotUpdateDoc godoc
// @Summary 更新机器人
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotUpdateRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/update [put]
func RobotUpdateDoc() {}

// RobotViewDoc godoc
// @Summary 获取机器人详情
// @Tags Robot
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.RobotResponse]
// @Router /robot/view [get]
func RobotViewDoc() {}

// ExportRobotLoginDataDoc godoc
// @Summary 导出机器人登录数据
// @Tags Robot
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[string]
// @Router /robot/export-login-data [get]
func ExportRobotLoginDataDoc() {}

// ImportRobotLoginDataDoc godoc
// @Summary 导入机器人登录数据
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerImportRobotLoginDataRequest true "登录数据"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/import-login-data [post]
func ImportRobotLoginDataDoc() {}

// RobotRestartClientDoc godoc
// @Summary 重启机器人客户端
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/restart-client [post]
func RobotRestartClientDoc() {}

// RobotRestartServerDoc godoc
// @Summary 重启机器人服务端
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/restart-server [post]
func RobotRestartServerDoc() {}

// RobotDockerImagePullDoc godoc
// @Summary 拉取机器人 Docker 镜像
// @Tags Robot
// @Security ApiTokenAuth
// @Produce text/event-stream
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.PullProgress
// @Router /robot/docker/image/pull [get]
func RobotDockerImagePullDoc() {}

// RobotStopAndRemoveClientDoc godoc
// @Summary 删除机器人客户端容器
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/docker/container/client/remove [delete]
func RobotStopAndRemoveClientDoc() {}

// RobotStopAndRemoveServerDoc godoc
// @Summary 删除机器人服务端容器
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/docker/container/server/remove [delete]
func RobotStopAndRemoveServerDoc() {}

// RobotStartClientDoc godoc
// @Summary 启动机器人客户端容器
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/docker/container/client/start [post]
func RobotStartClientDoc() {}

// RobotStartServerDoc godoc
// @Summary 启动机器人服务端容器
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotStartServerRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/docker/container/server/start [post]
func RobotStartServerDoc() {}

// RobotRemoveDoc godoc
// @Summary 删除机器人
// @Tags Robot
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotCommonRequest true "机器人"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/remove [delete]
func RobotRemoveDoc() {}

// RobotStateDoc godoc
// @Summary 检查机器人状态
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/state [get]
func RobotStateDoc() {}

// LoginSliderAutoVerifyDoc godoc
// @Summary 自动提交滑块验证
// @Tags RobotLogin
// @Accept json
// @Produce json
// @Param request body dto.SliderVerifyRequest true "滑块参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/login/slider-auto [post]
func LoginSliderAutoVerifyDoc() {}

// LoginSliderVerifyDoc godoc
// @Summary 获取滑块验证页面
// @Tags RobotLogin
// @Produce text/html
// @Param data62 query string true "data62"
// @Param ticket query string true "ticket"
// @Param device_type query string false "设备类型"
// @Param client_version query int false "客户端版本"
// @Success 200 {string} string
// @Router /robot/login/slider [get]
func LoginSliderVerifyDoc() {}

// LoginSliderVerifySubmitDoc godoc
// @Summary 提交滑块验证结果
// @Tags RobotLogin
// @Produce json
// @Param data query string true "滑块数据"
// @Param verifyid query string true "验证ID"
// @Param secticket query string true "安全票据"
// @Param randstr query string true "随机串"
// @Param ticket query string true "登录票据"
// @Param client_version query int false "客户端版本"
// @Param device_type query string false "设备类型"
// @Param ret query int false "返回码"
// @Param errorCode query int false "错误码"
// @Param errorMessage query string false "错误信息"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/login/slider-verify-submit [get]
func LoginSliderVerifySubmitDoc() {}

// RobotLoginDoc godoc
// @Summary 机器人登录
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body SwaggerRobotLoginRequest true "登录参数"
// @Success 200 {object} dto.Response[dto.RobotLoginResponse]
// @Router /robot/login [post]
func RobotLoginDoc() {}

// RobotLoginCheckDoc godoc
// @Summary 检查机器人登录状态
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotLoginCheckRequest true "登录检查参数"
// @Success 200 {object} dto.Response[dto.RobotLoginCheckResponse]
// @Router /robot/login-check [post]
func RobotLoginCheckDoc() {}

// RobotLogin2FADoc godoc
// @Summary 提交机器人二次验证
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.RobotLogin2FARequest true "二次验证参数"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/login/2fa [post]
func RobotLogin2FADoc() {}

// LoginData62LoginDoc godoc
// @Summary 使用 data62 登录
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.LoginRequest true "账号密码"
// @Success 200 {object} dto.Response[dto.UnifyAuthResponse]
// @Router /robot/login/data62 [post]
func LoginData62LoginDoc() {}

// LoginData62SMSAgainDoc godoc
// @Summary data62 登录重发短信
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.LoginData62SMSAgainRequest true "短信参数"
// @Success 200 {object} dto.Response[dto.UnifyAuthResponse]
// @Router /robot/login/data62-sms-again [post]
func LoginData62SMSAgainDoc() {}

// LoginData62SMSVerifyDoc godoc
// @Summary data62 登录短信验证
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.LoginData62SMSVerifyRequest true "短信验证码参数"
// @Success 200 {object} dto.Response[dto.UnifyAuthResponse]
// @Router /robot/login/data62-sms-verify [post]
func LoginData62SMSVerifyDoc() {}

// LoginA16Data1Doc godoc
// @Summary 使用 A16 登录
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Accept json
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param request body dto.LoginRequest true "账号密码"
// @Success 200 {object} dto.Response[dto.UnifyAuthResponse]
// @Router /robot/login/a16 [post]
func LoginA16Data1Doc() {}

// RobotLogoutDoc godoc
// @Summary 机器人登出
// @Tags RobotLogin
// @Security ApiTokenAuth
// @Produce json
// @Param id query int true "机器人实例ID"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /robot/logout [delete]
func RobotLogoutDoc() {}

// WxappQrcodeAuthLoginDoc godoc
// @Summary 小程序二维码授权登录
// @Tags WXApp
// @Security ApiTokenAuth
// @Accept multipart/form-data
// @Produce json
// @Param id query int true "机器人实例ID"
// @Param qrcode formData file true "二维码图片"
// @Success 200 {object} dto.Response[dto.SwaggerEmpty]
// @Router /wxapp/qrcode-auth-login [post]
func WxappQrcodeAuthLoginDoc() {}

// ProxyPprofGetDoc godoc
// @Summary 代理机器人 pprof GET
// @Tags Pprof
// @Security ApiTokenAuth
// @Produce plain
// @Param id query int true "机器人实例ID"
// @Param pprof_path path string true "pprof 路径"
// @Success 200 {string} string
// @Router /pprof/{pprof_path} [get]
func ProxyPprofGetDoc() {}

// ProxyPprofPostDoc godoc
// @Summary 代理机器人 pprof POST
// @Tags Pprof
// @Security ApiTokenAuth
// @Produce plain
// @Param id query int true "机器人实例ID"
// @Param pprof_path path string true "pprof 路径"
// @Success 200 {string} string
// @Router /pprof/{pprof_path} [post]
func ProxyPprofPostDoc() {}
