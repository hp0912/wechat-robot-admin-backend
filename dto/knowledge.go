package dto

type ListResponse[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

type AddKnowledgeDocumentRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Source   string `json:"source"`
	Category string `json:"category"`
}

type UpdateKnowledgeDocumentRequest struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Source  string `json:"source"`
}

type SearchKnowledgeRequest struct {
	Query    string `json:"query" binding:"required"`
	Category string `json:"category"`
	Limit    int    `json:"limit"`
}

type ListKnowledgeRequest struct {
	Category string `form:"category" json:"category"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
}

type DeleteKnowledgeRequest struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type SaveMemoryRequest struct {
	ContactWxID string `json:"contact_wxid" binding:"required"`
	ChatRoomID  string `json:"chat_room_id"`
	Type        string `json:"type" binding:"required"`
	Key         string `json:"key" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Importance  int    `json:"importance"`
}

type SearchMemoryRequest struct {
	ContactWxID string `form:"contact_wxid" json:"contact_wxid"`
	Query       string `form:"query" json:"query"`
	Limit       int    `form:"limit" json:"limit"`
}

type DeleteMemoryRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type AddImageKnowledgeRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url" binding:"required"`
	Category    string `json:"category"`
}

type DeleteImageKnowledgeRequest struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type ListImageKnowledgeRequest struct {
	Category string `form:"category" json:"category"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
}

type SearchImageKnowledgeByTextRequest struct {
	Query    string `json:"query" binding:"required"`
	Category string `json:"category"`
	Limit    int    `json:"limit"`
}

type SearchImageKnowledgeByImageRequest struct {
	ImageURL string `json:"image_url" binding:"required"`
	Category string `json:"category"`
	Limit    int    `json:"limit"`
}

type CreateKnowledgeCategoryRequest struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateKnowledgeCategoryRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type DeleteKnowledgeCategoryRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type KnowledgeDocument struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Source     string `json:"source"`
	Category   string `json:"category"`
	ChunkIndex int    `json:"chunk_index"`
	ChunkTotal int    `json:"chunk_total"`
	VectorID   string `json:"vector_id"`
	Enabled    bool   `json:"enabled"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type Memory struct {
	ID           int64  `json:"id"`
	ContactWxID  string `json:"contact_wxid"`
	ChatRoomID   string `json:"chat_room_id"`
	Type         string `json:"type"`
	Key          string `json:"key"`
	Content      string `json:"content"`
	Source       string `json:"source"`
	Importance   int    `json:"importance"`
	AccessCount  int    `json:"access_count"`
	LastAccessAt int64  `json:"last_access_at"`
	ExpireAt     int64  `json:"expire_at"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type ImageKnowledgeDocument struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Category    string `json:"category"`
	VectorID    string `json:"vector_id"`
	Enabled     bool   `json:"enabled"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type KnowledgeCategory struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsBuiltin   bool   `json:"is_builtin"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type VectorSearchResult struct {
	ID      string            `json:"ID"`
	Score   float32           `json:"Score"`
	Payload map[string]string `json:"Payload"`
}
