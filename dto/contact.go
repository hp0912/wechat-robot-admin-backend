package dto

type ContactType string

const (
	ContactTypeFriend   ContactType = "friend"
	ContactTypeChatRoom ContactType = "chat_room"
)

type SyncContactsRequest struct {
	ID int64 `form:"id" json:"id"`
}

type GetContactsRequest struct {
	ID         int64    `form:"id" json:"id" binding:"required"`
	ContactIDs []string `form:"contact_ids" json:"contact_ids"`
	Type       string   `form:"type" json:"type"`
	Keyword    string   `form:"keyword" json:"keyword"`
}

type GetContactsResponse struct {
	ID            int64       `json:"id"`
	WechatID      string      `json:"wechat_id"` // 添加索引长度
	Alias         string      `json:"alias"`     // 微信号
	Nickname      string      `json:"nickname"`
	Avatar        string      `json:"avatar"`
	Type          ContactType `json:"type"`
	Remark        string      `json:"remark"`
	Pyinitial     string      `json:"pyinitial"`       // 昵称拼音首字母大写
	QuanPin       string      `json:"quan_pin"`        // 昵称拼音全拼小写
	Sex           int         `json:"sex"`             // 性别 0：未知 1：男 2：女
	Country       string      `json:"country"`         // 国家
	Province      string      `json:"province"`        // 省份
	City          string      `json:"city"`            // 城市
	Signature     string      `json:"signature"`       // 个性签名
	SnsBackground string      `json:"sns_background"`  // 朋友圈背景图
	ChatRoomOwner string      `json:"chat_room_owner"` // 群主微信号
	CreatedAt     int64       `json:"created_at"`
	LastActiveAt  int64       `json:"last_active_at"` // 最近活跃时间
	UpdatedAt     int64       `json:"updated_at"`
}

type FriendSearchRequest struct {
	ID          int64  `form:"id" json:"id" binding:"required"`
	ToUserName  string `form:"to_username" json:"to_username" binding:"required"`
	FromScene   int    `form:"from_scene" json:"from_scene"`
	SearchScene int    `form:"search_scene" json:"search_scene"`
}

type FriendSendRequestRequest struct {
	ID            int64  `form:"id" json:"id" binding:"required"`
	V1            string `form:"v1" json:"V1" binding:"required"`
	V2            string `form:"v2" json:"V2" binding:"required"`
	Opcode        int    `form:"opcode" json:"Opcode"`
	Scene         int    `form:"scene" json:"Scene"`
	VerifyContent string `form:"verify_content" json:"VerifyContent"`
}

type FriendSetRemarksRequest struct {
	ID      int64  `form:"id" json:"id" binding:"required"`
	ToWxid  string `form:"to_wxid" json:"to_wxid" binding:"required"`
	Remarks string `form:"remarks" json:"remarks" binding:"required"`
}

type FriendPassVerifyRequest struct {
	ID              int64 `form:"id" json:"id" binding:"required"`
	SystemMessageID int64 `form:"system_message_id" json:"system_message_id"`
}

type FriendDeleteRequest struct {
	ID        int64  `form:"id" json:"id" binding:"required"`
	ContactID string `form:"contact_id" json:"contact_id"`
}
