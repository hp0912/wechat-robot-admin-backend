package dto

type GetGlobalSettingsRequest struct {
	ID int64 `form:"id" json:"id"  binding:"required"`
}

type GetGlobalSettingsResponse struct {
	ID                        int64          `form:"id" json:"id"`
	ChatAIEnabled             bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatAITrigger             string         `form:"chat_ai_trigger" json:"chat_ai_trigger"`
	ChatBaseURL               string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey                string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel                 string         `form:"chat_model" json:"chat_model"`
	ChatPrompt                string         `form:"chat_prompt" json:"chat_prompt"`
	ImageAIEnabled            bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageModel                string         `form:"image_model" json:"image_model"`
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
	LeaveChatRoomAlertEnabled bool           `form:"leave_chat_room_alert_enabled" json:"leave_chat_room_alert_enabled"`
	LeaveChatRoomAlertText    string         `form:"leave_chat_room_alert_text" json:"leave_chat_room_alert_text"`
	ChatRoomRankingEnabled    bool           `form:"chat_room_ranking_enabled" json:"chat_room_ranking_enabled"`
	ChatRoomRankingDailyCron  string         `form:"chat_room_ranking_daily_cron" json:"chat_room_ranking_daily_cron"`
	ChatRoomRankingWeeklyCron string         `form:"chat_room_ranking_weekly_cron" json:"chat_room_ranking_weekly_cron"`
	ChatRoomRankingMonthCron  string         `form:"chat_room_ranking_month_cron" json:"chat_room_ranking_month_cron"`
	ChatRoomSummaryEnabled    bool           `form:"chat_room_summary_enabled" json:"chat_room_summary_enabled"`
	ChatRoomSummaryModel      string         `form:"chat_room_summary_model" json:"chat_room_summary_model"`
	ChatRoomSummaryCron       string         `form:"chat_room_summary_cron" json:"chat_room_summary_cron"`
	NewsEnabled               bool           `form:"news_enabled" json:"news_enabled"`
	NewsType                  string         `form:"news_type" json:"news_type"`
	NewsCron                  string         `form:"news_cron" json:"news_cron"`
	MorningEnabled            bool           `form:"morning_enabled" json:"morning_enabled"`
	MorningCron               string         `form:"morning_cron" json:"morning_cron"`
	FriendSyncCron            string         `form:"friend_sync_cron" json:"friend_sync_cron"`
}

type SaveGlobalSettingsRequest struct {
	ID                        int64          `form:"id" json:"id"  binding:"required"`
	ConfigID                  int64          `form:"config_id" json:"config_id"`
	ChatAIEnabled             bool           `form:"chat_ai_enabled" json:"chat_ai_enabled"`
	ChatAITrigger             string         `form:"chat_ai_trigger" json:"chat_ai_trigger"`
	ChatBaseURL               string         `form:"chat_base_url" json:"chat_base_url"`
	ChatAPIKey                string         `form:"chat_api_key" json:"chat_api_key"`
	ChatModel                 string         `form:"chat_model" json:"chat_model"`
	ChatPrompt                string         `form:"chat_prompt" json:"chat_prompt"`
	ImageAIEnabled            bool           `form:"image_ai_enabled" json:"image_ai_enabled"`
	ImageModel                string         `form:"image_model" json:"image_model"`
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
	LeaveChatRoomAlertEnabled bool           `form:"leave_chat_room_alert_enabled" json:"leave_chat_room_alert_enabled"`
	LeaveChatRoomAlertText    string         `form:"leave_chat_room_alert_text" json:"leave_chat_room_alert_text"`
	ChatRoomRankingEnabled    bool           `form:"chat_room_ranking_enabled" json:"chat_room_ranking_enabled"`
	ChatRoomRankingDailyCron  string         `form:"chat_room_ranking_daily_cron" json:"chat_room_ranking_daily_cron"`
	ChatRoomRankingWeeklyCron string         `form:"chat_room_ranking_weekly_cron" json:"chat_room_ranking_weekly_cron"`
	ChatRoomRankingMonthCron  string         `form:"chat_room_ranking_month_cron" json:"chat_room_ranking_month_cron"`
	ChatRoomSummaryEnabled    bool           `form:"chat_room_summary_enabled" json:"chat_room_summary_enabled"`
	ChatRoomSummaryModel      string         `form:"chat_room_summary_model" json:"chat_room_summary_model"`
	ChatRoomSummaryCron       string         `form:"chat_room_summary_cron" json:"chat_room_summary_cron"`
	NewsEnabled               bool           `form:"news_enabled" json:"news_enabled"`
	NewsType                  string         `form:"news_type" json:"news_type"`
	NewsCron                  string         `form:"news_cron" json:"news_cron"`
	MorningEnabled            bool           `form:"morning_enabled" json:"morning_enabled"`
	MorningCron               string         `form:"morning_cron" json:"morning_cron"`
	FriendSyncCron            string         `form:"friend_sync_cron" json:"friend_sync_cron"`
}
