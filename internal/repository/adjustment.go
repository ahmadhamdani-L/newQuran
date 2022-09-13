package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Adjustment interface {
	Find(ctx *abstraction.Context, m *model.AdjustmentFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentModel, error)
	Create(ctx *abstraction.Context, e *model.Adjustment) (*model.AdjustmentModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.Adjustment) (*model.AdjustmentModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentModel) (*model.AdjustmentModel, error)
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

func (r *adjustment) Find(ctx *abstraction.Context, m *model.AdjustmentFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.AdjustmentModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.AdjustmentModel{})

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

func (r *adjustment) FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentModel

	

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
		conn.Preload(clause.Associations).Find(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustment) Create(ctx *abstraction.Context, e *model.Adjustment) (*model.AdjustmentModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentModel
	data.Adjustment = *e
	var Nama = ctx.Auth.Name
	data.CreatedBy = Nama
	err := conn.Create(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Model(data).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustment) Update(ctx *abstraction.Context, id *int, e *model.Adjustment) (*model.AdjustmentModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AdjustmentModel

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	data.Adjustment = *e
	var Nama = ctx.Auth.Name
	data.ModifiedBy = Nama
	err = conn.Model(data).UpdateColumns(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *adjustment) Delete(ctx *abstraction.Context, id *int, e *model.AdjustmentModel) (*model.AdjustmentModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
