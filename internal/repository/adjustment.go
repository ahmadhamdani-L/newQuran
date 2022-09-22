package repository

import (
	"fmt"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Adjustment interface {
	Find(ctx *abstraction.Context, m *model.AdjustmentFilterModel, p *abstraction.Pagination) (*[]model.AdjustmentEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.AdjustmentEntityModel, error)
	Create(ctx *abstraction.Context, e *model.AdjustmentEntityModel) (*model.AdjustmentEntityModel, error)
	CreateWithDetail(ctx *abstraction.Context, e *model.AdjustmentEntityModel, ed *model.AdjustmentDetailEntityModel) (*model.AdjustmentEntityModel, error)
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

func (r *adjustment) CreateWithDetail(ctx *abstraction.Context, e *model.AdjustmentEntityModel, ed *model.AdjustmentDetailEntityModel) (*model.AdjustmentEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data []model.TrialBalanceEntityModel
	var data1 model.CoaEntityModel
	
	nolimit := 1000
		criteriaTB := dto.AdjustmentGetRequest{}
		criteriaTB.Pagination.PageSize = &nolimit


	getAdjustment, _, err := r.Find(ctx , &criteriaTB.AdjustmentFilterModel, &criteriaTB.Pagination)
		if err != nil {
			return nil, err
		}

	versions := len(*getAdjustment) + 1

	e.Versions = versions

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	err = conn.Where("versions = ?", e.VersionsTb).First(&data).
		WithContext(ctx.Request().Context()).Error
		conn.Preload("Company").Preload("TrialBalanceDetail").First(&data).WithContext(ctx.Request().Context())
	if err != nil {
		return nil, err
	}

	ed.CoaCode = data[0].TrialBalanceDetail[0].TrialBalanceDetailEntity.Code
	ed.AdjustmentId = e.ID
	err = conn.Where("code = ?", ed.CoaCode).First(&data1).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	ed.Description = data1.Name

	if err := conn.Create(ed).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(ed).First(ed).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
var data3 model.TrialBalanceDetailEntityModel
	err = conn.Where("trial_balance_id = ?", data[0].Entity.ID).First(&data3).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
data3.AmountBeforeAje = data3.AmountAfterAje
	data3.AmountAjeCr= ed.BalanceSheetCr
	data3.AmountAjeDr = ed.BalanceSheetDr
	data3.AmountAfterAje = data3.AmountAjeCr + data3.AmountAjeDr + data3.AmountBeforeAje
	err = conn.Model(data3).UpdateColumns(&data3).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

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
