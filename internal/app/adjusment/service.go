package adjusment

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
	Find(ctx *abstraction.Context, payload *dto.AdjusmentGetRequest) (*dto.AdjusmentGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.AdjusmentGetByIDRequest) (*dto.AdjusmentGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.AdjusmentCreateRequest) (*dto.AdjusmentCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.AdjusmentUpdateRequest) (*dto.AdjusmentUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.AdjusmentDeleteRequest) (*dto.AdjusmentDeleteResponse, error)
}

type service struct {
	Repository repository.Adjusment
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AdjusmentRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.AdjusmentGetRequest) (*dto.AdjusmentGetResponse, error) {
	var result *dto.AdjusmentGetResponse
	var datas *[]model.AdjusmentEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.AdjusmentFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjusmentGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.AdjusmentGetByIDRequest) (*dto.AdjusmentGetByIDResponse, error) {
	var result *dto.AdjusmentGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjusmentGetByIDResponse{
		AdjusmentEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.AdjusmentCreateRequest) (*dto.AdjusmentCreateResponse, error) {
	var result *dto.AdjusmentCreateResponse
	var data *model.AdjusmentEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// data.Context = ctx

		// data.AdjusmentEntity = payload.AdjusmentEntity
		data, err = s.Repository.Create(ctx, &payload.AdjusmentEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.AdjusmentCreateResponse{
		AdjusmentEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.AdjusmentUpdateRequest) (*dto.AdjusmentUpdateResponse, error) {
	var result *dto.AdjusmentUpdateResponse
	var data *model.AdjusmentEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		data, err = s.Repository.Update(ctx, &payload.ID, &payload.AdjusmentEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AdjusmentUpdateResponse{
		AdjusmentEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.AdjusmentDeleteRequest) (*dto.AdjusmentDeleteResponse, error) {
	var result *dto.AdjusmentDeleteResponse
	var data *model.AdjusmentEntityModel

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

	result = &dto.AdjusmentDeleteResponse{
		AdjusmentEntityModel: *data,
	}

	return result, nil
}
