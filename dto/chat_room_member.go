package dto

type ChatRoomRequestBase struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id" binding:"required"`
}

type ChatRoomMemberRequest struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id" binding:"required"`
	Keyword    string `form:"keyword" json:"keyword"`
}

type ChatRoomOperateRequest struct {
	ChatRoomRequestBase
	Content string `form:"content" json:"content" binding:"required"`
}

type CreateChatRoomRequest struct {
	ID         int64    `form:"id" json:"id"  binding:"required"`
	ContactIDs []string `form:"contact_ids" json:"contact_ids" binding:"required"`
}

type InviteChatRoomMemberRequest struct {
	ID         int64    `form:"id" json:"id"  binding:"required"`
	ChatRoomID string   `form:"chat_room_id" json:"chat_room_id" binding:"required"`
	ContactIDs []string `form:"contact_ids" json:"contact_ids" binding:"required"`
}

type ChatRoomJoinRequest struct {
	ID              int64 `form:"id" json:"id"  binding:"required"`
	SystemMessageID int64 `form:"system_message_id" json:"system_message_id"`
}

type DelChatRoomMemberRequest struct {
	ChatRoomRequestBase
	MemberIDs []string `form:"member_ids" json:"member_ids" binding:"required"`
}

type ChatRoomMember struct {
	ID                   int64  `json:"id"`                     // 主键ID
	ChatRoomID           string `json:"chat_room_id"`           // 群ID
	WechatID             string `json:"wechat_id"`              // 微信ID
	Alias                string `json:"alias"`                  // 微信号
	Nickname             string `json:"nickname"`               // 昵称
	Avatar               string `json:"avatar"`                 // 头像
	InviterWechatID      string `json:"inviter_wechat_id"`      // 邀请人微信ID
	IsAdmin              bool   `json:"is_admin"`               // 是否群管理员
	IsBlacklisted        bool   `json:"is_blacklisted"`         // 是否在黑名单
	IsLeaved             bool   `json:"is_leaved"`              // 是否已经离开群聊
	Score                *int64 `json:"score"`                  // 积分
	TemporaryScore       *int64 `json:"temporary_score"`        // 临时积分
	TemporaryScoreExpiry *int64 `json:"temporary_score_expiry"` // 临时积分有效期
	FrozenScore          *int64 `json:"frozen_score"`           // 冻结积分
	FrozenTemporaryScore *int64 `json:"frozen_temporary_score"` // 冻结临时积分
	Remark               string `json:"remark"`                 // 备注
	JoinedAt             int64  `json:"joined_at"`              // 加入时间
	LastActiveAt         int64  `json:"last_active_at"`         // 最近活跃时间
	LeavedAt             *int64 `json:"leaved_at"`              // 离开时间
}
