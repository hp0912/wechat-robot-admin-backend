package model

type CommonConfig struct {
	ID                  int64  `gorm:"column:id;primaryKey;autoIncrement;comment:'公共配置表主键ID'" json:"id"`
	AIEnabled           bool   `gorm:"column:ai_enabled;default:false;comment:'是否启用AI聊天功能'" json:"ai_enabled"`
	RankingEnabled      bool   `gorm:"column:ranking_enabled;default:false;comment:'是否启用群聊排行榜功能'" json:"ranking_enabled"`
	SummaryEnabled      bool   `gorm:"column:summary_enabled;default:false;comment:'是否启用聊天记录总结功能'" json:"summary_enabled"`
	WelcomeEnabled      bool   `gorm:"column:welcome_enabled;default:false;comment:'是否启用新成员加群欢迎功能'" json:"welcome_enabled"`
	ChatURL             string `gorm:"column:chat_url;default:'';comment:'聊天AI的基础URL地址'" json:"chat_url"`
	ChatKey             string `gorm:"column:chat_key;default:'';comment:'聊天AI的API密钥'" json:"chat_key"`
	ChatModel           string `gorm:"column:chat_model;default:'';comment:'聊天AI使用的模型名称'" json:"chat_model"`
	ChatPersona         string `gorm:"column:chat_persona;type:text;default:'';comment:'聊天AI的人设配置'" json:"chat_persona"`
	SummaryModel        string `gorm:"column:summary_model;default:'';comment:'聊天总结使用的AI模型名称'" json:"summary_model"`
	ImageURL            string `gorm:"column:image_url;default:'';comment:'绘图AI的基础URL地址'" json:"image_url"`
	ImageKey            string `gorm:"column:image_key;default:'';comment:'绘图AI的API密钥'" json:"image_key"`
	ImageSecret         string `gorm:"column:image_secret;default:'';comment:'绘图AI的API密钥secret'" json:"image_secret"`
	ImageModel          string `gorm:"column:image_model;default:'';comment:'绘图AI使用的模型名称'" json:"image_model"`
	ImageScheduler      string `gorm:"column:image_scheduler;default:'';comment:'绘图AI的请求调度配置'" json:"image_scheduler"`
	NewsEnabled         bool   `gorm:"column:news_enabled;default:false;comment:'是否启用每日早报功能'" json:"news_enabled"`
	NewsCron            string `gorm:"column:news_cron;default:'';comment:'每日早报的定时任务表达式'" json:"news_cron"`
	FriendSyncEnabled   bool   `gorm:"column:friend_sync_enabled;default:false;comment:'是否启用好友同步功能'" json:"friend_sync_enabled"`
	FriendSyncCron      string `gorm:"column:friend_sync_cron;default:'';comment:'好友同步的定时任务表达式'" json:"friend_sync_cron"`
	GroupSummaryEnabled bool   `gorm:"column:group_summary_enabled;default:false;comment:'是否启用群聊总结功能'" json:"group_summary_enabled"`
	GroupSummaryCron    string `gorm:"column:group_summary_cron;default:'';comment:'群聊总结的定时任务表达式'" json:"group_summary_cron"`
	MorningEnabled      bool   `gorm:"column:morning_enabled;default:false;comment:'是否启用早安问候功能'" json:"morning_enabled"`
	MorningCron         string `gorm:"column:morning_cron;default:'';comment:'早安问候的定时任务表达式'" json:"morning_cron"`
	RankingCron         string `gorm:"column:ranking_cron;default:'';comment:'群聊排行榜的定时任务表达式'" json:"ranking_cron"`
}

// TableName 指定表名
func (CommonConfig) TableName() string {
	return "common-configs"
}
