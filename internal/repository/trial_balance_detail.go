package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/model"

	"gorm.io/gorm"
)

type TrialBalanceDetail interface {
	Find(ctx *abstraction.Context, m *model.TrialBalanceDetailFilterModel, p *abstraction.Pagination) (*[]model.TrialBalanceDetailEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.TrialBalanceDetailEntityModel, error)
	Create(ctx *abstraction.Context, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error)
	FindWithCode(ctx *abstraction.Context, code *string) (*[]model.TrialBalanceDetailEntityModel, error)
}

type trialbalancedetail struct {
	abstraction.Repository
}

func NewTrialBalanceDetail(db *gorm.DB) *trialbalancedetail {
	return &trialbalancedetail{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *trialbalancedetail) Find(ctx *abstraction.Context, m *model.TrialBalanceDetailFilterModel, p *abstraction.Pagination) (*[]model.TrialBalanceDetailEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.TrialBalanceDetailEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.TrialBalanceDetailEntityModel{})
	//filter
	query = r.Filter(ctx, query, *m)

	//sort
	if p.Sort == nil {
		sort := "asc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "id"
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

	if err := query.Preload("TrialBalance").Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
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

func (r *trialbalancedetail) FindWithCode(ctx *abstraction.Context, code *string) (*[]model.TrialBalanceDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data []model.TrialBalanceDetailEntityModel
	tmp := fmt.Sprintf("%s", *code)
	if err := conn.Where("code LIKE ?", tmp+"%").Find(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *trialbalancedetail) FindByID(ctx *abstraction.Context, id *int) (*model.TrialBalanceDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.TrialBalanceDetailEntityModel
	if err := conn.Where("id = ?", &id).Preload("TrialBalance").First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *trialbalancedetail) Create(ctx *abstraction.Context, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Preload("TrialBalance").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (r *trialbalancedetail) Update(ctx *abstraction.Context, id *int, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).Preload("TrialBalance").WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Where("id = ?", &id).Preload("TrialBalance").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil

}

func (r *trialbalancedetail) Delete(ctx *abstraction.Context, id *int, e *model.TrialBalanceDetailEntityModel) (*model.TrialBalanceDetailEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Where("id =?", id).Delete(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}
