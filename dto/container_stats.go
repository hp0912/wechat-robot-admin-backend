package dto

// ContainerStats 表示单个容器的资源使用统计信息
type ContainerStats struct {
	Name        string `json:"name"`         // 容器名称
	Status      string `json:"status"`       // 容器状态
	MemoryUsage string `json:"memory_usage"` // 内存使用情况
	CPUsage     string `json:"cpu_usage"`    // CPU使用情况
	DiskRead    string `json:"disk_read"`    // 磁盘读取
	DiskWrite   string `json:"disk_write"`   // 磁盘写入
}

// RobotContainerStatsResponse 包含机器人客户端和服务端容器的统计信息
type RobotContainerStatsResponse struct {
	Client ContainerStats `json:"client"` // 客户端容器统计
	Server ContainerStats `json:"server"` // 服务端容器统计
}
