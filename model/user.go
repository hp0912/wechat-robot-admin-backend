package model

import "gorm.io/gorm"

type User struct {
	Id            int64          `json:"id"`
	WeChatId      string         `json:"wechat_id" gorm:"column:wechat_id;index"`
	DisplayName   string         `json:"display_name" gorm:"index" validate:"max=20"`
	Role          int            `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status        int            `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	AvatarUrl     string         `json:"avatar_url" gorm:"type:varchar(500);column:avatar_url;default:''"`
	LastLoginTime int64          `json:"last_login_time" gorm:"bigint;default:0"`
	CreatedTime   int64          `json:"created_time" gorm:"bigint"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
