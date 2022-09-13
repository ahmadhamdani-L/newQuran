package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type AdjusmentGetRequest struct {
	abstraction.Pagination
	model.AdjusmentFilterModel
}
type AdjusmentGetResponse struct {
	Datas          []model.AdjusmentModel
	PaginationInfo abstraction.PaginationInfo
}
type AdjusmentGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.AdjusmentModel `json:"data"`
	} `json:"body"`
}

// GetByID
type AdjusmentGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjusmentGetByIDResponse struct {
	model.AdjusmentModel
}
type AdjusmentGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data AdjusmentGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type AdjusmentCreateRequest struct {
	model.Adjusment
}
type AdjusmentCreateResponse struct {
	model.AdjusmentModel
}
type AdjusmentCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjusmentCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type AdjusmentUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.Adjusment
}
type AdjusmentUpdateResponse struct {
	model.AdjusmentModel
}
type AdjusmentUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjusmentUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type AdjusmentDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjusmentDeleteResponse struct {
	model.AdjusmentModel
}
type AdjusmentDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AdjusmentDeleteResponse `json:"data"`
	} `json:"body"`
}
