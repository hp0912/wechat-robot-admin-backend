package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/template"
	"wechat-robot-admin-backend/utils"
	"wechat-robot-admin-backend/vars"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	dockerImage "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RobotManageService struct {
	ctx context.Context
}

func NewRobotManageService(ctx context.Context) *RobotManageService {
	return &RobotManageService{
		ctx: ctx,
	}
}

func (sv *RobotManageService) RobotList(req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	return repository.NewRobotRepo(sv.ctx, vars.DB).RobotList(req, pager)
}

// 辅助方法：获取Docker客户端
func (sv *RobotManageService) getDockerClient() (*client.Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	return dockerClient, nil
}

func (sv *RobotManageService) DockerStartClientAndServer(ctx *gin.Context, robot *model.Robot) error {
	// 创建Docker客户端
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 服务端容器配置
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	serverContainerName := fmt.Sprintf("server_%s", robot.RobotCode)
	serverConfig := &container.Config{
		Image: "registry.cn-shenzhen.aliyuncs.com/houhou/wechat-ipad:latest",
		Env: []string{
			fmt.Sprintf("WECHAT_PORT=%s", "9000"),
			fmt.Sprintf("REDIS_HOST=%s", vars.RedisSettings.Host),
			fmt.Sprintf("REDIS_PORT=%s", vars.RedisSettings.Port),
			fmt.Sprintf("REDIS_PASSWORD=%s", vars.RedisSettings.Password),
			fmt.Sprintf("REDIS_DB=%d", robot.RedisDB),
			fmt.Sprintf("WECHAT_CLIENT_HOST=%s", fmt.Sprintf("%s:%d", clientContainerName, 9000)),
		},
	}

	serverHostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}

	// 服务端网络配置
	serverNetworkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			vars.DockerNetwork: {},
		},
	}

	// 创建服务端容器
	serverResp, err := dockerClient.ContainerCreate(
		sv.ctx,
		serverConfig,
		serverHostConfig,
		serverNetworkConfig,
		nil,
		serverContainerName,
	)
	if err != nil {
		return fmt.Errorf("创建服务端容器失败: %v", err)
	}

	// 启动服务端容器
	err = dockerClient.ContainerStart(sv.ctx, serverResp.ID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("启动服务端容器失败: %v", err)
	}

	// 客户端容器配置
	clientConfig := &container.Config{
		Image: "registry.cn-shenzhen.aliyuncs.com/houhou/wechat-robot-client:latest",
		Env: []string{
			fmt.Sprintf("GIN_MODE=%s", "release"),
			fmt.Sprintf("ROBOT_ID=%d", robot.ID),
			fmt.Sprintf("ROBOT_CODE=%s", robot.RobotCode),
			fmt.Sprintf("ROBOT_START_TIMEOUT=%s", "60"),
			fmt.Sprintf("MYSQL_DRIVER=%s", vars.MysqlSettings.Driver),
			fmt.Sprintf("MYSQL_HOST=%s", vars.MysqlSettings.Host),
			fmt.Sprintf("MYSQL_PORT=%s", vars.MysqlSettings.Port),
			fmt.Sprintf("MYSQL_USER=%s", vars.MysqlSettings.User),
			fmt.Sprintf("MYSQL_PASSWORD=%s", vars.MysqlSettings.Password),
			fmt.Sprintf("MYSQL_ADMIN_DB=%s", vars.MysqlSettings.Db),
			fmt.Sprintf("MYSQL_DB=%s", robot.RobotCode),
			fmt.Sprintf("MYSQL_SCHEMA=%s", vars.MysqlSettings.Schema),
		},
	}

	clientHostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}

	// 客户端网络配置
	clientNetworkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			vars.DockerNetwork: {},
		},
	}

	// 创建客户端容器
	clientResp, err := dockerClient.ContainerCreate(
		sv.ctx,
		clientConfig,
		clientHostConfig,
		clientNetworkConfig,
		nil,
		clientContainerName,
	)
	if err != nil {
		return fmt.Errorf("创建客户端容器失败: %v", err)
	}

	// 启动客户端容器
	err = dockerClient.ContainerStart(sv.ctx, clientResp.ID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("启动客户端容器失败: %v", err)
	}

	return nil
}

func (sv *RobotManageService) DockerStopAndRemoveClientAndServer(ctx *gin.Context, robot *model.Robot) error {
	// 使用Docker SDK停止并删除容器
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 停止并删除服务端容器
	serverContainerName := fmt.Sprintf("server_%s", robot.RobotCode)
	err = sv.stopAndRemoveContainer(dockerClient, serverContainerName)
	if err != nil {
		return fmt.Errorf("删除服务端容器失败: %v", err)
	}

	// 停止并删除客户端容器
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	err = sv.stopAndRemoveContainer(dockerClient, clientContainerName)
	if err != nil {
		return fmt.Errorf("删除客户端容器失败: %v", err)
	}

	return nil
}

func (sv *RobotManageService) RobotCreate(ctx *gin.Context, req dto.RobotCreateRequest) error {
	session := sessions.Default(ctx)
	wechatId := session.Get("wechat_id")
	role := session.Get("role")
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	redisDb, err := respo.GetMaxRedisDB()
	if err != nil {
		return err
	}
	// 一个账号最多创建2个机器人
	robots := respo.GetByOwner(wechatId.(string), true)
	if len(robots) >= 2 && role.(int) != vars.RoleRootUser {
		return errors.New("一个账号最多创建2个机器人")
	}

	robot := &model.Robot{
		RobotCode:    req.RobotCode,
		Owner:        wechatId.(string),
		DeviceID:     utils.CreateDeviceID(""),
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
	// 开始建表
	err = newDB.Exec(fmt.Sprintf("USE `%s`;\n%s", robot.RobotCode, template.RobotSqlTemplate)).Error
	if err != nil {
		return err
	}
	// 插入一条公共配置记录
	commonConf := fmt.Sprintf("INSERT INTO `%s`.`%s` (`owner`, `chat_ai_enabled`, `chat_base_url`, `chat_api_key`, `chat_model`, `chat_prompt`, `friend_sync_cron`) VALUES ('%s', 1, '%s', '%s', '%s', '%s', '0 * * * *');",
		robot.RobotCode, "global_settings", robot.WeChatID, "https://ai-api.houhoukang.com/", vars.OpenAIApiKey, "gpt-4o-mini", "我是一个聊天机器人。")
	err = newDB.Exec(commonConf).Error
	if err != nil {
		return err
	}

	err = sv.DockerStartClientAndServer(ctx, robot)
	if err != nil {
		return err
	}

	return nil
}

// RobotView 查看机器人元数据
func (sv *RobotManageService) RobotView(robotID int64) *model.Robot {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	return respo.GetByID(robotID)
}

// RobotStopAndRemoveClientAndServer 删除机器人容器
func (sv *RobotManageService) RobotStopAndRemoveClientAndServer(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot := respo.GetByID(robotID)
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err := sv.DockerStopAndRemoveClientAndServer(ctx, robot)
	if err != nil {
		return err
	}
	return nil
}

// RobotStartClientAndServer 启动机器人容器
func (sv *RobotManageService) RobotStartClientAndServer(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot := respo.GetByID(robotID)
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err := sv.DockerStartClientAndServer(ctx, robot)
	if err != nil {
		return err
	}
	return nil
}

// RobotRemove 删除机器人
func (sv *RobotManageService) RobotRemove(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
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

	err = sv.DockerStopAndRemoveClientAndServer(ctx, robot)
	if err != nil {
		return err
	}

	return nil
}

func (sv *RobotManageService) RobotDockerImagePull(ctx *gin.Context, progressChan chan<- dto.PullProgress) error {
	defer close(progressChan)
	// 创建Docker客户端
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		progressChan <- dto.PullProgress{
			Status: "error",
			Error:  fmt.Sprintf("创建Docker客户端失败: %v", err),
		}
		return fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	defer dockerClient.Close()
	// 定义需要拉取的镜像列表
	images := []string{
		"registry.cn-shenzhen.aliyuncs.com/houhou/wechat-ipad:latest",
		"registry.cn-shenzhen.aliyuncs.com/houhou/wechat-robot-client:latest",
	}
	// 逐个拉取镜像
	for _, image := range images {
		progressChan <- dto.PullProgress{
			Image:  image,
			Status: "start",
		}
		// 拉取镜像
		reader, err := dockerClient.ImagePull(sv.ctx, image, dockerImage.PullOptions{})
		if err != nil {
			progressChan <- dto.PullProgress{
				Image:  image,
				Status: "error",
				Error:  fmt.Sprintf("拉取镜像 %s 失败: %v", image, err),
			}
			return fmt.Errorf("拉取镜像 %s 失败: %v", image, err)
		}
		// 解析拉取进度
		err = sv.parseDockerPullProgress(reader, image, progressChan)
		reader.Close()
		if err != nil {
			progressChan <- dto.PullProgress{
				Image:  image,
				Status: "error",
				Error:  fmt.Sprintf("解析进度失败: %v", err),
			}
			return err
		}
		progressChan <- dto.PullProgress{
			Image:  image,
			Status: "complete",
		}
	}
	progressChan <- dto.PullProgress{
		Status: "all_complete",
	}
	return nil
}

// 解析Docker拉取进度
func (sv *RobotManageService) parseDockerPullProgress(reader io.ReadCloser, image string, progressChan chan<- dto.PullProgress) error {
	decoder := json.NewDecoder(reader)
	for {
		var progress map[string]interface{}
		if err := decoder.Decode(&progress); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		status, _ := progress["status"].(string)
		progressDetail, _ := progress["progressDetail"].(map[string]interface{})
		var progressStr string
		if progressDetail != nil {
			current, _ := progressDetail["current"].(float64)
			total, _ := progressDetail["total"].(float64)
			if total > 0 {
				percentage := (current / total) * 100
				progressStr = fmt.Sprintf("%.1f%%", percentage)
			}
		}
		progressChan <- dto.PullProgress{
			Image:    image,
			Status:   status,
			Progress: progressStr,
		}
	}
	return nil
}

// 辅助方法：停止并删除容器
func (sv *RobotManageService) stopAndRemoveContainer(dockerClient *client.Client, containerName string) error {
	// 根据容器名查找容器ID
	listFilters := filters.NewArgs()
	listFilters.Add("name", containerName)

	containers, err := dockerClient.ContainerList(sv.ctx, container.ListOptions{
		All:     true,
		Filters: listFilters,
	})
	if err != nil {
		return err
	}

	// 如果找不到容器，直接返回
	if len(containers) == 0 {
		return nil
	}

	// 容器存在，先停止
	timeout := 30
	err = dockerClient.ContainerStop(sv.ctx, containers[0].ID, container.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		return err
	}

	// 删除容器
	removeOptions := container.RemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}
	err = dockerClient.ContainerRemove(sv.ctx, containers[0].ID, removeOptions)
	if err != nil {
		return err
	}

	return nil
}

func (sv *RobotManageService) RobotRestart(robotID int64, restartType string) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot := respo.GetByID(robotID)
	if robot == nil {
		return errors.New("机器人不存在")
	}

	// 使用Docker SDK重启容器
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 根据重启类型确定容器名
	containerName := fmt.Sprintf("%s_%s", restartType, robot.RobotCode)

	// 根据容器名找到容器
	listFilters := filters.NewArgs()
	listFilters.Add("name", containerName)

	containers, err := dockerClient.ContainerList(sv.ctx, container.ListOptions{
		All:     true,
		Filters: listFilters,
	})
	if err != nil {
		return fmt.Errorf("查询容器失败: %v", err)
	}

	if len(containers) == 0 {
		return fmt.Errorf("找不到容器: %s", containerName)
	}

	// 重启容器
	timeout := 30
	err = dockerClient.ContainerRestart(sv.ctx, containers[0].ID, container.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		return fmt.Errorf("重启容器失败: %v", err)
	}

	return nil
}

func (sv *RobotManageService) RobotRestartClient(robotID int64) error {
	return sv.RobotRestart(robotID, "client")
}

func (sv *RobotManageService) RobotRestartServer(robotID int64) error {
	err := sv.RobotRestart(robotID, "server")
	if err != nil {
		return err
	}
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot := model.Robot{
		ID:     robotID,
		Status: model.RobotStatusOffline,
	}
	respo.Update(&robot)
	return nil
}
