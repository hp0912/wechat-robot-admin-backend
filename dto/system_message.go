package dto

type SystemMessage struct {
	ID          int64  `json:"id"`
	MsgID       int64  `json:"msg_id"`
	ClientMsgID int64  `json:"client_msg_id"`
	Type        int    `json:"type"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Content     string `json:"content"`
	FromWxid    string `json:"from_wxid"`
	ToWxid      string `json:"to_wxid"`
	Status      int    `json:"status"`
	IsRead      bool   `json:"is_read"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
