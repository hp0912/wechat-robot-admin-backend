package dto

type RobotListRequest struct {
	Owner   string `form:"owner" json:"owner"`
	Status  string `form:"status" json:"status"`
	Keyword string `form:"keyword" json:"keyword"`
}

type RobotCreateRequest struct {
	RobotCode string `form:"robot_code" json:"robot_code" binding:"required"`
}
