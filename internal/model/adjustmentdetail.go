package model

import (
	"quran/internal/abstraction"
)

type AdjustmentDetailEntity struct {
	AdjustmentId int `json:"adjustment_id" `
	CoaCode string `json:"coa_code" `
	ReffNumber string `json:"reff_number" `
	Description string `json:"description" `
	BalanceSheetDr float64 `json:"balance_sheet_dr" sql:"type:decimal(20,8);"`
	BalanceSheetCr float64 `json:"balance_sheet_cr" sql:"type:decimal(20,8);"`
	IncomeStatementDr float64 `json:"income_statement_dr" sql:"type:decimal(20,8);"`
	IncomeStatementCr float64 `json:"income_statement_cr" sql:"type:decimal(20,8);"`
}


type AdjustmentDetailFilter struct {

	TrxNumber   *string `query:"trx_number" filter:"ILIKE"`
}

type  AdjustmentDetailEntityModel struct {
// abstraction
	// abstraction.Entity
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	// entity
	AdjustmentDetailEntity

	// relations
	// SampleChilds []SampleChildEntityModel `json:"sample_childs" gorm:"foreignKey:SampleId"`
	Adjustment AdjustmentEntityModel `json:"adjustment" gorm:"foreignKey:AdjustmentId"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`

}

type AdjustmentDetailFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	AdjustmentFilter
}

func (AdjustmentDetailEntityModel) TableName() string {
	return "adjustment_detail"
}

// func (m *AdjustmentDetailEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.CreatedAt = *date.DateTodayLocal()
// 	m.CreatedBy = m.Context.Auth.Name
// 	return
// }

// func (m *AdjustmentDetailEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
// 	m.ModifiedAt = date.DateTodayLocal()
// 	m.ModifiedBy = &m.Context.Auth.Name
// 	return
// }
