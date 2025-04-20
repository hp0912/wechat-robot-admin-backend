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
	ID   int64  `form:"id" json:"id" binding:"required"`
	Uuid string `form:"uuid" json:"uuid"`
}

type RobotLoginResponse struct {
	Awken bool   `json:"awken"` // 微信登录凭证
	Uuid  string `json:"uuid"`
}

type AcctSectRespData struct {
	Username   string `json:"userName"`   // 原始微信Id
	Alias      string `json:"alias"`      // 自定义的微信号
	BindMobile string `json:"bindMobile"` // 绑定的手机号
	FsUrl      string `json:"fsurl"`      // 可能是头像地址
	Nickname   string `json:"nickName"`   // 昵称
}

type RobotLoginCheckResponse struct {
	Uuid                    string           `json:"uuid"`
	Status                  int              `json:"status"` // 状态
	PushLoginUrlexpiredTime int              `json:"pushLoginUrlexpiredTime"`
	ExpiredTime             int              `json:"expiredTime"`  // 过期时间(秒)
	HeadImgUrl              string           `json:"headImgUrl"`   // 头像
	NickName                string           `json:"nickName"`     // 昵称
	AcctSectResp            AcctSectRespData `json:"acctSectResp"` // 账号信息-登录成功之后才有
}
