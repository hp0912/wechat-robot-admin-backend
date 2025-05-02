package dto

type ContactType string

const (
	ContactTypeFriend ContactType = "friend"
	ContactTypeGroup  ContactType = "group"
)

type GetContactsRequest struct {
	ID      int64  `form:"id" json:"id"`
	Type    string `form:"type" json:"type"`
	Keyword string `form:"keyword" json:"keyword"`
}

type GetContactsResponse struct {
	ID            int64       `json:"id"`
	WechatID      string      `json:"wechat_id"` // 添加索引长度
	Owner         string      `json:"owner"`     // 联系人所有者
	Alias         string      `json:"alias"`     // 微信号
	Nickname      string      `json:"nickname"`
	Avatar        string      `json:"avatar"`
	Type          ContactType `json:"type"`
	Remark        string      `json:"remark"`
	Pyinitial     string      `json:"pyinitial"`      // 昵称拼音首字母大写
	QuanPin       string      `json:"quan_pin"`       // 昵称拼音全拼小写
	Sex           int         `json:"sex"`            // 性别 0：未知 1：男 2：女
	Country       string      `json:"country"`        // 国家
	Province      string      `json:"province"`       // 省份
	City          string      `json:"city"`           // 城市
	Signature     string      `json:"signature"`      // 个性签名
	SnsBackground string      `json:"sns_background"` // 朋友圈背景图
	CreatedAt     int64       `json:"created_at"`
	UpdatedAt     int64       `json:"updated_at"`
}
