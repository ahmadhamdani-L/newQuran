package model

import (
	"quran/internal/abstraction"
)

type TrialBalanceDetailEntity struct {
	Code            string  `json:"code" validate:"required"`
	AmountBeforeAje float64 `json:"amount_before_aje" validate:"required"`
	AmountAjeDr     float64 `json:"amount_aje_dr" validate:"required"`
	AmountAjeCr     float64 `json:"amount_aje_cr" validate:"required"`
	AmountAfterAje  float64 `json:"amount_after_aje" validate:"required"`
	ReffAjeDr       string  `json:"reff_aje_dr" validate:"required"`
	ReffAjeCr       string  `json:"reff_aje_cr" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	TrialBalanceID  int     `json:"trial_balance_id" validate:"required"`
}

type TrialBalanceDetailFilter struct {
	Code            *string  `query:"code" filter:"LIKE"`
	AmountBeforeAje *float64 `query:"amount_before_aje"`
	AmountAjeDr     *float64 `query:"amount_aje_dr"`
	AmountAjeCr     *float64 `query:"amount_aje_cr"`
	AmountAfterAje  *float64 `query:"amount_after_aje"`
	ReffAjeDr       *string  `query:"reff_aje_dr" filter:"ILIKE"`
	ReffAjeCr       *string  `query:"reff_aje_cr" filter:"ILIKE"`
	Description     *string  `query:"description" filter:"ILIKE"`
	TrialBalanceID  *int     `query:"trial_balance_id" validate:"required"`
}

type TrialBalanceDetailEntityModel struct {
	// abstraction
	// abstraction.Entity
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	// entity
	TrialBalanceDetailEntity

	// relations
	// SampleChilds []SampleChildEntityModel `json:"sample_childs" gorm:"foreignKey:SampleId"`
	TrialBalance TrialBalanceEntityModel `json:"trial_balance" gorm:"foreignKey:TrialBalanceID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type TrialBalanceDetailFilterModel struct {
	// abstraction
	// abstraction.Filter

	// filter
	TrialBalanceDetailFilter
}

func (TrialBalanceDetailEntityModel) TableName() string {
	return "trial_balance_detail"
}

// func (m *TrialBalanceDetailEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	m.CreatedAt = *date.DateTodayLocal()
// 	m.CreatedBy = m.Context.Auth.Name
// 	return
// }

// func (m *TrialBalanceDetailEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
// 	m.ModifiedAt = date.DateTodayLocal()
// 	m.ModifiedBy = &m.Context.Auth.Name
// 	return
// }
