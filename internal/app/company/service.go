package company

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
	Repository repository.Company
	Db         *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.CompanyGetRequest) (*dto.CompanyGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.CompanyGetByIDRequest) (*dto.CompanyGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CompanyCreateRequest) (*dto.CompanyCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.CompanyUpdateRequest) (*dto.CompanyUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.CompanyDeleteRequest) (*dto.CompanyDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.CompanyRepository
	db := f.Db
	return &service{
		Repository: repository,
		Db:         db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.CompanyGetRequest) (*dto.CompanyGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.CompanyFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.CompanyGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CompanyGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.CompanyGetByIDRequest) (*dto.CompanyGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.CompanyGetByIDResponse{}, err
	}
	result := &dto.CompanyGetByIDResponse{
		CompanyEntityModel: *data,
	}
	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CompanyCreateRequest) (*dto.CompanyCreateResponse, error) {
	var data model.CompanyEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.CompanyEntity = payload.CompanyEntity
		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CompanyCreateResponse{}, err
	}
	result := &dto.CompanyCreateResponse{
		CompanyEntityModel: data,
	}
	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.CompanyUpdateRequest) (*dto.CompanyUpdateResponse, error) {
	var data model.CompanyEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.CompanyEntity = payload.CompanyEntity
		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CompanyUpdateResponse{}, err
	}
	result := &dto.CompanyUpdateResponse{
		CompanyEntityModel: data,
	}
	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.CompanyDeleteRequest) (*dto.CompanyDeleteResponse, error) {
	var data model.CompanyEntityModel
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
		return &dto.CompanyDeleteResponse{}, err
	}
	result := &dto.CompanyDeleteResponse{
		CompanyEntityModel: data,
	}
	return result, nil
}
