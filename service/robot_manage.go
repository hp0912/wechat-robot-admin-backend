package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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
	"github.com/go-resty/resty/v2"
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
	if pager.PageSize <= 0 || pager.PageSize > 50 {
		return nil, 0, errors.New("参数异常")
	}
	robots, total, err := repository.NewRobotRepo(sv.ctx, vars.DB).RobotList(req, pager)
	if err != nil {
		return nil, 0, err
	}
	client := resty.New()
	client.SetTimeout(2 * time.Second)
	for index := range robots {
		var robotLoginData dto.Response[dto.RobotLoginData]
		_, err := client.R().
			SetHeader("Content-Type", "application/json;chartset=utf-8").
			SetResult(&robotLoginData).
			Get(robots[index].GetBaseURL() + "/get-cached-info")
		if err = robotLoginData.CheckError(err); err != nil {
			//
		}
		robots[index].DeviceType = robots[index].ParseDeviceType(robotLoginData.Data.DeviceType)
		robots[index].WeChatVersion = robots[index].ParseDeviceVersion(robotLoginData.Data.ClientVersion)
	}
	return robots, total, nil
}

// 辅助方法：获取Docker客户端
func (sv *RobotManageService) getDockerClient() (*client.Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	return dockerClient, nil
}

func (sv *RobotManageService) DockerStartWeChatClient(ctx *gin.Context, robot *model.Robot) error {
	// 创建Docker客户端
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 需要的镜像
	clientImage := "registry.cn-shenzhen.aliyuncs.com/houhou/wechat-robot-client:latest"
	// 确保镜像存在，不存在则先拉取
	if err := sv.ensureImage(sv.ctx, dockerClient, clientImage); err != nil {
		return fmt.Errorf("准备客户端镜像失败: %v", err)
	}

	// 客户端容器配置
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	clientConfig := &container.Config{
		Image: clientImage,
		Env: []string{
			fmt.Sprintf("GIN_MODE=%s", "release"),
			fmt.Sprintf("WECHAT_CLIENT_PORT=%s", "9000"),
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
			fmt.Sprintf("REDIS_HOST=%s", vars.RedisSettings.Host),
			fmt.Sprintf("REDIS_PORT=%s", vars.RedisSettings.Port),
			fmt.Sprintf("REDIS_PASSWORD=%s", vars.RedisSettings.Password),
			fmt.Sprintf("REDIS_DB=%d", robot.RedisDB),
			fmt.Sprintf("WORD_CLOUD_URL=%s", "http://word-cloud-server:9000/api/v1/word-cloud/gen"),
			fmt.Sprintf("THIRD_PARTY_API_KEY=%s", vars.ThirdPartyApiKey),
		},
	}

	clientHostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{"max-size": "50m", "max-file": "3"},
		},
	}

	if vars.HostDataDir != "" {
		hostSkillsDir := fmt.Sprintf("%s/wechat-robot/%s/data/skills", vars.HostDataDir, robot.RobotCode)
		clientHostConfig.Binds = []string{
			fmt.Sprintf("%s:/data/skills", hostSkillsDir),
		}
	}

	// 为该机器人创建（或复用）独立隔离网络，并将公共服务容器接入。
	// 动态容器仅加入此网络，不加入公共网络，从而实现跨组网络隔离。
	if _, err := sv.ensureRobotNetwork(dockerClient, robot.RobotCode); err != nil {
		return fmt.Errorf("创建机器人隔离网络失败: %v", err)
	}

	// 客户端网络配置
	clientNetworkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			sv.robotNetworkName(robot.RobotCode): {},
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

func (sv *RobotManageService) DockerStartWeChatServer(ctx *gin.Context, robot *model.Robot, pprofEnable bool) error {
	// 创建Docker客户端
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 需要的镜像
	serverImage := "registry.cn-shenzhen.aliyuncs.com/houhou/wechat-ipad:latest"
	// 确保镜像存在，不存在则先拉取
	if err := sv.ensureImage(sv.ctx, dockerClient, serverImage); err != nil {
		return fmt.Errorf("准备服务端镜像失败: %v", err)
	}

	// 服务端容器配置
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	serverContainerName := fmt.Sprintf("server_%s", robot.RobotCode)
	serverConfig := &container.Config{
		Image: serverImage,
		Env: []string{
			fmt.Sprintf("WECHAT_PORT=%s", "9000"),
			fmt.Sprintf("REDIS_HOST=%s", vars.RedisSettings.Host),
			fmt.Sprintf("REDIS_PORT=%s", vars.RedisSettings.Port),
			fmt.Sprintf("REDIS_PASSWORD=%s", vars.RedisSettings.Password),
			fmt.Sprintf("REDIS_DB=%d", robot.RedisDB),
			fmt.Sprintf("WECHAT_CLIENT_HOST=%s", fmt.Sprintf("%s:%d", clientContainerName, 9000)),
			fmt.Sprintf("UUID_URL=%s", vars.UUIDURL),
		},
	}
	if pprofEnable {
		serverConfig.Env = append(serverConfig.Env, "ENABLE_PPROF=true")
		serverConfig.Env = append(serverConfig.Env, "PPROF_ADDR=0.0.0.0")
		serverConfig.Env = append(serverConfig.Env, "PPROF_PORT=9010")
	}

	serverHostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{"max-size": "50m", "max-file": "3"},
		},
	}

	// 服务端网络配置：与 client 使用同一隔离网络，确保同组互通、跨组隔离。
	serverNetworkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			sv.robotNetworkName(robot.RobotCode): {},
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

	return nil
}

func (sv *RobotManageService) DockerStopAndRemoveWeChatClient(ctx *gin.Context, robot *model.Robot) error {
	// 使用Docker SDK停止并删除容器
	dockerClient, err := sv.getDockerClient()
	if err != nil {
		return err
	}
	defer dockerClient.Close()

	// 停止并删除客户端容器
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	err = sv.stopAndRemoveContainer(dockerClient, clientContainerName)
	if err != nil {
		return fmt.Errorf("删除客户端容器失败: %v", err)
	}

	return nil
}

func (sv *RobotManageService) DockerStopAndRemoveWeChatServer(ctx *gin.Context, robot *model.Robot) error {
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
	robots, err := respo.GetByOwner(wechatId.(string), true)
	if err != nil {
		return err
	}
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
	err = respo.Create(robot)
	if err != nil {
		return err
	}
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
	commonConf := fmt.Sprintf("INSERT INTO `%s`.`%s` (`chat_ai_enabled`, `chat_base_url`, `chat_api_key`, `chat_model`, `image_recognition_model`, `chat_prompt`, `friend_sync_cron`) VALUES (0, '%s', '%s', 'gpt-4o-mini', 'gpt-4o-mini', '%s', '55 * * * *');",
		robot.RobotCode, "global_settings", "https://new-api.houhoukang.com/", vars.OpenAIApiKey, "你是一个聊天机器人。")
	err = newDB.Exec(commonConf).Error
	if err != nil {
		return err
	}

	// 插入一条官方 MCP 服务配置
	mcpServerConf := fmt.Sprintf("INSERT INTO `%s`.mcp_servers (name, is_built_in, description, transport, enabled, priority, command, args, working_dir, env, url, client_name, auth_type, auth_token, auth_username, auth_password, headers, tls_skip_verify, connect_timeout, read_timeout, write_timeout, max_retries, retry_interval, heartbeat_enable, heartbeat_interval, capabilities, custom_config, tags, last_connected_at, last_error, connection_count, error_count, created_at, updated_at, deleted_at) VALUES ('BuiltInPlugin', 1, '官方内置 MCP 服务', 'stream', 1, 100, '', 'null', '', '{}', 'http://wechat-robot-mcp-server:9000/mcp', '', 'none', '', '', '', '{}', 0, 30, 60, 60, 3, 5, 1, 60, 'null', 'null', '[\"官方\", \"群聊总结\"]', null, '', 0, 0, '2025-11-14 21:28:26', '2025-11-14 21:28:26', null);", robot.RobotCode)
	err = newDB.Exec(mcpServerConf).Error
	if err != nil {
		return err
	}

	err = sv.DockerStartWeChatClient(ctx, robot)
	if err != nil {
		return err
	}

	err = sv.DockerStartWeChatServer(ctx, robot, false)
	if err != nil {
		return err
	}

	return nil
}

// RobotView 查看机器人元数据
func (sv *RobotManageService) RobotView(robotID int64) (*model.Robot, error) {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return nil, err
	}
	if robot == nil {
		return nil, nil
	}
	client := resty.New()
	client.SetTimeout(2 * time.Second)
	var robotLoginData dto.Response[dto.RobotLoginData]
	_, err = client.R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&robotLoginData).
		Get(robot.GetBaseURL() + "/get-cached-info")
	if err = robotLoginData.CheckError(err); err != nil {
		//
	}
	robot.DeviceType = robot.ParseDeviceType(robotLoginData.Data.DeviceType)
	robot.WeChatVersion = robot.ParseDeviceVersion(robotLoginData.Data.ClientVersion)
	return robot, nil
}

// RobotStopAndRemoveWeChatClient 删除机器人客户端容器
func (sv *RobotManageService) RobotStopAndRemoveWeChatClient(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return err
	}
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err = sv.DockerStopAndRemoveWeChatClient(ctx, robot)
	if err != nil {
		return err
	}
	return nil
}

// RobotStopAndRemoveClientAndServer 删除机器人容器
func (sv *RobotManageService) RobotStopAndRemoveWeChatServer(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return err
	}
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err = sv.DockerStopAndRemoveWeChatServer(ctx, robot)
	if err != nil {
		return err
	}
	return nil
}

// RobotStartWeChatClient 启动机器人客户端容器
func (sv *RobotManageService) RobotStartWeChatClient(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return err
	}
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err = sv.DockerStartWeChatClient(ctx, robot)
	if err != nil {
		return err
	}
	return nil
}

// RobotStartWeChatServer 启动机器人服务端容器
func (sv *RobotManageService) RobotStartWeChatServer(ctx *gin.Context, req dto.RobotStartServerRequest) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(req.ID)
	if err != nil {
		return err
	}
	if robot == nil {
		return errors.New("机器人不存在")
	}
	err = sv.DockerStartWeChatServer(ctx, robot, req.PprofEnable)
	if err != nil {
		return err
	}
	return nil
}

// RobotRemove 删除机器人
func (sv *RobotManageService) RobotRemove(ctx *gin.Context, robotID int64) error {
	respo := repository.NewRobotRepo(sv.ctx, vars.DB)
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return err
	}
	if robot == nil {
		return errors.New("机器人不存在")
	}

	// 先尝试退出登录
	robotLoginService := NewRobotLoginService(sv.ctx)
	err = robotLoginService.RobotLogout(robot)
	if err != nil {
		log.Println("删除机器人容器前，机器人登出失败:", err)
	}

	// 删除机器人实例数据
	err = respo.Delete(robotID)
	if err != nil {
		return err
	}

	// 删除机器人数据库
	err = vars.DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`;", robot.RobotCode)).Error
	if err != nil {
		return err
	}

	err = sv.DockerStopAndRemoveWeChatClient(ctx, robot)
	if err != nil {
		return err
	}

	err = sv.DockerStopAndRemoveWeChatServer(ctx, robot)
	if err != nil {
		return err
	}

	// 两个容器都已删除后，清理该机器人的隔离网络。
	if dc, dcErr := sv.getDockerClient(); dcErr == nil {
		defer dc.Close()
		if netErr := sv.removeRobotNetwork(dc, robot.RobotCode); netErr != nil {
			log.Printf("删除机器人隔离网络失败: %v", netErr)
		}
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
	robot, err := respo.GetByID(robotID)
	if err != nil {
		return err
	}
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
	timeout := 60
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
	return sv.RobotRestart(robotID, "server")
}

func (sv *RobotManageService) robotNetworkName(robotCode string) string {
	return fmt.Sprintf("robot-%s", robotCode)
}

func (sv *RobotManageService) robotNetworkPublicContainers() []string {
	return []string{
		"wechat-admin-mysql",
		"wechat-admin-redis",
		"wechat-server",
		"wechat-slider",
		"wechat-uuid",
		"word-cloud-server",
		"jimeng-api",
		"wechat-robot-mcp-server",
		"xiaohongshu-mcp",
		"wechat-robot-admin-backend",
		"netease-cloud-music",
	}
}

// ensureRobotNetwork creates the per-robot isolated bridge network if it does not
// already exist, then connects all required public service containers to it so
// that the robot pair can reach shared dependencies. The robot containers
// themselves are attached ONLY to this network; they are never added to
// vars.DockerNetwork, so containers in different robot groups cannot reach each
// other even though they all share the same public services.
func (sv *RobotManageService) ensureRobotNetwork(dockerClient *client.Client, robotCode string) (string, error) {
	networkName := sv.robotNetworkName(robotCode)

	// Look up existing network by exact name.
	netFilters := filters.NewArgs()
	netFilters.Add("name", networkName)
	nets, err := dockerClient.NetworkList(sv.ctx, network.ListOptions{Filters: netFilters})
	if err != nil {
		return "", fmt.Errorf("查询网络失败: %v", err)
	}

	var networkID string
	for _, n := range nets {
		if n.Name == networkName {
			networkID = n.ID
			break
		}
	}

	if networkID == "" {
		resp, err := dockerClient.NetworkCreate(sv.ctx, networkName, network.CreateOptions{
			Driver: "bridge",
		})
		if err != nil {
			return "", fmt.Errorf("创建网络 %s 失败: %v", networkName, err)
		}
		networkID = resp.ID
	}

	// Connect required public service containers (best-effort; already-connected is fine).
	for _, ctrName := range sv.robotNetworkPublicContainers() {
		ctrFilters := filters.NewArgs()
		ctrFilters.Add("name", ctrName)
		ctrs, err := dockerClient.ContainerList(sv.ctx, container.ListOptions{All: true, Filters: ctrFilters})
		if err != nil || len(ctrs) == 0 {
			continue
		}
		if err := dockerClient.NetworkConnect(sv.ctx, networkID, ctrs[0].ID, nil); err != nil {
			// Ignore "already exists in network" errors; log everything else.
			log.Printf("连接容器 %s 到网络 %s: %v", ctrName, networkName, err)
		}
	}

	return networkID, nil
}

// removeRobotNetwork disconnects all public service containers from the per-robot
// network and then removes the network itself. Called when a robot is fully deleted.
func (sv *RobotManageService) removeRobotNetwork(dockerClient *client.Client, robotCode string) error {
	networkName := sv.robotNetworkName(robotCode)

	netFilters := filters.NewArgs()
	netFilters.Add("name", networkName)
	nets, err := dockerClient.NetworkList(sv.ctx, network.ListOptions{Filters: netFilters})
	if err != nil {
		return fmt.Errorf("查询网络失败: %v", err)
	}

	var networkID string
	for _, n := range nets {
		if n.Name == networkName {
			networkID = n.ID
			break
		}
	}
	if networkID == "" {
		return nil // Already gone.
	}

	for _, ctrName := range sv.robotNetworkPublicContainers() {
		ctrFilters := filters.NewArgs()
		ctrFilters.Add("name", ctrName)
		ctrs, err := dockerClient.ContainerList(sv.ctx, container.ListOptions{All: true, Filters: ctrFilters})
		if err != nil || len(ctrs) == 0 {
			continue
		}
		_ = dockerClient.NetworkDisconnect(sv.ctx, networkID, ctrs[0].ID, true)
	}

	if err := dockerClient.NetworkRemove(sv.ctx, networkID); err != nil {
		return fmt.Errorf("删除网络 %s 失败: %v", networkName, err)
	}
	return nil
}

// ensureImage 确保镜像已存在；若不存在则拉取
func (sv *RobotManageService) ensureImage(ctx context.Context, dockerClient *client.Client, image string) error {
	// 先尝试 inspect
	if _, err := dockerClient.ImageInspect(ctx, image); err == nil {
		return nil // 已存在
	}
	// 不存在则拉取
	reader, err := dockerClient.ImagePull(ctx, image, dockerImage.PullOptions{})
	if err != nil {
		return fmt.Errorf("拉取镜像 %s 失败: %w", image, err)
	}
	defer reader.Close()
	// 读取完输出以便 docker 正常完成（忽略具体进度，这里只是确保镜像到位）
	_, _ = io.Copy(io.Discard, reader)
	// 再次确认
	if _, err := dockerClient.ImageInspect(ctx, image); err != nil {
		return fmt.Errorf("镜像 %s 拉取后仍不可用: %w", image, err)
	}
	return nil
}

func (sv *RobotManageService) ExportRobotLoginData(robot *model.Robot) (string, error) {
	client := resty.New()
	client.SetTimeout(2 * time.Second)
	var robotLoginData dto.Response[dto.RobotLoginData]
	_, err := client.R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&robotLoginData).
		Get(robot.GetBaseURL() + "/get-cached-info")
	if err = robotLoginData.CheckError(err); err != nil {
		return "", err
	}
	dataBytes, err := json.MarshalIndent(robotLoginData.Data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(dataBytes), nil
}

func (sv *RobotManageService) ImportRobotLoginData(robot *model.Robot, data string) error {
	client := resty.New()
	var robotLoginData dto.Response[struct{}]
	_, err := client.R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]string{
			"data": data,
		}).
		SetResult(&robotLoginData).
		Post(robot.GetBaseURL() + "/import-login-data")
	if err = robotLoginData.CheckError(err); err != nil {
		return err
	}
	return nil
}
