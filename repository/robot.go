package repository

import (
	"context"
	"wechat-robot-client/dto"
	"wechat-robot-client/model"
	"wechat-robot-client/pkg/appx"
	"wechat-robot-client/vars"

	"github.com/gin-gonic/gin"
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

func (r *Robot) RobotList(ctx *gin.Context, req dto.RobotListRequest, pager appx.Pager) ([]*model.Robot, int64, error) {
	query := vars.DB.Model(&model.Robot{})
	if req.Owner != "" {
		query = query.Where("owner = ?", req.Owner)
	}
	if req.Keyword != "" {
		query = query.Where("nickname LIKE ?", "%"+req.Keyword)
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

func (r *Robot) GetMaxRedisDB() (uint, error) {
	var maxDB uint
	if err := vars.DB.Model(&model.Robot{}).Unscoped().Select("COALESCE(MAX(redis_db), 0)").Scan(&maxDB).Error; err != nil {
		return 0, err
	}
	return maxDB, nil
}
