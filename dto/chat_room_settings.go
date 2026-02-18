package dto

type GetChatRoomSettingsRequest struct {
	ID         int64  `form:"id" json:"id"  binding:"required"`
	ChatRoomID string `form:"chat_room_id" json:"chat_room_id"  binding:"required"`
}

type GetChatRoomSettingsResponse struct {
	ID                        int64          `form:"id" json:"id"`
	ChatAIEnabled             bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatAITrigger             string         `form:"chat_ai_trigger" json:"chat_ai_trigger"`
	ChatBaseURL               string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey                string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel                 string         `form:"chat_model" json:"chat_model"`
	ImageRecognitionModel     string         `form:"image_recognition_model" json:"image_recognition_model"`
	ChatPrompt                string         `form:"chat_prompt" json:"chat_prompt"`
	MaxCompletionTokens       int            `form:"max_completion_tokens" json:"max_completion_tokens"`
	ImageAIEnabled            bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageAISettings           map[string]any `form:"image_ai_settings" json:"image_ai_settings"`
	TTSEnabled                bool           `form:"tts_enabled" json:"tts_enabled"`
	TTSSettings               map[string]any `form:"tts_settings" json:"tts_settings"`
	LTTSSettings              map[string]any `form:"ltts_settings" json:"ltts_settings"`
	WelcomeEnabled            bool           `form:"welcome_enabled" json:"welcome_enabled"`
	WelcomeType               string         `form:"welcome_type" json:"welcome_type"`
	WelcomeText               string         `form:"welcome_text" json:"welcome_text"`
	WelcomeEmojiMD5           string         `form:"welcome_emoji_md5" json:"welcome_emoji_md5"`
	WelcomeEmojiLen           int64          `form:"welcome_emoji_len" json:"welcome_emoji_len"`
	WelcomeImageURL           string         `form:"welcome_image_url" json:"welcome_image_url"`
	WelcomeURL                string         `form:"welcome_url" json:"welcome_url"`
	ShortVideoParsingEnabled  bool           `form:"short_video_parsing_enabled" json:"short_video_parsing_enabled"`
	WxhbNotifyEnabled         bool           `form:"wxhb_notify_enabled" json:"wxhb_notify_enabled"`
	WxhbNotifyMemberList      string         `form:"wxhb_notify_member_list" json:"wxhb_notify_member_list"`
	PodcastEnabled            bool           `form:"podcast_enabled" json:"podcast_enabled"`
	PodcastConfig             map[string]any `form:"podcast_config" json:"podcast_config"`
	PatEnabled                bool           `form:"pat_enabled" json:"pat_enabled"`
	PatType                   string         `form:"pat_type" json:"pat_type"`
	PatText                   string         `form:"pat_text" json:"pat_text"`
	PatVoiceTimbre            string         `form:"pat_voice_timbre" json:"pat_voice_timbre"`
	LeaveChatRoomAlertEnabled bool           `form:"leave_chat_room_alert_enabled" json:"leave_chat_room_alert_enabled"`
	LeaveChatRoomAlertText    string         `form:"leave_chat_room_alert_text" json:"leave_chat_room_alert_text"`
	ChatRoomRankingEnabled    bool           `form:"chat_room_ranking_enabled" json:"chat_room_ranking_enabled"`
	ChatRoomSummaryEnabled    bool           `form:"chat_room_summary_enabled" json:"chat_room_summary_enabled"`
	ChatRoomSummaryModel      string         `form:"chat_room_summary_model" json:"chat_room_summary_model"`
	NewsEnabled               bool           `form:"news_enabled" json:"news_enabled"`
	NewsType                  string         `form:"news_type" json:"news_type"`
	MorningEnabled            bool           `form:"morning_enabled" json:"morning_enabled"`
}

type SaveChatRoomSettingsRequest struct {
	ID                        int64          `form:"id" json:"id"  binding:"required"`
	ConfigID                  int64          `form:"config_id" json:"config_id"`
	ChatRoomID                string         `form:"chat_room_id" json:"chat_room_id"`
	ChatAIEnabled             bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatAITrigger             string         `form:"chat_ai_trigger" json:"chat_ai_trigger"`
	ChatBaseURL               string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey                string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel                 string         `form:"chat_model" json:"chat_model"`
	ImageRecognitionModel     string         `form:"image_recognition_model" json:"image_recognition_model"`
	ChatPrompt                string         `form:"chat_prompt" json:"chat_prompt"`
	MaxCompletionTokens       int            `form:"max_completion_tokens" json:"max_completion_tokens"`
	ImageAIEnabled            bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageAISettings           map[string]any `form:"image_ai_settings" json:"image_ai_settings"`
	TTSEnabled                bool           `form:"tts_enabled" json:"tts_enabled"`
	TTSSettings               map[string]any `form:"tts_settings" json:"tts_settings"`
	LTTSSettings              map[string]any `form:"ltts_settings" json:"ltts_settings"`
	WelcomeEnabled            bool           `form:"welcome_enabled" json:"welcome_enabled"`
	WelcomeType               string         `form:"welcome_type" json:"welcome_type"`
	WelcomeText               string         `form:"welcome_text" json:"welcome_text"`
	WelcomeEmojiMD5           string         `form:"welcome_emoji_md5" json:"welcome_emoji_md5"`
	WelcomeEmojiLen           int64          `form:"welcome_emoji_len" json:"welcome_emoji_len"`
	WelcomeImageURL           string         `form:"welcome_image_url" json:"welcome_image_url"`
	WelcomeURL                string         `form:"welcome_url" json:"welcome_url"`
	ShortVideoParsingEnabled  bool           `form:"short_video_parsing_enabled" json:"short_video_parsing_enabled"`
	WxhbNotifyEnabled         bool           `form:"wxhb_notify_enabled" json:"wxhb_notify_enabled"`
	WxhbNotifyMemberList      string         `form:"wxhb_notify_member_list" json:"wxhb_notify_member_list"`
	PodcastEnabled            bool           `form:"podcast_enabled" json:"podcast_enabled"`
	PodcastConfig             map[string]any `form:"podcast_config" json:"podcast_config"`
	PatEnabled                bool           `form:"pat_enabled" json:"pat_enabled"`
	PatType                   string         `form:"pat_type" json:"pat_type"`
	PatText                   string         `form:"pat_text" json:"pat_text"`
	PatVoiceTimbre            string         `form:"pat_voice_timbre" json:"pat_voice_timbre"`
	LeaveChatRoomAlertEnabled bool           `form:"leave_chat_room_alert_enabled" json:"leave_chat_room_alert_enabled"`
	LeaveChatRoomAlertText    string         `form:"leave_chat_room_alert_text" json:"leave_chat_room_alert_text"`
	ChatRoomRankingEnabled    bool           `form:"chat_room_ranking_enabled" json:"chat_room_ranking_enabled"`
	ChatRoomSummaryEnabled    bool           `form:"chat_room_summary_enabled" json:"chat_room_summary_enabled"`
	ChatRoomSummaryModel      string         `form:"chat_room_summary_model" json:"chat_room_summary_model"`
	NewsEnabled               bool           `form:"news_enabled" json:"news_enabled"`
	NewsType                  string         `form:"news_type" json:"news_type"`
	MorningEnabled            bool           `form:"morning_enabled" json:"morning_enabled"`
}
