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
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RobotCode     string         `gorm:"column:robot_code;type:varchar(64);not null;uniqueIndex:idx_robot_code" json:"robot_code"`         // 当前机器人的唯一标识
	RobotName     string         `gorm:"column:robot_name;type:varchar(255);not null;default:'WeChat Robot'" json:"robot_name"`            // 当前机器人的名称
	Owner         string         `gorm:"column:owner;type:varchar(64);not null;index:idx_owner" json:"owner"`                              // 当前机器人的拥有者
	DeviceID      string         `gorm:"column:device_id;type:varchar(255)" json:"device_id"`                                              // 当前机器人登陆的设备Id
	DeviceName    string         `gorm:"column:device_name;type:varchar(255)" json:"device_name"`                                          // 当前机器人登陆的设备名称
	WeChatID      string         `gorm:"column:wechat_id;type:varchar(64);index:idx_wechat_id" json:"wechat_id"`                           // 当前机器人登陆的微信id
	Alias         string         `gorm:"column:alias;type:varchar(64)" json:"alias"`                                                       // 当前机器人登陆的自定义微信号
	BindMobile    string         `gorm:"column:bind_mobile;type:varchar(15)" json:"bind_mobile"`                                           // 当前机器人登陆的手机号
	Nickname      string         `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                                                // 当前机器人登陆的微信昵称
	Avatar        string         `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                                                    // 当前机器人登陆的微信头像
	Status        RobotStatus    `gorm:"column:status;type:enum('online','offline','error');not null;default:'offline'" json:"status"`     // 当前机器人登陆的状态
	DBUsername    string         `gorm:"column:db_username;type:varchar(64);not null;default:''" json:"db_username"`                       // 当前机器人登陆的数据库用户名
	DBPassword    string         `gorm:"column:db_password;type:varchar(255);not null;default:''" json:"db_password"`                      // 当前机器人登陆的数据库密码
	RedisDB       uint64         `gorm:"column:redis_db;type:bigint unsigned;not null;default:1;uniqueIndex:idx_redis_db" json:"redis_db"` // 当前机器人登陆的Redis数据库
	ErrorMessage  string         `gorm:"column:error_message;type:text" json:"error_message"`
	Profile       datatypes.JSON `gorm:"column:profile;type:json" json:"profile"` // 当前机器人登陆的微信个人资料
	ProfileExt    datatypes.JSON `gorm:"column:profile_ext;type:json" json:"profile_ext"`
	LastLoginAt   int64          `gorm:"column:last_login_at;type:bigint" json:"last_login_at"`
	CreatedAt     int64          `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt     int64          `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index:idx_deleted_at" json:"-"`
	DeviceType    DeviceType     `gorm:"->;<-:false;-:migration" json:"device_type"`
	WeChatVersion string         `gorm:"->;<-:false;-:migration" json:"wechat_version"`
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
