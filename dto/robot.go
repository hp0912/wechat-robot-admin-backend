package dto

type RobotListRequest struct {
	Owner   string `form:"owner" json:"owner"`
	Status  string `form:"status" json:"status"`
	Keyword string `form:"keyword" json:"keyword"`
}

type RobotCreateRequest struct {
	RobotCode string `form:"robot_code" json:"robot_code" binding:"required"`
}

type RobotCommonRequest struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type RobotLoginResponse struct {
	Qrcode string `json:"qrcode"`
	Uuid   string `json:"uuid"`
}

type RobotLoginCheckResponse struct {
	LoggedIn string `json:"logged_in"`
}
