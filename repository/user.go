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

func (r *User) GetUserByWechatID(robotID string, preloads ...string) *model.User {
	return r.takeOne(preloads, func(g *gorm.DB) *gorm.DB {
		return g.Where("wechat_id = ?", robotID)
	})
}
