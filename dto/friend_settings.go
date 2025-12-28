package dto

type GetFriendSettingsRequest struct {
	ID        int64  `form:"id" json:"id"  binding:"required"`
	ContactID string `form:"contact_id" json:"contact_id"  binding:"required"`
}

type GetFriendSettingsResponse struct {
	ID                    int64          `form:"id" json:"id"`
	ChatAIEnabled         bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatBaseURL           string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey            string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel             string         `form:"chat_model" json:"chat_model"`
	ImageRecognitionModel string         `form:"image_recognition_model" json:"image_recognition_model"`
	ChatPrompt            string         `form:"chat_prompt" json:"chat_prompt"`
	MaxCompletionTokens   int            `form:"max_completion_tokens" json:"max_completion_tokens"`
	ImageAIEnabled        bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageModel            string         `form:"image_model" json:"image_model"`
	ImageAISettings       map[string]any `form:"image_ai_settings" json:"image_ai_settings"`
	TTSEnabled            bool           `form:"tts_enabled" json:"tts_enabled"`
	TTSSettings           map[string]any `form:"tts_settings" json:"tts_settings"`
	LTTSSettings          map[string]any `form:"ltts_settings" json:"ltts_settings"`
}

type SaveFriendSettingsRequest struct {
	ID                    int64          `form:"id" json:"id"  binding:"required"`
	ConfigID              int64          `form:"config_id" json:"config_id"`
	WeChatID              string         `form:"wechat_id" json:"wechat_id"  binding:"required"`
	ChatAIEnabled         bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatBaseURL           string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey            string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel             string         `form:"chat_model" json:"chat_model"`
	ImageRecognitionModel string         `form:"image_recognition_model" json:"image_recognition_model"`
	ChatPrompt            string         `form:"chat_prompt" json:"chat_prompt"`
	MaxCompletionTokens   int            `form:"max_completion_tokens" json:"max_completion_tokens"`
	ImageAIEnabled        bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageModel            string         `form:"image_model" json:"image_model"`
	ImageAISettings       map[string]any `form:"image_ai_settings" json:"image_ai_settings"`
	TTSEnabled            bool           `form:"tts_enabled" json:"tts_enabled"`
	TTSSettings           map[string]any `form:"tts_settings" json:"tts_settings"`
	LTTSSettings          map[string]any `form:"ltts_settings" json:"ltts_settings"`
}
