package dto

// 进度信息结构体
type PullProgress struct {
	Image    string `json:"image"`
	Status   string `json:"status"`
	Progress string `json:"progress,omitempty"`
	Error    string `json:"error,omitempty"`
}
