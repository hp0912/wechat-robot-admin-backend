package dto

type SystemSettingsRequest struct {
	ID                         int64   `json:"id"`
	SystemSettingsID           int64   `json:"system_settings_id"`
	APITokenEnabled            *bool   `json:"api_token_enabled"`
	OfflineNotificationEnabled *bool   `json:"offline_notification_enabled"`
	NotificationType           string  `json:"notification_type"`
	PushPlusURL                *string `json:"push_plus_url"`
	PushPlusToken              *string `json:"push_plus_token"`
	AutoVerifyUser             *bool   `json:"auto_verify_user"`
	VerifyUserDelay            int     `json:"verify_user_delay"`
	AutoChatroomInvite         *bool   `json:"auto_chatroom_invite"`
}

type SystemSettings struct {
	ID                         int64   `json:"id"`
	APITokenEnabled            *bool   `json:"api_token_enabled"`
	APIToken                   *string `json:"api_token"`
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
