package model

import (
	"gorm.io/gorm"
)

// RobotStatus 表示机器人状态的枚举类型
type RobotStatus string

const (
	RobotStatusOnline  RobotStatus = "online"
	RobotStatusOffline RobotStatus = "offline"
	RobotStatusError   RobotStatus = "error"
)

// Robot 表示微信机器人实例
type Robot struct {
	ID           int64          `gorm:"primarykey" json:"id"`
	RobotCode    string         `gorm:"column:robot_code;index;unique,length:64" json:"robot_code"` // 当前机器人的唯一标识
	Owner        string         `gorm:"column:owner;index;length:64" json:"owner"`                  // 当前机器人的拥有者
	DeviceId     string         `gorm:"column:device_id;" json:"device_id"`                         // 当前机器人登陆的设备Id
	DeviceName   string         `gorm:"column:device_name" json:"device_mame"`                      // 当前机器人登陆的设备名称
	WeChatID     string         `gorm:"column:wechat_id;index;length:64" json:"wechat_id"`          // 当前机器人登陆的微信id
	Nickname     string         `gorm:"column:nickname" json:"nickname"`                            // 当前机器人登陆的微信昵称
	Avatar       string         `gorm:"column:avatar" json:"avatar"`                                // 当前机器人登陆的微信头像
	Status       RobotStatus    `gorm:"column:status;default:'offline'" json:"status"`              // 当前机器人登陆的状态
	RedisDB      uint           `gorm:"column:redis_db;default:1" json:"redis_db"`                  // 当前机器人登陆的Redis数据库
	ErrorMessage string         `gorm:"column:error_message" json:"error_message"`
	LastLoginAt  int64          `gorm:"column:last_login_at" json:"last_login_at"`
	CreatedAt    int64          `json:"created_at"`
	UpdatedAt    int64          `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (Robot) TableName() string {
	return "robot"
}
