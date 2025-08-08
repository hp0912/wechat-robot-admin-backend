package dto

type SystemSettings struct {
	ID                         int64   `json:"id"`
	OfflineNotificationEnabled *bool   `json:"offline_notification_enabled"`
	NotificationType           string  `json:"notification_type"`
	PushPlusURL                *string `json:"push_plus_url"`
	PushPlusToken              *string `json:"push_plus_token"`
	AutoVerifyUser             *bool   `json:"auto_verify_user"`
	VerifyUserDelay            int     `json:"verify_user_delay"`
	AutoChatroomInvite         *bool   `json:"auto_chatroom_invite"`
	CreatedAt                  int64   `json:"created_at"`
	UpdatedAt                  int64   `json:"updated_at"`
}
