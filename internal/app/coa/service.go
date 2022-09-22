package coa

import (
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/internal/repository"
	"quran/pkg/util/response"
	"quran/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type service struct {
	Repository repository.Coa
	Db         *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.CoaGetRequest) (*dto.CoaGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.CoaGetByIDRequest) (*dto.CoaGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CoaCreateRequest) (*dto.CoaCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.CoaUpdateRequest) (*dto.CoaUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.CoaDeleteRequest) (*dto.CoaDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.CoaRepository
	db := f.Db
	return &service{
		Repository: repository,
		Db:         db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.CoaGetRequest) (*dto.CoaGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.CoaFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.CoaGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.CoaGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.CoaGetByIDRequest) (*dto.CoaGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.CoaGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CoaGetByIDResponse{
		CoaEntityModel: *data,
	}
	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CoaCreateRequest) (*dto.CoaCreateResponse, error) {
	var data model.CoaEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.CoaEntity = payload.CoaEntity

		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CoaCreateResponse{}, err
	}

	result := &dto.CoaCreateResponse{
		CoaEntityModel: data,
	}
	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.CoaUpdateRequest) (*dto.CoaUpdateResponse, error) {
	var data model.CoaEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.CoaEntity = payload.CoaEntity

		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CoaUpdateResponse{}, err
	}
	result := &dto.CoaUpdateResponse{
		CoaEntityModel: data,
	}
	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.CoaDeleteRequest) (*dto.CoaDeleteResponse, error) {
	var data model.CoaEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		result, err := s.Repository.Delete(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CoaDeleteResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CoaDeleteResponse{
		CoaEntityModel: data,
	}
	return result, nil
}
