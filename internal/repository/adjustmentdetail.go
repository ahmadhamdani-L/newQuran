package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdjustmentDetail interface {
	Find(ctx *abstraction.Context, m *model.AdjustmentDetailFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentDetailEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentDetailEntityModel, error)
	Create(ctx *abstraction.Context, e *model.AdjustmentDetailEntityModel) (*model.AdjustmentDetailEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.AdjustmentDetailEntity) (*model.AdjustmentDetailEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentDetailEntityModel) (*model.AdjustmentDetailEntityModel, error)
}

type adjustmentdetail struct {
	abstraction.Repository
}

func NewAdjustmentDetail(db *gorm.DB) *adjustmentdetail {
	return &adjustmentdetail{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *adjustmentdetail) Find(ctx *abstraction.Context, m *model.AdjustmentDetailFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentDetailEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.AdjustmentDetailEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.AdjustmentDetailEntityModel{})

	// filter
	query = r.Filter(ctx, query, m)

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "created_at"
		p.SortBy = &sortBy
	}
	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	// pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}
	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *adjustmentdetail) FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentDetailEntityModel

	

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
		conn.Preload(clause.Associations).Find(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustmentdetail) Create(ctx *abstraction.Context, e *model.AdjustmentDetailEntityModel) (*model.AdjustmentDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil
}

// func (r *trialbalance) Create(ctx *abstraction.Context, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error) {
// 	conn := r.CheckTrx(ctx)

// 	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := conn.Model(e).Preload("Company").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}

// 	return e, nil
// }

func (r *adjustmentdetail) Update(ctx *abstraction.Context, id *int, e *model.AdjustmentDetailEntity) (*model.AdjustmentDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentDetailEntityModel

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	data.AdjustmentDetailEntity = *e
	err = conn.Model(data).UpdateColumns(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustmentdetail) Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentDetailEntityModel) (*model.AdjustmentDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
