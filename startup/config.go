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

	// mysql
	vars.MysqlSettings.Driver = os.Getenv("MYSQL_DRIVER")
	vars.MysqlSettings.Host = os.Getenv("MYSQL_HOST")
	vars.MysqlSettings.Port = os.Getenv("MYSQL_PORT")
	vars.MysqlSettings.User = os.Getenv("MYSQL_USER")
	vars.MysqlSettings.Password = os.Getenv("MYSQL_PASSWORD")
	// 机器人ID就是数据库名
	vars.MysqlSettings.Db = os.Getenv("ROBOT_ID")
	vars.MysqlSettings.AdminDb = os.Getenv("MYSQL_ADMIN_DB")
	vars.MysqlSettings.Schema = os.Getenv("MYSQL_SCHEMA")
}
