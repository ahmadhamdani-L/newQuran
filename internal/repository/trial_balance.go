package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
)

type TrialBalance interface {
	Find(ctx *abstraction.Context, m *model.TrialBalanceFilterModel, p *abstraction.Pagination) (*[]model.TrialBalanceEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.TrialBalanceEntityModel, error)
	Create(ctx *abstraction.Context, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error)
}

type trialbalance struct {
	abstraction.Repository
}

func NewTrialBalance(db *gorm.DB) *trialbalance {
	return &trialbalance{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *trialbalance) Find(ctx *abstraction.Context, m *model.TrialBalanceFilterModel, p *abstraction.Pagination) (*[]model.TrialBalanceEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.TrialBalanceEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.TrialBalanceEntityModel{})
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

	if err := query.Preload("Company").Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
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

func (r *trialbalance) FindByID(ctx *abstraction.Context, id *int) (*model.TrialBalanceEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.TrialBalanceEntityModel
	if err := conn.Where("id = ?", &id).Preload("Company").Preload("TrialBalanceDetail").First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *trialbalance) Create(ctx *abstraction.Context, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Preload("Company").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (r *trialbalance) Update(ctx *abstraction.Context, id *int, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).Preload("Company").WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Where("id = ?", &id).Preload("Company").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil

}

func (r *trialbalance) Delete(ctx *abstraction.Context, id *int, e *model.TrialBalanceEntityModel) (*model.TrialBalanceEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Where("id =?", id).Delete(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}
