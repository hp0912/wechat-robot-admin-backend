package dto

type Response[T any] struct {
	Code    int    `json:"code"`    // 状态码
	Data    T      `json:"data"`    // 数据
	Message string `json:"message"` // 消息
}
