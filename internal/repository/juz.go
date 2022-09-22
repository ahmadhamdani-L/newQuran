package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Juz interface {
	Find(ctx *abstraction.Context, m *model.JuzFilterModel, p *abstraction.Pagination) (*[]model.JuzEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.JuzEntityModel, error)
	Create(ctx *abstraction.Context, e *model.JuzEntity) (*model.JuzEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.JuzEntity) (*model.JuzEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.JuzEntityModel) (*model.JuzEntityModel, error)
}

type juz struct {
	abstraction.Repository
}

func NewJuz(db *gorm.DB) *juz {
	return &juz{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *juz) Find(ctx *abstraction.Context, m *model.JuzFilterModel, p *abstraction.Pagination) (*[]model.JuzEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.JuzEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.JuzEntityModel{})

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

func (r *juz) FindByID(ctx *abstraction.Context, id *int) (*model.JuzEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.JuzEntityModel

	

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
		conn.Preload(clause.Associations).Find(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *juz) Create(ctx *abstraction.Context, e *model.JuzEntity) (*model.JuzEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.JuzEntityModel
	data.JuzEntity = *e
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

func (r *juz) Update(ctx *abstraction.Context, id *int, e *model.JuzEntity) (*model.JuzEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.JuzEntityModel

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	data.JuzEntity = *e
	err = conn.Model(data).UpdateColumns(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *juz) Delete(ctx *abstraction.Context, id *int, e *model.JuzEntityModel) (*model.JuzEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
