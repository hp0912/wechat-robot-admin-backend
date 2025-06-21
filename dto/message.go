package dto

type MessageRevokeRequest struct {
	ID        int64 `form:"id" json:"id"  binding:"required"`
	MessageID int64 `form:"message_id" json:"message_id" binding:"required"`
}

type RobotMessageRevokeRequest struct {
	MessageID int64 `form:"message_id" json:"message_id" binding:"required"`
}

type RobotSendTextMessageRequest struct {
	ToWxid  string   `form:"to_wxid" json:"to_wxid" binding:"required"`
	Content string   `form:"content" json:"content" binding:"required"`
	At      []string `form:"at" json:"at"`
}

type SendTextMessageRequest struct {
	ID int64 `form:"id" json:"id"  binding:"required"`
	RobotSendTextMessageRequest
}

type SendImageMessageRequest struct {
	ID     int64  `form:"id" json:"id"  binding:"required"`
	ToWxid string `form:"to_wxid" json:"to_wxid" binding:"required"`
}

type SendVideoMessageRequest struct {
	ID     int64  `form:"id" json:"id"  binding:"required"`
	ToWxid string `form:"to_wxid" json:"to_wxid" binding:"required"`
}

type SendVoiceMessageRequest struct {
	ID     int64  `form:"id" json:"id"  binding:"required"`
	ToWxid string `form:"to_wxid" json:"to_wxid" binding:"required"`
}

type RobotSendAITTSMessageRequest struct {
	ID      int64  `form:"id" json:"id"  binding:"required"`
	ToWxid  string `form:"to_wxid" json:"to_wxid" binding:"required"`
	Speaker string `form:"speaker" json:"speaker" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type TimbreResponse struct {
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	Count    int      `json:"count"`
	Speakers []string `json:"speakers"`
}

type RobotSendAITTSMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Speaker   string `json:"speaker"`
	Text      string `json:"text"`
	Audiopath string `json:"audiopath"`
}
