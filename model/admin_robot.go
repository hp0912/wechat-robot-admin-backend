package model

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// RobotStatus 表示机器人状态的枚举类型
type RobotStatus string

const (
	RobotStatusOnline  RobotStatus = "online"
	RobotStatusOffline RobotStatus = "offline"
	RobotStatusError   RobotStatus = "error"
)

type DeviceType string

const (
	DeviceTypeiPad       DeviceType = "iPad"
	DeviceTypeiPhone     DeviceType = "iPhone"
	DeviceTypeAndroid    DeviceType = "Android"
	DeviceTypePadAndroid DeviceType = "Pad-Android"
	DeviceTypeWin        DeviceType = "Windows"
	DeviceTypeMac        DeviceType = "Mac"
	DeviceTypeCar        DeviceType = "Car"
	DeviceTypeUnknown    DeviceType = "Unknown"
)

// Robot 表示微信机器人实例
type Robot struct {
	ID           int64          `gorm:"primarykey" json:"id"`
	RobotCode    string         `gorm:"column:robot_code;index;unique,length:64" json:"robot_code"` // 当前机器人的唯一标识
	Owner        string         `gorm:"column:owner;index;length:64" json:"owner"`                  // 当前机器人的拥有者
	DeviceID     string         `gorm:"column:device_id;" json:"device_id"`                         // 当前机器人登陆的设备Id
	DeviceName   string         `gorm:"column:device_name" json:"device_name"`                      // 当前机器人登陆的设备名称
	WeChatID     string         `gorm:"column:wechat_id;index;length:64" json:"wechat_id"`          // 当前机器人登陆的微信id
	Alias        string         `gorm:"column:alias;length:64" json:"alias"`                        // 当前机器人登陆的自定义微信号
	BindMobile   string         `gorm:"column:bind_mobile" json:"bind_mobile"`                      // 当前机器人登陆的手机号
	Nickname     string         `gorm:"column:nickname" json:"nickname"`                            // 当前机器人登陆的微信昵称
	Avatar       string         `gorm:"column:avatar" json:"avatar"`                                // 当前机器人登陆的微信头像
	Status       RobotStatus    `gorm:"column:status;default:'offline'" json:"status"`              // 当前机器人登陆的状态
	RedisDB      uint           `gorm:"column:redis_db;default:1" json:"redis_db"`                  // 当前机器人登陆的Redis数据库
	ErrorMessage string         `gorm:"column:error_message" json:"error_message"`
	Profile      datatypes.JSON `gorm:"column:profile" json:"profile"` // 当前机器人登陆的微信个人资料
	ProfileExt   datatypes.JSON `gorm:"column:profile_ext" json:"profile_ext"`
	LastLoginAt  int64          `gorm:"column:last_login_at" json:"last_login_at"`
	// 代理设置
	ProxyEnabled  bool           `gorm:"column:proxy_enabled;default:false" json:"proxy_enabled"` // 是否启用代理
	ProxyIP       string         `gorm:"column:proxy_ip" json:"proxy_ip"`                         // 代理IP地址，格式: ip:port
	ProxyUser     string         `gorm:"column:proxy_user" json:"proxy_user"`                     // 代理用户名（可选）
	ProxyPassword string         `gorm:"column:proxy_password" json:"proxy_password"`             // 代理密码（可选）
	CreatedAt     int64          `json:"created_at"`
	UpdatedAt     int64          `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	DeviceType    DeviceType     `gorm:"->;<-:false" json:"device_type"`
	WeChatVersion string         `gorm:"->;<-:false" json:"wechat_version"`
}

// TableName 指定表名
func (Robot) TableName() string {
	return "robot"
}

func (r *Robot) GetBaseURL() string {
	if os.Getenv("DEV_ROBOT_CLIENT_URL") != "" {
		return os.Getenv("DEV_ROBOT_CLIENT_URL")
	}
	return fmt.Sprintf("http://client_%s:%d/api/v1/robot", r.RobotCode, 9000)
}

func (r *Robot) ParseDeviceType(raw string) DeviceType {
	deviceType := strings.ToLower(raw)
	if strings.Contains(deviceType, "ipad") {
		return DeviceTypeiPad
	} else if strings.Contains(deviceType, "iphone") {
		return DeviceTypeiPhone
	} else if strings.Contains(deviceType, "android") {
		if strings.Contains(deviceType, "pad") {
			return DeviceTypePadAndroid
		}
		return DeviceTypeAndroid
	} else if strings.Contains(deviceType, "windows") {
		return DeviceTypeWin
	} else if strings.Contains(deviceType, "mac") {
		return DeviceTypeMac
	} else if strings.Contains(deviceType, "car") {
		return DeviceTypeCar
	}
	return DeviceTypeUnknown
}

func (r *Robot) ParseDeviceVersion(version int32) string {
	if version <= 0 {
		return "Unknown"
	}
	major := 0x0f & (version >> 24)
	minor := 0xff & (version >> 16)
	patch := 0xff & (version >> 8)
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// GetProxyInfo 获取代理信息，返回ProxyInfo结构体
func (r *Robot) GetProxyInfo() *ProxyInfo {
	if !r.ProxyEnabled || r.ProxyIP == "" {
		return nil
	}
	return &ProxyInfo{
		ProxyIp:       r.ProxyIP,
		ProxyUser:     r.ProxyUser,
		ProxyPassword: r.ProxyPassword,
	}
}

// ProxyInfo 代理信息结构体（与client包中的ProxyInfo保持一致）
type ProxyInfo struct {
	ProxyIp       string `json:"ProxyIp"`
	ProxyUser     string `json:"ProxyUser"`
	ProxyPassword string `json:"ProxyPassword"`
}
