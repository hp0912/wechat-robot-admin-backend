package dto

// SystemPrompt 系统提示词
type SystemPrompt struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// CreateSystemPromptRequest 创建系统提示词请求
type CreateSystemPromptRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateSystemPromptRequest 更新系统提示词请求
type UpdateSystemPromptRequest struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// DeleteSystemPromptRequest 删除系统提示词请求
type DeleteSystemPromptRequest struct {
	ID int64 `json:"id" binding:"required"`
}

// GetSystemPromptRequest 获取系统提示词请求
type GetSystemPromptRequest struct {
	ID       int64 `form:"id" json:"id" binding:"required"`
	PromptID int64 `form:"prompt_id" json:"prompt_id" binding:"required"`
}

// ListSystemPromptRequest 获取系统提示词列表请求
type ListSystemPromptRequest struct {
	Keyword string `form:"keyword" json:"keyword"`
}
