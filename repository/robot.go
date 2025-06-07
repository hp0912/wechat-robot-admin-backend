package repository

import (
	"context"
	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"

	"gorm.io/gorm"
)

type Robot struct {
	Ctx context.Context
	DB  *gorm.DB
}

func NewRobotRepo(ctx context.Context, db *gorm.DB) *Robot {
	return &Robot{
		Ctx: ctx,
		DB:  db,
	}
}

func (r *Robot) RobotList(req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	query := r.DB.WithContext(r.Ctx).Model(&model.Robot{})
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

func (r *Robot) GetByID(id int64) (*model.Robot, error) {
	var robot model.Robot
	err := r.DB.WithContext(r.Ctx).Where("id = ?", id).First(&robot).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &robot, nil
}

func (r *Robot) GetByOwner(owner string, unscoped bool) ([]*model.Robot, error) {
	query := r.DB.WithContext(r.Ctx).Model(&model.Robot{}).Where("owner = ?", owner)
	if unscoped {
		query = query.Unscoped()
	}
	var robots []*model.Robot
	if err := query.Find(&robots).Error; err != nil {
		return nil, err
	}
	return robots, nil
}

func (r *Robot) GetMaxRedisDB() (uint, error) {
	var maxDB uint
	if err := r.DB.WithContext(r.Ctx).Model(&model.Robot{}).Unscoped().Select("COALESCE(MAX(redis_db), 0)").Scan(&maxDB).Error; err != nil {
		return 0, err
	}
	return maxDB, nil
}

func (r *Robot) Create(data *model.Robot) error {
	return r.DB.WithContext(r.Ctx).Create(data).Error
}

func (r *Robot) Update(data *model.Robot) error {
	return r.DB.WithContext(r.Ctx).Where("id = ?", data.ID).Updates(data).Error
}

func (r *Robot) Delete(id int64) error {
	return r.DB.WithContext(r.Ctx).Where("id = ?", id).Delete(&model.Robot{}).Error
}
