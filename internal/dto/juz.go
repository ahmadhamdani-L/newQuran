package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type JuzGetRequest struct {
	abstraction.Pagination
	model.JuzFilterModel
}
type JuzGetResponse struct {
	Datas          []model.JuzEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type JuzGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.JuzEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type JuzGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type JuzGetByIDResponse struct {
	model.JuzEntityModel
}
type JuzGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data JuzGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type JuzCreateRequest struct {
	model.JuzEntity
}
type JuzCreateResponse struct {
	model.JuzEntityModel
}
type JuzCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data JuzCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type JuzUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.JuzEntity
}
type JuzUpdateResponse struct {
	model.JuzEntityModel
}
type JuzUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data JuzUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type JuzDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type JuzDeleteResponse struct {
	model.JuzEntityModel
}
type JuzDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data JuzDeleteResponse `json:"data"`
	} `json:"body"`
}
