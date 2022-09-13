package adjustment

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
	Find(ctx *abstraction.Context, payload *dto.AdjustmentGetRequest) (*dto.AdjustmentGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.AdjustmentGetByIDRequest) (*dto.AdjustmentGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.AdjustmentCreateRequest) (*dto.AdjustmentCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.AdjustmentUpdateRequest) (*dto.AdjustmentUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.AdjustmentDeleteRequest) (*dto.AdjustmentDeleteResponse, error)
}

type service struct {
	Repository repository.Adjustment
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AdjustmentRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.AdjustmentGetRequest) (*dto.AdjustmentGetResponse, error) {
	var result *dto.AdjustmentGetResponse
	var datas *[]model.AdjustmentModel

	datas, info, err := s.Repository.Find(ctx, &payload.AdjustmentFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjustmentGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.AdjustmentGetByIDRequest) (*dto.AdjustmentGetByIDResponse, error) {
	var result *dto.AdjustmentGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AdjustmentGetByIDResponse{
		AdjustmentModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.AdjustmentCreateRequest) (*dto.AdjustmentCreateResponse, error) {
	var result *dto.AdjustmentCreateResponse
	var data *model.AdjustmentModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// data.Context = ctx

		// data.Adjustment = payload.Adjustment
		data, err = s.Repository.Create(ctx, &payload.Adjustment)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.AdjustmentCreateResponse{
		AdjustmentModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.AdjustmentUpdateRequest) (*dto.AdjustmentUpdateResponse, error) {
	var result *dto.AdjustmentUpdateResponse
	var data *model.AdjustmentModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		data, err = s.Repository.Update(ctx, &payload.ID, &payload.Adjustment)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AdjustmentUpdateResponse{
		AdjustmentModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.AdjustmentDeleteRequest) (*dto.AdjustmentDeleteResponse, error) {
	var result *dto.AdjustmentDeleteResponse
	var data *model.AdjustmentModel

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

	result = &dto.AdjustmentDeleteResponse{
		AdjustmentModel: *data,
	}

	return result, nil
}