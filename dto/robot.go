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

type RobotLoginCheckRequest struct {
	Uuid string `form:"uuid" json:"uuid" binding:"required"`
}

type RobotLoginResponse struct {
	AutoLogin  bool   `json:"auto_login"`  // 自动登陆
	AwkenLogin bool   `json:"awken_login"` // 微信登录凭证
	Uuid       string `json:"uuid"`
	Data62     string `json:"data62"`
}

type AcctSectResp struct {
	Username   string `json:"userName"`   // 原始微信Id
	Alias      string `json:"alias"`      // 自定义的微信号
	BindMobile string `json:"bindMobile"` // 绑定的手机号
	FsUrl      string `json:"fsurl"`      // 可能是头像地址
	Nickname   string `json:"nickName"`   // 昵称
}

type RobotLoginCheckResponse struct {
	Uuid                    string       `json:"uuid"`
	Status                  int          `json:"status"` // 状态
	PushLoginUrlexpiredTime int          `json:"pushLoginUrlexpiredTime"`
	ExpiredTime             int          `json:"expiredTime"`  // 过期时间(秒)
	HeadImgUrl              string       `json:"headImgUrl"`   // 头像
	NickName                string       `json:"nickName"`     // 昵称
	Ticket                  string       `json:"ticket"`       // 登录票据
	AcctSectResp            AcctSectResp `json:"acctSectResp"` // 账号信息-登录成功之后才有
}

type RobotLogin2FARequest struct {
	ID     int64  `form:"id" json:"id" binding:"required"`
	Uuid   string `form:"uuid" json:"uuid" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required"`
	Ticket string `form:"ticket" json:"ticket" binding:"required"`
	Data62 string `form:"data62" json:"data62" binding:"required"`
}

type SliderVerifyRequest struct {
	Data62 string `form:"data62" json:"data62" binding:"required"`
	Ticket string `form:"ticket" json:"ticket" binding:"required"`
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SilderOCR struct {
	Flag    int    `json:"flag"`
	Data    string `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	URL     string `json:"url"`
	Remark  string `json:"remark"`
	Success bool   `json:"Success"`
}

type UnifyAuthResponse struct {
	CheckUrl          string        `json:"checkUrl,omitempty"`
	AgainUrl          string        `json:"AgainUrl,omitempty"`
	Cookie            string        `json:"Cookie,omitempty"`
	UnifyAuthSectFlag *uint32       `json:"unifyAuthSectFlag,omitempty"`
	AuthSectResp      *AuthSectResp `json:"authSectResp,omitempty"`
	// AcctSectResp      *AcctSectResp      `json:"acctSectResp,omitempty"`
	// NetworkSectResp   *NetworkSectResp   `json:"networkSectResp,omitempty"`
	// AxAuthSecRespList *AxAuthSecRespList `json:"axAuthSecRespList,omitempty"`
}

type AuthSectResp struct {
	Uin                  *uint32               `json:"uin,omitempty"`
	SvrPubEcdhkey        *ECDHKey              `json:"svrPubEcdhkey,omitempty"`
	SessionKey           *SKBuiltinBufferT     `json:"sessionKey,omitempty"`
	AutoAuthKey          *SKBuiltinBufferT     `json:"autoAuthKey,omitempty"`
	WtloginRspBuffFlag   *uint32               `json:"wtloginRspBuffFlag,omitempty"`
	WtloginRspBuff       *SKBuiltinBufferT     `json:"wtloginRspBuff,omitempty"`
	WtloginImgRespInfo   *WTLoginImgRespInfo   `json:"wtloginImgRespInfo,omitempty"`
	WxVerifyCodeRespInfo *WxVerifyCodeRespInfo `json:"wxVerifyCodeRespInfo,omitempty"`
	CliDbencryptKey      *SKBuiltinBufferT     `json:"cliDbencryptKey,omitempty"`
	CliDbencryptInfo     *SKBuiltinBufferT     `json:"cliDbencryptInfo,omitempty"`
	AuthKey              *string               `json:"authKey,omitempty"`
	A2Key                *SKBuiltinBufferT     `json:"a2Key,omitempty"`
	ApplyBetaUrl         *string               `json:"applyBetaUrl,omitempty"`
	ShowStyle            *ShowStyleKey         `json:"showStyle,omitempty"`
	AuthTicket           *string               `json:"authTicket,omitempty"`
	NewVersion           *uint32               `json:"newVersion,omitempty"`
	UpdateFlag           *uint32               `json:"updateFlag,omitempty"`
	AuthResultFlag       *uint32               `json:"authResultFlag,omitempty"`
	Fsurl                *string               `json:"fsurl,omitempty"`
	MmtlsControlBitFlag  *uint32               `json:"mmtlsControlBitFlag,omitempty"`
	ServerTime           *uint32               `json:"serverTime,omitempty"`
	ClientSessionKey     *SKBuiltinBufferT     `json:"clientSessionKey,omitempty"`
	ServerSessionKey     *SKBuiltinBufferT     `json:"serverSessionKey,omitempty"`
	EcdhControlFlag      *uint32               `json:"ecdhControlFlag,omitempty"`
}

type ECDHKey struct {
	Nid *int32            `json:"nid,omitempty"`
	Key *SKBuiltinBufferT `json:"key,omitempty"`
}

type WTLoginImgRespInfo struct {
	ImgEncryptKey *string           `json:"imgEncryptKey,omitempty"`
	Ksid          *SKBuiltinBufferT `json:"ksid,omitempty"`
	ImgSid        *string           `json:"imgSid,omitempty"`
	ImgBuf        *SKBuiltinBufferT `json:"imgBuf,omitempty"`
}

type ShowStyleKey struct {
	KeyCount *uint32 `json:"keyCount,omitempty"`
	// Key      []string `json:"key,omitempty"`
}

type WxVerifyCodeRespInfo struct {
	VerifySignature *string           `json:"verifySignature,omitempty"`
	VerifyBuff      *SKBuiltinBufferT `json:"verifyBuff,omitempty"`
}

type LoginData62SMSAgainRequest struct {
	Url    string
	Cookie string
}

type LoginData62SMSVerifyRequest struct {
	Url    string
	Cookie string
	Sms    string
}
