package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
)

type Company interface {
	Find(ctx *abstraction.Context, m *model.CompanyFilterModel, p *abstraction.Pagination) (*[]model.CompanyEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.CompanyEntityModel, error)
	Create(ctx *abstraction.Context, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error)
}

type company struct {
	abstraction.Repository
}

func NewCompany(db *gorm.DB) *company {
	return &company{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *company) Find(ctx *abstraction.Context, m *model.CompanyFilterModel, p *abstraction.Pagination) (*[]model.CompanyEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.CompanyEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.CompanyEntityModel{})

	//filter
	query = r.Filter(ctx, query, *m)

	//sort
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

	//pagination
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
	offset := limit * (*p.Page - 1)
	query = query.Limit(limit).Offset(offset)

	if err := query.Preload("ParentCompany").Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
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

func (r *company) FindByID(ctx *abstraction.Context, id *int) (*model.CompanyEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.CompanyEntityModel
	if err := conn.Where("id = ?", &id).First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *company) Create(ctx *abstraction.Context, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (r *company) Update(ctx *abstraction.Context, id *int, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Where("id = ?", &id).Preload("ParentCompany").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil

}

func (r *company) Delete(ctx *abstraction.Context, id *int, e *model.CompanyEntityModel) (*model.CompanyEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Where("id =?", id).Delete(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}
