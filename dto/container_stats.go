package dto

type MemoryUsage struct {
	Usage   string `json:"usage"`   // 内存使用情况
	Limit   string `json:"limit"`   // 内存限制
	Percent string `json:"percent"` // 内存使用百分比
}

// ContainerStats 表示单个容器的资源使用统计信息
type ContainerStats struct {
	Name        string      `json:"name"`         // 容器名称
	Status      string      `json:"status"`       // 容器状态
	MemoryUsage MemoryUsage `json:"memory_usage"` // 内存使用情况
	CPUsage     string      `json:"cpu_usage"`    // CPU使用情况
	DiskRead    string      `json:"disk_read"`    // 磁盘读取
	DiskWrite   string      `json:"disk_write"`   // 磁盘写入
}

// RobotContainerStatsResponse 包含机器人客户端和服务端容器的统计信息
type RobotContainerStatsResponse struct {
	Client ContainerStats `json:"client"` // 客户端容器统计
	Server ContainerStats `json:"server"` // 服务端容器统计
}

// RobotContainerLogsResponse 包含机器人客户端和服务端容器的日志
type RobotContainerLogsResponse struct {
	Client []string `json:"client"` // 客户端容器日志
	Server []string `json:"server"` // 服务端容器日志
}
