package trialbalancedetail

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
	Repository repository.TrialBalanceDetail
	Db         *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.TrialBalanceDetailGetRequest) (*dto.TrialBalanceDetailGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.TrialBalanceDetailGetByIDRequest) (*dto.TrialBalanceDetailGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.TrialBalanceDetailCreateRequest) (*dto.TrialBalanceDetailCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.TrialBalanceDetailUpdateRequest) (*dto.TrialBalanceDetailUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.TrialBalanceDetailDeleteRequest) (*dto.TrialBalanceDetailDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.TrialBalanceDetailRepository
	db := f.Db
	return &service{
		Repository: repository,
		Db:         db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.TrialBalanceDetailGetRequest) (*dto.TrialBalanceDetailGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.TrialBalanceDetailFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.TrialBalanceDetailGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	result := &dto.TrialBalanceDetailGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.TrialBalanceDetailGetByIDRequest) (*dto.TrialBalanceDetailGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.TrialBalanceDetailGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.TrialBalanceDetailGetByIDResponse{
		TrialBalanceDetailEntityModel: *data,
	}
	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.TrialBalanceDetailCreateRequest) (*dto.TrialBalanceDetailCreateResponse, error) {
	var data model.TrialBalanceDetailEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.TrialBalanceDetailEntity = payload.TrialBalanceDetailEntity

		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.TrialBalanceDetailCreateResponse{}, err
	}

	result := &dto.TrialBalanceDetailCreateResponse{
		TrialBalanceDetailEntityModel: data,
	}
	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.TrialBalanceDetailUpdateRequest) (*dto.TrialBalanceDetailUpdateResponse, error) {
	var data model.TrialBalanceDetailEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.TrialBalanceDetailEntity = payload.TrialBalanceDetailEntity

		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.TrialBalanceDetailUpdateResponse{}, err
	}
	result := &dto.TrialBalanceDetailUpdateResponse{
		TrialBalanceDetailEntityModel: data,
	}
	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.TrialBalanceDetailDeleteRequest) (*dto.TrialBalanceDetailDeleteResponse, error) {
	var data model.TrialBalanceDetailEntityModel
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
		return &dto.TrialBalanceDetailDeleteResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.TrialBalanceDetailDeleteResponse{
		TrialBalanceDetailEntityModel: data,
	}
	return result, nil
}
