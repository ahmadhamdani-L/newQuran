package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type Surah interface {
	Find(ctx *abstraction.Context, m *model.SurahFilterModel, p *abstraction.Pagination) (*[]model.SurahEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.SurahEntityModel, error)
	Create(ctx *abstraction.Context, e *model.SurahEntity) (*model.SurahEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.SurahEntity) (*model.SurahEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.SurahEntityModel) (*model.SurahEntityModel, error)
}

type surah struct {
	abstraction.Repository
}

func NewSurah(db *gorm.DB) *surah {
	return &surah{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *surah) Find(ctx *abstraction.Context, m *model.SurahFilterModel, p *abstraction.Pagination) (*[]model.SurahEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.SurahEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.SurahEntityModel{})

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

func (r *surah) FindByID(ctx *abstraction.Context, id *int) (*model.SurahEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SurahEntityModel
	
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
		// conn.Preload(string(clause.Associations)).Find(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *surah) Create(ctx *abstraction.Context, e *model.SurahEntity) (*model.SurahEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SurahEntityModel
	data.SurahEntity = *e
	var Nama = ctx.Auth.Name
	data.CreatedBy = Nama
	data.JuzId = data.SurahEntity.JuzId
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

func (r *surah) Update(ctx *abstraction.Context, id *int, e *model.SurahEntity) (*model.SurahEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SurahEntityModel
	data.SurahEntity = *e

	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(data).UpdateColumns(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *surah) Delete(ctx *abstraction.Context, id *int, e *model.SurahEntityModel) (*model.SurahEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
		
	if err != nil {
		return nil, err
	}

	return e, nil
}
