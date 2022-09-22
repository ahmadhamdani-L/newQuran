package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type CompanyGetRequest struct {
	abstraction.Pagination
	model.CompanyFilterModel
}
type CompanyGetResponse struct {
	Datas          []model.CompanyEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type CompanyGetResponseDoc struct {
	Body struct {
		Meta res.Meta                   `json:"meta"`
		Data []model.CompanyEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type CompanyGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CompanyGetByIDResponse struct {
	model.CompanyEntityModel
}
type CompanyGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CompanyGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type CompanyCreateRequest struct {
	model.CompanyEntity
}
type CompanyCreateResponse struct {
	model.CompanyEntityModel
}
type CompanyCreateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data CompanyCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type CompanyUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.CompanyEntity
}
type CompanyUpdateResponse struct {
	model.CompanyEntityModel
}
type CompanyUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data CompanyUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type CompanyDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CompanyDeleteResponse struct {
	model.CompanyEntityModel
}
type CompanyDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data CompanyDeleteResponse `json:"data"`
	} `json:"body"`
}
