package model

import (
	"quran/internal/abstraction"
	"quran/pkg/util/date"
	"gorm.io/gorm"
)

type AdjustmentEntity struct {
	TrxNumber   string `json:"trx_number" validate:"required" gorm:"index:idx_trx_number_adjustment,unique"`
	Note string `json:"note" validate:"required"`
	CompanyId int `json:"company_id" `
	Period string `json:"period" validate:"required"`
}


type AdjustmentFilter struct {

	TrxNumber   *string `query:"trx_number" filter:"ILIKE"`
}

type  AdjustmentEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	AdjustmentEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type AdjustmentFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	AdjustmentFilter
}

func (AdjustmentEntityModel) TableName() string {
	return "adjustment"
}

func (m *AdjustmentEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *AdjustmentEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
