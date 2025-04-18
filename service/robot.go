package service

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"wechat-robot-client/dto"
	"wechat-robot-client/model"
	"wechat-robot-client/pkg/appx"
	"wechat-robot-client/repository"
	"wechat-robot-client/utils"
	"wechat-robot-client/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DockerComposeFileContext struct {
	WECHAT_PORT         string
	REDIS_HOST          string
	REDIS_PORT          string
	REDIS_PASSWORD      string
	REDIS_DB            string
	GIN_MODE            string
	ROBOT_CODE          string
	ROBOT_START_TIMEOUT string
	MYSQL_DRIVER        string
	MYSQL_HOST          string
	MYSQL_PORT          string
	MYSQL_USER          string
	MYSQL_PASSWORD      string
	MYSQL_ADMIN_DB      string
	MYSQL_DB            string
	MYSQL_SCHEMA        string
	DOCKER_NETWORK      string
}

type RobotService struct {
	ctx context.Context
}

func NewRobotService(ctx context.Context) *RobotService {
	return &RobotService{
		ctx: ctx,
	}
}

func (r *RobotService) RobotList(ctx *gin.Context, req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	return repository.NewRobotRepo(r.ctx, vars.DB).RobotList(ctx, req, pager)
}

func (r *RobotService) RobotCreate(ctx *gin.Context, req dto.RobotCreateRequest) error {
	session := sessions.Default(ctx)
	wechatId := session.Get("wechat_id")
	respo := repository.NewRobotRepo(r.ctx, vars.DB)
	redisDb, err := respo.GetMaxRedisDB()
	if err != nil {
		return err
	}
	robot := &model.Robot{
		RobotCode:    req.RobotCode,
		Owner:        wechatId.(string),
		DeviceId:     utils.CreateDeviceID(""),
		DeviceName:   utils.CreateDeviceName(),
		WeChatID:     "", // 登陆后才会有
		Nickname:     "",
		Avatar:       vars.RobotDefaultAvatar,
		Status:       model.RobotStatusOffline,
		ErrorMessage: "",
		LastLoginAt:  0,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		RedisDB:      redisDb + 1,
	}
	respo.Create(robot)
	// 创建机器人实例数据库
	err = vars.DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", robot.RobotCode)).Error
	if err != nil {
		return err
	}
	// 创建机器人实例表
	newDsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		vars.MysqlSettings.User, vars.MysqlSettings.Password, vars.MysqlSettings.Host, vars.MysqlSettings.Port, robot.RobotCode)
	mysqlConfig := mysql.Config{
		DSN: newDsn,
	}
	// gorm 配置
	gormConfig := gorm.Config{}
	newDB, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		return err
	}
	db, err := newDB.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	// 读取建表模版
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("无法获取运行时信息")
	}
	projectRoot := filepath.Join(filepath.Dir(filename), "..") // 上一级为项目根目录
	sqlFilePath := filepath.Join(projectRoot, "admin.sql")     // TODO admin.sql
	// 检查文件是否存在
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		return errors.New("建表模版不存在")
	}
	// 读取文件内容
	content, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}
	// 开始建表
	err = newDB.Exec(fmt.Sprintf("USE `%s`;\n%s", robot.RobotCode, string(content))).Error
	if err != nil {
		return err
	}
	// 生成docker-compose.yml
	// 读取模板文件
	templateFile := filepath.Join(projectRoot, "docker-compose", "docker-compose.temp")
	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	// 解析模板
	tmpl, err := template.New("docker-compose").Parse(string(tmplContent))
	if err != nil {
		return err
	}
	data := DockerComposeFileContext{
		WECHAT_PORT:         "9000",
		REDIS_HOST:          vars.RedisSettings.Host,
		REDIS_PORT:          vars.RedisSettings.Port,
		REDIS_PASSWORD:      vars.RedisSettings.Password,
		REDIS_DB:            fmt.Sprintf("%d", robot.RedisDB),
		GIN_MODE:            "release",
		ROBOT_CODE:          robot.RobotCode,
		ROBOT_START_TIMEOUT: "60",
		MYSQL_DRIVER:        vars.MysqlSettings.Driver,
		MYSQL_HOST:          vars.MysqlSettings.Host,
		MYSQL_PORT:          vars.MysqlSettings.Port,
		MYSQL_USER:          vars.MysqlSettings.User,
		MYSQL_PASSWORD:      vars.MysqlSettings.Password,
		MYSQL_ADMIN_DB:      vars.MysqlSettings.Db,
		MYSQL_DB:            robot.RobotCode,
		MYSQL_SCHEMA:        vars.MysqlSettings.Schema,
		DOCKER_NETWORK:      vars.DockerNetwork,
	}
	// 创建目标文件
	outputFilePath := filepath.Join(projectRoot, "docker-compose", fmt.Sprintf("docker-compose-%s.yml", robot.RobotCode))
	err = r.CreateDockerComposeFile(tmpl, outputFilePath, data)
	if err != nil {
		return err
	}
	// 通过 docker-compose 启动微信客户端和服务端
	cmdParts := strings.Fields(vars.DockerComposeCmd)
	cmdArgs := append(cmdParts[1:], "-f", outputFilePath, "up", "-d")
	cmd := exec.Command(cmdParts[0], cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// RobotView 查看机器人元数据
func (r *RobotService) RobotView(ctx *gin.Context, robotID int64) *model.Robot {
	respo := repository.NewRobotRepo(r.ctx, vars.DB)
	return respo.GetByID(robotID)
}

// RobotRemove 删除机器人
func (r *RobotService) RobotRemove(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(r.ctx, vars.DB)
	robot := respo.GetByID(robotID)
	if robot == nil {
		return errors.New("机器人不存在")
	}
	// 删除机器人实例数据
	respo.DeleteById(robotID)
	// 删除机器人数据库
	err := vars.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`;", robot.RobotCode)).Error
	if err != nil {
		return err
	}
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("无法获取运行时信息")
	}
	projectRoot := filepath.Join(filepath.Dir(filename), "..") // 上一级为项目根目录
	dockerComposeFile := filepath.Join(projectRoot, "docker-compose", fmt.Sprintf("docker-compose-%s.yml", robot.RobotCode))
	// 判断docker-compose文件是否存在
	_, err = os.Stat(dockerComposeFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	// 通过docker-compose停止微信客户端和服务端
	cmdParts := strings.Fields(vars.DockerComposeCmd)
	cmdArgs := append(cmdParts[1:], "-f", dockerComposeFile, "down")
	cmd := exec.Command(cmdParts[0], cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	// 删除docker-compose文件
	err = os.Remove(dockerComposeFile)
	if err != nil {
		return nil
	}
	return nil
}

func (r *RobotService) CreateDockerComposeFile(tmpl *template.Template, outputFilePath string, data DockerComposeFileContext) error {
	outputFile, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	// 渲染模板并写入目标文件
	if err := tmpl.Execute(outputFile, data); err != nil {
		return err
	}
	return nil
}
