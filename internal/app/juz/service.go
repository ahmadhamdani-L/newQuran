package juz

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
	Find(ctx *abstraction.Context, payload *dto.JuzGetRequest) (*dto.JuzGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.JuzGetByIDRequest) (*dto.JuzGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.JuzCreateRequest) (*dto.JuzCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.JuzUpdateRequest) (*dto.JuzUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.JuzDeleteRequest) (*dto.JuzDeleteResponse, error)
}

type service struct {
	Repository repository.Juz
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.JuzRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.JuzGetRequest) (*dto.JuzGetResponse, error) {
	var result *dto.JuzGetResponse
	var datas *[]model.JuzEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.JuzFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.JuzGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.JuzGetByIDRequest) (*dto.JuzGetByIDResponse, error) {
	var result *dto.JuzGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.JuzGetByIDResponse{
		JuzEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.JuzCreateRequest) (*dto.JuzCreateResponse, error) {
	var result *dto.JuzCreateResponse
	var data *model.JuzEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// data.Context = ctx

		// data.JuzEntity = payload.JuzEntity
		data, err = s.Repository.Create(ctx, &payload.JuzEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.JuzCreateResponse{
		JuzEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.JuzUpdateRequest) (*dto.JuzUpdateResponse, error) {
	var result *dto.JuzUpdateResponse
	var data *model.JuzEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		data, err = s.Repository.Update(ctx, &payload.ID, &payload.JuzEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.JuzUpdateResponse{
		JuzEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.JuzDeleteRequest) (*dto.JuzDeleteResponse, error) {
	var result *dto.JuzDeleteResponse
	var data *model.JuzEntityModel

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

	result = &dto.JuzDeleteResponse{
		JuzEntityModel: *data,
	}

	return result, nil
}
