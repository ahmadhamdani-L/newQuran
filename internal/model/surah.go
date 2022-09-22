package model

import (
	"quran/internal/abstraction"
	"quran/pkg/util/date"
	"gorm.io/gorm"
)

type SurahEntity struct {
	NamaSurah string `json:"nama_surah" validate:"required" gorm:"index:idx_surah_nama_surah,unique"`
	NoSurah   string `json:"no_surah" validate:"required"`
	JuzId int `json:"juz_id" gorm:"not null"`
}

type SurahFilter struct {
	NamaSurah *string `query:"nama_surah" filter:"ILIKE"`
	NoSurah   *string `query:"no_surah"`
}

type SurahEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	SurahEntity

	// relations
	JuzId int `json:"juz_id"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SurahFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	SurahFilter
}

func (SurahEntityModel) TableName() string {
	return "surahs"
}

func (m *SurahEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *SurahEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
