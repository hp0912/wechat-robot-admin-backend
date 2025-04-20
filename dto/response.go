package dto

import "fmt"

type Response[T any] struct {
	Code    int    `json:"code"`    // 状态码
	Data    T      `json:"data"`    // 数据
	Message string `json:"message"` // 消息
}

func (r Response[T]) IsSuccess() bool {
	return r.Code == 200
}

func (r Response[T]) CheckError(err error) error {
	if err != nil {
		return err
	}
	if r.Code != 200 {
		return fmt.Errorf("error code: %d, message: %s", r.Code, r.Message)
	}
	return nil
}
