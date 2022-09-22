package factory

import (
	"quran/database"
	"quran/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db             *gorm.DB
	UserRepository repository.User
	JuzRepository  repository.Juz
	SurahRepository  repository.Surah
	AdjustmentRepository repository.Adjustment
	AdjustmentDetailRepository repository.AdjustmentDetail
	CoaRepository                              repository.Coa
	CoaGroupRepository                         repository.CoaGroup
	CompanyRepository                          repository.Company
	TrialBalanceRepository                     repository.TrialBalance
	TrialBalanceDetailRepository               repository.TrialBalanceDetail
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("QURAN")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
	f.JuzRepository = repository.NewJuz(f.Db)
	f.SurahRepository = repository.NewSurah(f.Db)
	f.AdjustmentRepository = repository.NewAdjustment(f.Db)
	f.AdjustmentDetailRepository = repository.NewAdjustmentDetail(f.Db)
	f.CoaRepository = repository.NewCoa(f.Db)
	f.CoaGroupRepository = repository.NewCoaGroup(f.Db)
	f.CompanyRepository = repository.NewCompany(f.Db)
	f.TrialBalanceRepository = repository.NewTrialBalance(f.Db)
	f.TrialBalanceDetailRepository = repository.NewTrialBalanceDetail(f.Db)
}
