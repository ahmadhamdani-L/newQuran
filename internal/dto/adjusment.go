package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type AdjustmentGetRequest struct {
	abstraction.Pagination
	model.AdjustmentFilterModel
}
type AdjustmentGetResponse struct {
	Datas          []model.AdjustmentModel
	PaginationInfo abstraction.PaginationInfo
}
type AdjustmentGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.AdjustmentModel `json:"data"`
	} `json:"body"`
}

// GetByID
type AdjustmentGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjustmentGetByIDResponse struct {
	model.AdjustmentModel
}
type AdjustmentGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data AdjustmentGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type AdjustmentCreateRequest struct {
	model.Adjustment
}
type AdjustmentCreateResponse struct {
	model.AdjustmentModel
}
type AdjustmentCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjustmentCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type AdjustmentUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.Adjustment
}
type AdjustmentUpdateResponse struct {
	model.AdjustmentModel
}
type AdjustmentUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjustmentUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type AdjustmentDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjustmentDeleteResponse struct {
	model.AdjustmentModel
}
type AdjustmentDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjustmentDeleteResponse `json:"data"`
	} `json:"body"`
}
