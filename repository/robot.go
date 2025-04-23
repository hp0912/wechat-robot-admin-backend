package repository

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
	"wechat-robot-admin-backend/vars"

	"gorm.io/gorm"
)

type Robot struct {
	Base[model.Robot]
}

func NewRobotRepo(ctx context.Context, db *gorm.DB) *Robot {
	return &Robot{
		Base[model.Robot]{
			Ctx: ctx,
			DB:  db,
		}}
}

func (r *Robot) RobotList(req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	query := vars.DB.Model(&model.Robot{})
	if req.Owner != "" {
		query = query.Where("owner = ?", req.Owner)
	}
	if req.Keyword != "" {
		query = query.Where("nickname LIKE ?", req.Keyword+"%").
			Or("robot_code LIKE ?", req.Keyword+"%").
			Or("wechat_id LIKE ?", req.Keyword+"%")
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}
	var robots []*model.Robot
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(pager.OffSet).Limit(pager.PageSize).Find(&robots).Error; err != nil {
		return nil, 0, err
	}
	return robots, total, nil
}

func (r *Robot) GetByOwner(owner string, unscoped bool, preloads ...string) []*model.Robot {
	filter := func(tx *gorm.DB) *gorm.DB {
		query := tx.Where("owner = ?", owner)
		if unscoped {
			query = query.Unscoped()
		}
		return query
	}
	return r.List(preloads, filter)
}

func (r *Robot) GetMaxRedisDB() (uint, error) {
	var maxDB uint
	if err := vars.DB.Model(&model.Robot{}).Unscoped().Select("COALESCE(MAX(redis_db), 0)").Scan(&maxDB).Error; err != nil {
		return 0, err
	}
	return maxDB, nil
}
