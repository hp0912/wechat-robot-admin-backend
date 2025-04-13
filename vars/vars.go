package vars

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

var WeChatServerAddress string
var WeChatServerToken string

var SessionSecret string
