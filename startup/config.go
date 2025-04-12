package startup

import (
	"log"
	"os"
	"strings"
	"wechat-robot-client/vars"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	loadEnvConfig()
	return nil
}

func loadEnvConfig() {
	// 本地开发模式
	isDevMode := strings.ToLower(os.Getenv("GO_ENV")) == "dev"
	if isDevMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("加载本地环境变量失败，请检查是否存在 .env 文件")
		}
	}

	vars.SessionSecret = os.Getenv("SESSION_SECRET")
	if vars.SessionSecret == "" {
		log.Fatal("SESSION_SECRET 环境变量未设置")
	}

	vars.WechatServerAddress = os.Getenv("WECHAT_SERVER_ADDRESS")
	vars.WechatServerToken = os.Getenv("WECHAT_SERVER_TOKEN")

	// mysql
	vars.MysqlSettings.Driver = os.Getenv("MYSQL_DRIVER")
	vars.MysqlSettings.Host = os.Getenv("MYSQL_HOST")
	vars.MysqlSettings.Port = os.Getenv("MYSQL_PORT")
	vars.MysqlSettings.User = os.Getenv("MYSQL_USER")
	vars.MysqlSettings.Password = os.Getenv("MYSQL_PASSWORD")
	vars.MysqlSettings.Db = os.Getenv("MYSQL_DB")
	vars.MysqlSettings.Schema = os.Getenv("MYSQL_SCHEMA")
}
