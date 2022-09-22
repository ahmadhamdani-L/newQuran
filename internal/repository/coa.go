package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
)

type Coa interface {
	Find(ctx *abstraction.Context, m *model.CoaFilterModel, p *abstraction.Pagination) (*[]model.CoaEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.CoaEntityModel, error)
	Create(ctx *abstraction.Context, e *model.CoaEntityModel) (*model.CoaEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.CoaEntityModel) (*model.CoaEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.CoaEntityModel) (*model.CoaEntityModel, error)
	FindWithCode(ctx *abstraction.Context, code *string) (*[]model.CoaEntityModel, error)
}

type coa struct {
	abstraction.Repository
}

func NewCoa(db *gorm.DB) *coa {
	return &coa{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *coa) Find(ctx *abstraction.Context, m *model.CoaFilterModel, p *abstraction.Pagination) (*[]model.CoaEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.CoaEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.CoaEntityModel{})
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

	if err := query.Preload("CoaGroup").Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
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

func (r *coa) FindByID(ctx *abstraction.Context, id *int) (*model.CoaEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.CoaEntityModel
	if err := conn.Where("id = ?", &id).Preload("CoaGroup").First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *coa) FindWithCode(ctx *abstraction.Context, code *string) (*[]model.CoaEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data []model.CoaEntityModel
	tmp := fmt.Sprintf("%s", *code)
	if err := conn.Where("code LIKE ?", tmp+"%").Find(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *coa) Create(ctx *abstraction.Context, e *model.CoaEntityModel) (*model.CoaEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Preload("CoaGroup").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (r *coa) Update(ctx *abstraction.Context, id *int, e *model.CoaEntityModel) (*model.CoaEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Where("id = ?", &id).Preload("CoaGroup").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil

}

func (r *coa) Delete(ctx *abstraction.Context, id *int, e *model.CoaEntityModel) (*model.CoaEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Where("id =?", id).Delete(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}
