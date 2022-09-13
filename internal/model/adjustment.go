package model

import (
	"quran/internal/abstraction"
	"quran/pkg/util/date"
	"time"

	"gorm.io/gorm"
)

type Adjustment struct {
	TrxNumber   string `json:"trx_number" validate:"required" gorm:"index:idx_trx_number_adjustment,unique"`
	Note string `json:"note" validate:"required"`
	CompanyId int `json:"company_id" `
	Period time.Time `json:"period" validate:"required"`
}


type AdjustmentFilter struct {

	TrxNumber   *string `query:"trx_number" filter:"ILIKE"`
}

type  AdjustmentModel struct {
	// abstraction
	abstraction.Entity

	// entity
	Adjustment

	//relations
	// Surahs []SurahEntityModel `json:"surahs" gorm:"foreignKey:JuzId"`
	
	

	// relations
	// SampleChilds []SampleChildEntityModel `json:"sample_childs" gorm:"foreignKey:SampleId"`


	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type AdjustmentFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	AdjustmentFilter
}

func (AdjustmentModel) TableName() string {
	return "adjustment"
}

func (m *AdjustmentModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	// m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *AdjustmentModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = *date.DateTodayLocal()
	// m.ModifiedBy = m.Context.Auth.Name
	return
}
