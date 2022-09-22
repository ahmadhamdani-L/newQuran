package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type AdjustmentDetailGetRequest struct {
	abstraction.Pagination
	model.AdjustmentDetailFilterModel
}
type AdjustmentDetailGetResponse struct {
	Datas          []model.AdjustmentDetailEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type AdjustmentDetailGetResponseDoc struct {
	Body struct {
		Meta res.Meta                              `json:"meta"`
		Data []model.AdjustmentDetailEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type AdjustmentDetailGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjustmentDetailGetByIDResponse struct {
	model.AdjustmentDetailEntityModel
}
type AdjustmentDetailGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                          `json:"meta"`
		Data AdjustmentDetailGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type AdjustmentDetailCreateRequest struct {
	model.AdjustmentDetailEntity
}
type AdjustmentDetailCreateResponse struct {
	model.AdjustmentDetailEntityModel
}
type AdjustmentDetailCreateResponseDoc struct {
	Body struct {
		Meta res.Meta                         `json:"meta"`
		Data AdjustmentDetailCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type AdjustmentDetailUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.AdjustmentDetailEntity
}
type AdjustmentDetailUpdateResponse struct {
	model.AdjustmentDetailEntityModel
}
type AdjustmentDetailUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta                         `json:"meta"`
		Data AdjustmentDetailUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type AdjustmentDetailDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type AdjustmentDetailDeleteResponse struct {
	model.AdjustmentDetailEntityModel
}
type AdjustmentDetailDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta                         `json:"meta"`
		Data AdjustmentDetailDeleteResponse `json:"data"`
	} `json:"body"`
}

// export
type AdjustmentDetailExportRequest struct {
	model.AdjustmentDetailFilterModel
}
type AdjustmentDetailExportResponse struct {
	File string `json:"file"`
}
type AdjustmentDetailExportResponseDoc struct {
	Body struct {
		Meta res.Meta                         `json:"meta"`
		Data AdjustmentDetailExportResponse `json:"data"`
	} `json:"body"`
}

// Import
type AdjustmentDetailImportRequest struct {
	Datas []model.AdjustmentDetailEntity
}
type AdjustmentDetailImportResponse struct {
	Datas          []model.AdjustmentDetailEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type AdjustmentDetailImportResponseDoc struct {
	Body struct {
		Meta res.Meta                              `json:"meta"`
		Data []model.AdjustmentDetailEntityModel `json:"data"`
	} `json:"body"`
}
