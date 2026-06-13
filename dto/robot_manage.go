package dto

import "wechat-robot-admin-backend/model"

// RobotResponse 表示机器人实例的接口响应结构。
type RobotResponse struct {
	ID            int64             `json:"id"`
	RobotCode     string            `json:"robot_code"`  // 当前机器人的唯一标识
	RobotName     string            `json:"robot_name"`  // 当前机器人的名称
	Proxy         *model.Proxy      `json:"proxy"`       // 当前机器人的代理配置
	Owner         string            `json:"owner"`       // 当前机器人的拥有者
	DeviceID      string            `json:"device_id"`   // 当前机器人登陆的设备Id
	DeviceName    string            `json:"device_name"` // 当前机器人登陆的设备名称
	WeChatID      string            `json:"wechat_id"`   // 当前机器人登陆的微信id
	Alias         string            `json:"alias"`       // 当前机器人登陆的自定义微信号
	BindMobile    string            `json:"bind_mobile"` // 当前机器人登陆的手机号
	Nickname      string            `json:"nickname"`    // 当前机器人登陆的微信昵称
	Avatar        string            `json:"avatar"`      // 当前机器人登陆的微信头像
	Status        model.RobotStatus `json:"status"`      // 当前机器人登陆的状态
	DBUsername    string            `json:"db_username"` // 当前机器人登陆的数据库用户名
	DBPassword    string            `json:"db_password"` // 当前机器人登陆的数据库密码
	RedisDB       uint64            `json:"redis_db"`    // 当前机器人登陆的Redis数据库
	ErrorMessage  string            `json:"error_message"`
	Profile       map[string]any    `json:"profile"`     // 当前机器人登陆的微信个人资料
	ProfileExt    map[string]any    `json:"profile_ext"` // 当前机器人登陆的扩展资料
	LastLoginAt   int64             `json:"last_login_at"`
	CreatedAt     int64             `json:"created_at"`
	UpdatedAt     int64             `json:"updated_at"`
	DeviceType    model.DeviceType  `json:"device_type"`
	WeChatVersion string            `json:"wechat_version"`
}

// 进度信息结构体
type PullProgress struct {
	Image    string `json:"image"`
	Status   string `json:"status"`
	Progress string `json:"progress,omitempty"`
	Error    string `json:"error,omitempty"`
}
