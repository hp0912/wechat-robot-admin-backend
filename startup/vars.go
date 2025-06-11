package startup

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"wechat-robot-admin-backend/template"
	"wechat-robot-admin-backend/utils/docker"
	"wechat-robot-admin-backend/vars"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupVars() error {
	if err := InitMySQLClient(); err != nil {
		return err
	}
	log.Println("MySQL连接成功")
	if err := InitMysqlTables(); err != nil {
		return err
	}
	log.Println("MySQL表初始化成功")
	if err := InitDockerNetwork(); err != nil {
		return err
	}
	log.Println("Docker网络连接成功")
	return nil
}

func InitMySQLClient() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		vars.MysqlSettings.User, vars.MysqlSettings.Password, vars.MysqlSettings.Host, vars.MysqlSettings.Port, vars.MysqlSettings.Db)
	mysqlConfig := mysql.Config{
		DSN:                     dsn,
		DontSupportRenameIndex:  true, // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn: true, // 用 `change` 重命名列
	}
	// gorm 配置
	gormConfig := gorm.Config{}
	// 是否开启调试模式
	if flag, _ := strconv.ParseBool(os.Getenv("GORM_DEBUG")); flag {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	vars.DB, err = gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	return err
}

func InitMysqlTables() error {
	// 配置不太一样，必须包含 multiStatements=true 以支持多条 SQL 语句执行
	// 因此新建一个数据库连接
	temporaryDsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		vars.MysqlSettings.User, vars.MysqlSettings.Password, vars.MysqlSettings.Host, vars.MysqlSettings.Port, vars.MysqlSettings.Db)
	mysqlConfig := mysql.Config{
		DSN: temporaryDsn,
	}
	// gorm 配置
	gormConfig := gorm.Config{}
	temporaryDB, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		return err
	}
	db, err := temporaryDB.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = temporaryDB.Exec(fmt.Sprintf("USE `%s`;\n%s", vars.MysqlSettings.Db, template.AdminSqlTemplate)).Error
	if err != nil {
		return err
	}
	return nil
}

func InitDockerNetwork() error {
	networkName := vars.DockerNetwork
	if networkName == "" {
		return fmt.Errorf("docker network name is not set")
	}
	if !docker.NetworkExists(networkName) {
		err := docker.CreateNetwork(networkName)
		if err != nil {
			return fmt.Errorf("failed to create Docker network: %v", err)
		}
	}
	return nil
}
