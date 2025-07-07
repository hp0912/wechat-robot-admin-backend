package dto

type WxappQrcodeAuthLoginRequest struct {
	ID  int64  `form:"id" json:"id"  binding:"required"`
	URL string `form:"url" json:"url"  binding:"required"`
}
