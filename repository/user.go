package repository

import (
	"context"
	"wechat-robot-admin-backend/model"

	"gorm.io/gorm"
)

type User struct {
	Ctx context.Context
	DB  *gorm.DB
}

func NewUserRepo(ctx context.Context, db *gorm.DB) *User {
	return &User{
		Ctx: ctx,
		DB:  db,
	}
}

func (u *User) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	err := u.DB.WithContext(u.Ctx).Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) GetUserByWeChatID(WxID string) (*model.User, error) {
	var user model.User
	err := u.DB.WithContext(u.Ctx).Where("wechat_id = ?", WxID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) GetUserByApiToken(token string) (*model.User, error) {
	var user model.User
	err := u.DB.WithContext(u.Ctx).Where("api_token = ?", token).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UserCount() (int64, error) {
	var count int64
	err := u.DB.WithContext(u.Ctx).Model(&model.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *User) Create(data *model.User) error {
	return u.DB.WithContext(u.Ctx).Create(data).Error
}

func (u *User) Update(data *model.User) error {
	return u.DB.WithContext(u.Ctx).Where("id = ?", data.ID).Updates(data).Error
}
