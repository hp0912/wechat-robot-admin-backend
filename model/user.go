package model

import "gorm.io/gorm"

type User struct {
	ID          int64          `gorm:"primarykey" json:"id"`
	WeChatId    string         `json:"wechat_id" gorm:"column:wechat_id;index"`
	DisplayName string         `json:"display_name" gorm:"column:display_name;" validate:"max=20"`
	ApiToken    string         `json:"api_token" gorm:"column:api_token;type:varchar(128);index"`
	Role        int            `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status      int            `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	AvatarUrl   string         `json:"avatar_url" gorm:"type:varchar(500);column:avatar_url;default:''"`
	LastLoginAt int64          `json:"last_login_at" gorm:"bigint;default:0"`
	CreatedAt   int64          `json:"created_at" gorm:"bigint"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
