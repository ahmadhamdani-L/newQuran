package model

import (
	"quran/internal/abstraction"
	"quran/pkg/util/date"

	"gorm.io/gorm"
)

type TrialBalanceEntity struct {
	Period      string `json:"period" validate:"required"`
	Versions    int    `json:"versions" validate:"required"`
	CompanyID   int    `json:"company_id" validate:"required"`
	FormatterID int    `json:"formatter_id" validate:"required"`
}

type TrialBalanceFilter struct {
	Period      *string `query:"period" filter:"LIKE"`
	Versions    *int    `query:"versions"`
	CompanyID   *int    `query:"company_id"`
	FormatterID *int    `query:"formatter_id"`
}

type TrialBalanceEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	TrialBalanceEntity

	// relations
	Company            CompanyEntityModel              `json:"company" gorm:"foreignKey:CompanyID"`
	// Formatter          FormatterEntityModel            `json:"formatter,omitempty" gorm:"foreignKey:FormatterID"`
	TrialBalanceDetail []TrialBalanceDetailEntityModel `json:"trial_balance_detail" gorm:"foreignKey:TrialBalanceID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type TrialBalanceFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	TrialBalanceFilter
}

func (TrialBalanceEntityModel) TableName() string {
	return "trial_balance"
}

func (m *TrialBalanceEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *TrialBalanceEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
