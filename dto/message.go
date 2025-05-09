package dto

type MessageRevokeRequest struct {
	ID        int64 `form:"id" json:"id"  binding:"required"`
	MessageID int64 `form:"message_id" json:"message_id" binding:"required"`
}

type RobotMessageRevokeRequest struct {
	MessageID int64 `form:"message_id" json:"message_id" binding:"required"`
}
