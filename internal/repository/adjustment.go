package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Adjustment interface {
	Find(ctx *abstraction.Context, m *model.AdjustmentFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentEntityModel, error)
	Create(ctx *abstraction.Context, e *model.AdjustmentEntityModel) (*model.AdjustmentEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.AdjustmentEntity) (*model.AdjustmentEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentEntityModel) (*model.AdjustmentEntityModel, error)
}

type adjustment struct {
	abstraction.Repository
}

func NewAdjustment(db *gorm.DB) *adjustment {
	return &adjustment{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *adjustment) Find(ctx *abstraction.Context, m *model.AdjustmentFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.AdjustmentEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.AdjustmentEntityModel{})

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

func (r *adjustment) FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentEntityModel

	

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
		conn.Preload(clause.Associations).Find(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustment) Create(ctx *abstraction.Context, e *model.AdjustmentEntityModel) (*model.AdjustmentEntityModel, error) {
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

func (r *adjustment) Update(ctx *abstraction.Context, id *int, e *model.AdjustmentEntity) (*model.AdjustmentEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentEntityModel

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	data.AdjustmentEntity = *e
	err = conn.Model(data).UpdateColumns(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustment) Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentEntityModel) (*model.AdjustmentEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
