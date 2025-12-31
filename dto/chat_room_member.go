package dto

type ChatRoomRequestBase struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id" binding:"required"`
}

type ChatRoomMemberListRequest struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id" binding:"required"`
	Keyword    string `form:"keyword" json:"keyword"`
}

type ChatRoomMemberRequest struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id" binding:"required"`
	WechatID   string `form:"wechat_id" json:"wechat_id" binding:"required"`
}

type UpdateChatRoomMemberRequest struct {
	ChatRoomID           string  `form:"chat_room_id" json:"chat_room_id" binding:"required"`
	WechatID             string  `form:"wechat_id" json:"wechat_id" binding:"required"`
	Batch                bool    `form:"batch" json:"batch"`
	IsAdmin              *bool   `form:"is_admin" json:"is_admin"`
	IsBlacklisted        *bool   `form:"is_blacklisted" json:"is_blacklisted"`
	TemporaryScoreAction *string `form:"temporary_score_action" json:"temporary_score_action"`
	TemporaryScore       *int64  `form:"temporary_score" json:"temporary_score"`
	ScoreAction          *string `form:"score_action" json:"score_action"`
	Score                *int64  `form:"score" json:"score"`
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
	Remark               string `json:"remark"`                 // 备注
	JoinedAt             int64  `json:"joined_at"`              // 加入时间
	LastActiveAt         int64  `json:"last_active_at"`         // 最近活跃时间
	LeavedAt             *int64 `json:"leaved_at"`              // 离开时间
}
