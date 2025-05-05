package dto

type ChatHistoryRequest struct {
	ID        int64  `form:"id" json:"id"  binding:"required"`
	ContactID string `form:"contact_id" json:"contact_id" binding:"required"`
	Keyword   string `form:"keyword" json:"keyword"`
}

// MessageType 以Go惯用形式定义了PC微信所有的官方消息类型。
type MessageType int

// AppMessageType 以Go惯用形式定义了PC微信所有的官方App消息类型。
type AppMessageType int

const (
	MsgTypeText           MessageType = 1     // 文本消息
	MsgTypeImage          MessageType = 3     // 图片消息
	MsgTypeVoice          MessageType = 34    // 语音消息
	MsgTypeVerify         MessageType = 37    // 认证消息
	MsgTypePossibleFriend MessageType = 40    // 好友推荐消息
	MsgTypeShareCard      MessageType = 42    // 名片消息
	MsgTypeVideo          MessageType = 43    // 视频消息
	MsgTypeEmoticon       MessageType = 47    // 表情消息
	MsgTypeLocation       MessageType = 48    // 地理位置消息
	MsgTypeApp            MessageType = 49    // APP消息
	MsgTypeVoip           MessageType = 50    // VOIP消息
	MsgTypeInit           MessageType = 51    // 微信初始化消息
	MsgTypeVoipNotify     MessageType = 52    // VOIP结束消息
	MsgTypeVoipInvite     MessageType = 53    // VOIP邀请
	MsgTypeMicroVideo     MessageType = 62    // 小视频消息
	MsgTypeUnknow         MessageType = 9999  // 未知消息
	MsgTypeSys            MessageType = 10000 // 系统消息
	MsgTypeRecalled       MessageType = 10002 // 消息撤回
)

const (
	AppMsgTypeText                  AppMessageType = 1      // 文本消息
	AppMsgTypeImg                   AppMessageType = 2      // 图片消息
	AppMsgTypeAudio                 AppMessageType = 3      // 语音消息
	AppMsgTypeVideo                 AppMessageType = 4      // 视频消息
	AppMsgTypeUrl                   AppMessageType = 5      // 文章消息
	AppMsgTypeAttach                AppMessageType = 6      // 附件消息
	AppMsgTypeOpen                  AppMessageType = 7      // Open
	AppMsgTypeEmoji                 AppMessageType = 8      // 表情消息
	AppMsgTypeVoiceRemind           AppMessageType = 9      // VoiceRemind
	AppMsgTypeScanGood              AppMessageType = 10     // ScanGood
	AppMsgTypeGood                  AppMessageType = 13     // Good
	AppMsgTypeEmotion               AppMessageType = 15     // Emotion
	AppMsgTypeCardTicket            AppMessageType = 16     // 名片消息
	AppMsgTypeRealtimeShareLocation AppMessageType = 17     // 地理位置消息
	AppMsgTypequote                 AppMessageType = 57     // 引用消息
	AppMsgTypeAttachUploading       AppMessageType = 74     // 附件上传中
	AppMsgTypeTransfers             AppMessageType = 2000   // 转账消息
	AppMsgTypeRedEnvelopes          AppMessageType = 2001   // 红包消息
	AppMsgTypeReaderType            AppMessageType = 100001 //自定义的消息
)

type ChatHistory struct {
	ID                 int64          `json:"id"`
	MsgId              int64          `json:"msg_id"`        // 消息Id
	ClientMsgId        int64          `json:"client_msg_id"` // 客户端消息Id
	IsGroup            bool           `json:"is_group"`
	IsAtMe             bool           `json:"is_atme"` // @所有人 好的
	IsRecalled         bool           `json:"is_recalled"`
	Type               MessageType    `json:"type"`                 // 消息类型
	AppMsgType         AppMessageType `json:"app_msg_type"`         // 消息类型
	Content            string         `json:"content"`              // 内容
	DisplayFullContent string         `json:"display_full_content"` // 显示的完整内容
	MessageSource      string         `json:"message_source"`
	FromWxID           string         `json:"from_wxid"`      // 消息来源
	SenderWxID         string         `json:"sender_wxid"`    // 消息发送者
	ToWxID             string         `json:"to_wxid"`        // 接收者
	AttachmentUrl      string         `json:"attachment_url"` // 文件地址
	CreatedAt          int64          `json:"created_at"`
	UpdatedAt          int64          `json:"updated_at"`
	SenderNickname     string         `json:"sender_nickname"`
	SenderAvatar       string         `json:"sender_avatar"`
}

type AttachDownloadRequest struct {
	ID        int64  `form:"id" json:"id"  binding:"required"`
	MessageID int64  `form:"message_id" json:"message_id" binding:"required"`
	AttachUrl string `form:"attach_url" json:"attach_url"`
}
