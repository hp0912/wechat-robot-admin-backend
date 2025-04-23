package vars

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

var DockerNetwork string
var DockerComposeCmd string

var WeChatServerAddress string
var WeChatServerToken string

var SessionSecret string

var OpenAIApiKey string
