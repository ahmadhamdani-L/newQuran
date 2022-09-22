package adjustmentdetail

import (
	"errors"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/internal/repository"
	"quran/pkg/util/response"
	res "quran/pkg/util/response"
	"quran/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.AdjustmentDetailGetRequest) (*dto.AdjustmentDetailGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.AdjustmentDetailGetByIDRequest) (*dto.AdjustmentDetailGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.AdjustmentDetailCreateRequest) (*dto.AdjustmentDetailCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.AdjustmentDetailUpdateRequest) (*dto.AdjustmentDetailUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.AdjustmentDetailDeleteRequest) (*dto.AdjustmentDetailDeleteResponse, error)
}

type service struct {
	Repository repository.AdjustmentDetail
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AdjustmentDetailRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.AdjustmentDetailGetRequest) (*dto.AdjustmentDetailGetResponse, error) {
	var result *dto.AdjustmentDetailGetResponse
	var datas *[]model.AdjustmentDetailEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.AdjustmentDetailFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjustmentDetailGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.AdjustmentDetailGetByIDRequest) (*dto.AdjustmentDetailGetByIDResponse, error) {
	var result *dto.AdjustmentDetailGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjustmentDetailGetByIDResponse{
		AdjustmentDetailEntityModel: *data,
	}

	return result, nil
}


func (s *service) Create(ctx *abstraction.Context, payload *dto.AdjustmentDetailCreateRequest) (*dto.AdjustmentDetailCreateResponse, error) {
	var data model.AdjustmentDetailEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.AdjustmentDetailEntity = payload.AdjustmentDetailEntity
		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.AdjustmentDetailCreateResponse{}, err
	}
	result := &dto.AdjustmentDetailCreateResponse{
		AdjustmentDetailEntityModel: data,
	}
	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.AdjustmentDetailUpdateRequest) (*dto.AdjustmentDetailUpdateResponse, error) {
	var result *dto.AdjustmentDetailUpdateResponse
	var data *model.AdjustmentDetailEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		data, err = s.Repository.Update(ctx, &payload.ID, &payload.AdjustmentDetailEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AdjustmentDetailUpdateResponse{
		AdjustmentDetailEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.AdjustmentDetailDeleteRequest) (*dto.AdjustmentDetailDeleteResponse, error) {
	var result *dto.AdjustmentDetailDeleteResponse
	var data *model.AdjustmentDetailEntityModel

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

	result = &dto.AdjustmentDetailDeleteResponse{
		AdjustmentDetailEntityModel: *data,
	}

	return result, nil
}
