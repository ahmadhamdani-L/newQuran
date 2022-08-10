package surah

import (
	"errors"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/internal/repository"
	res "quran/pkg/util/response"
	"quran/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.SurahGetRequest) (*dto.SurahGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.SurahGetByIDRequest) (*dto.SurahGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.SurahCreateRequest) (*dto.SurahCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.SurahUpdateRequest) (*dto.SurahUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.SurahDeleteRequest) (*dto.SurahDeleteResponse, error)
}

type service struct {
	Repository repository.Surah
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.SurahRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.SurahGetRequest) (*dto.SurahGetResponse, error) {
	var result *dto.SurahGetResponse
	var datas *[]model.SurahEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.SurahFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SurahGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.SurahGetByIDRequest) (*dto.SurahGetByIDResponse, error) {
	var result *dto.SurahGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SurahGetByIDResponse{
		SurahEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.SurahCreateRequest) (*dto.SurahCreateResponse, error) {
	var result *dto.SurahCreateResponse
	var data *model.SurahEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// data.Context = ctx

		// data.SurahEntity = payload.SurahEntity
		data, err = s.Repository.Create(ctx, &payload.SurahEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.SurahCreateResponse{
		SurahEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.SurahUpdateRequest) (*dto.SurahUpdateResponse, error) {
	var result *dto.SurahUpdateResponse
	var data *model.SurahEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		// data.Context = ctx
		// data.SurahEntity = payload.SurahEntity
		data, err = s.Repository.Update(ctx, &payload.ID, &payload.SurahEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SurahUpdateResponse{
		SurahEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.SurahDeleteRequest) (*dto.SurahDeleteResponse, error) {
	var result *dto.SurahDeleteResponse
	var data *model.SurahEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data, err = s.Repository.Delete(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SurahDeleteResponse{
		SurahEntityModel: *data,
	}

	return result, nil
}
