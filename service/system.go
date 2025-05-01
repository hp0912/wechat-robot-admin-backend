package service

import (
	"context"
	"encoding/json"
	"fmt"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type SystemService struct {
	ctx context.Context
}

func NewSystemService(ctx context.Context) *SystemService {
	return &SystemService{
		ctx: ctx,
	}
}

func (s *SystemService) getDockerClient() (*client.Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("创建Docker客户端失败: %v", err)
	}
	return dockerClient, nil
}

// RobotContainerStats 获取机器人容器的资源使用情况
func (s *SystemService) RobotContainerStats(robot *model.Robot) (dto.RobotContainerStatsResponse, error) {
	// 创建Docker客户端
	dockerClient, err := s.getDockerClient()
	if err != nil {
		return dto.RobotContainerStatsResponse{}, err
	}
	defer dockerClient.Close()

	// 获取客户端容器的状态
	clientContainerName := fmt.Sprintf("client_%s", robot.RobotCode)
	clientStats, err := s.getContainerStats(dockerClient, clientContainerName)
	if err != nil {
		return dto.RobotContainerStatsResponse{}, fmt.Errorf("获取客户端容器状态失败: %v", err)
	}

	// 获取服务端容器的状态
	serverContainerName := fmt.Sprintf("server_%s", robot.RobotCode)
	serverStats, err := s.getContainerStats(dockerClient, serverContainerName)
	if err != nil {
		return dto.RobotContainerStatsResponse{}, fmt.Errorf("获取服务端容器状态失败: %v", err)
	}

	return dto.RobotContainerStatsResponse{
		Client: clientStats,
		Server: serverStats,
	}, nil
}

// getContainerStats 获取单个容器的资源使用情况
func (s *SystemService) getContainerStats(dockerClient *client.Client, containerName string) (dto.ContainerStats, error) {
	// 初始化返回结构
	stats := dto.ContainerStats{
		Name: containerName,
	}

	// 根据容器名查找容器
	listFilters := filters.NewArgs()
	listFilters.Add("name", containerName)

	containers, err := dockerClient.ContainerList(s.ctx, container.ListOptions{
		All:     true,
		Filters: listFilters,
	})
	if err != nil {
		return stats, err
	}

	if len(containers) == 0 {
		return stats, fmt.Errorf("找不到容器: %s", containerName)
	}

	containerID := containers[0].ID
	stats.Status = string(containers[0].State)

	// 获取容器统计信息
	containerStats, err := dockerClient.ContainerStats(s.ctx, containerID, false)
	if err != nil {
		return stats, err
	}
	defer containerStats.Body.Close()

	// 解析统计数据
	var statsJSON container.StatsResponse
	if err := json.NewDecoder(containerStats.Body).Decode(&statsJSON); err != nil {
		return stats, err
	}

	// 计算内存使用率
	memUsage := float64(statsJSON.MemoryStats.Usage)
	memLimit := float64(statsJSON.MemoryStats.Limit)
	memUsageMB := memUsage / 1024 / 1024
	memLimitMB := memLimit / 1024 / 1024
	memPercent := 0.0
	if memLimit > 0 {
		memPercent = (memUsage / memLimit) * 100
	}
	stats.MemoryUsage = fmt.Sprintf("%.2f MB / %.2f MB (%.2f%%)", memUsageMB, memLimitMB, memPercent)

	// 计算CPU使用率
	var cpuPercent float64
	cpuDelta := float64(statsJSON.CPUStats.CPUUsage.TotalUsage) - float64(statsJSON.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(statsJSON.CPUStats.SystemUsage) - float64(statsJSON.PreCPUStats.SystemUsage)

	if systemDelta > 0 && cpuDelta > 0 {
		numCPUs := float64(len(statsJSON.CPUStats.CPUUsage.PercpuUsage))
		// 如果numCPUs为0，则默认为1，避免除以0的错误
		if numCPUs == 0 {
			numCPUs = 1
		}
		cpuPercent = (cpuDelta / systemDelta) * numCPUs * 100.0
	}
	stats.CPUsage = fmt.Sprintf("%.2f%%", cpuPercent)

	// 计算磁盘I/O
	var readBytes, writeBytes uint64
	for _, blkioStat := range statsJSON.BlkioStats.IoServiceBytesRecursive {
		switch blkioStat.Op {
		case "Read":
			readBytes += blkioStat.Value
		case "Write":
			writeBytes += blkioStat.Value
		}
	}
	readMB := float64(readBytes) / 1024 / 1024
	writeMB := float64(writeBytes) / 1024 / 1024
	stats.DiskRead = fmt.Sprintf("%.2f MB", readMB)
	stats.DiskWrite = fmt.Sprintf("%.2f MB", writeMB)

	return stats, nil
}
