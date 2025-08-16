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

type SendFileMessageRequest struct {
	ID              int64  `form:"id" json:"id" binding:"required"`
	ToWxid          string `form:"to_wxid" json:"to_wxid" binding:"required"`
	ClientAppDataId string `form:"client_app_data_id" json:"client_app_data_id" binding:"required"`
	Filename        string `form:"filename" json:"filename" binding:"required"`
	FileHash        string `form:"file_hash" json:"file_hash" binding:"required"`
	FileSize        int64  `form:"file_size" json:"file_size" binding:"required"`
	ChunkIndex      int64  `form:"chunk_index" json:"chunk_index"`
	TotalChunks     int64  `form:"total_chunks" json:"total_chunks" binding:"required"`
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
