package vars

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

var DockerNetwork string
var DockerComposeCmd string

var LoginMethod string
var LoginToken string
var WeChatServerAddress string // 登录管理员后台的微信鉴权服务器地址 公众号扫码登录
var WeChatServerToken string
var WeChatOfficialAccountAuthURL string // 微信公众号授权地址

var SliderServerBaseURL string
var SliderToken string

var SessionSecret string

var OpenAIApiKey string

var ThirdPartyApiKey string

var UploadFileChunkSize int64 = 50000
