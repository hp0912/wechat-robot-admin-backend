package repository

import (
	"context"
	"wechat-robot-client/model"

	"gorm.io/gorm"
)

type User struct {
	Base[model.User]
}

func NewUserRepo(ctx context.Context, db *gorm.DB) *User {
	return &User{
		Base[model.User]{
			Ctx: ctx,
			DB:  db,
		}}
}

func (u *User) GetUserByID(id int64, preloads ...string) *model.User {
	return u.takeOne(preloads, func(g *gorm.DB) *gorm.DB {
		return g.Where("id = ?", id)
	})
}

func (u *User) GetUserByWeChatID(WxID string, preloads ...string) *model.User {
	return u.takeOne(preloads, func(g *gorm.DB) *gorm.DB {
		return g.Where("wechat_id = ?", WxID)
	})
}
