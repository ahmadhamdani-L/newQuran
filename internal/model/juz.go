package model

import (
	"quran/internal/abstraction"
	"quran/pkg/util/date"
	"gorm.io/gorm"

)

type JuzEntity struct {

	NamaJuz   string `json:"nama_juz" validate:"required" gorm:"index:idx_juz_nama_juz,unique"`
	NoJuz string `json:"no_juz" validate:"required"`
}

type JuzFilter struct {
	NamaJuz   *string `query:"nama_juz" filter:"ILIKE"`
	NoJuz *string `query:"no_juz"`
}

type JuzEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	JuzEntity

	//relations
	Surahs []SurahEntityModel `json:"surahs" gorm:"foreignKey:JuzId"`
	
	

	// relations
	// SampleChilds []SampleChildEntityModel `json:"sample_childs" gorm:"foreignKey:SampleId"`


	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JuzFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	JuzFilter
}

func (JuzEntityModel) TableName() string {
	return "juzs"
}

func (m *JuzEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	// m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *JuzEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = *date.DateTodayLocal()
	// m.ModifiedBy = m.Context.Auth.Name
	return
}
