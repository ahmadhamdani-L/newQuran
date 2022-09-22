package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type CoaGetRequest struct {
	abstraction.Pagination
	model.CoaFilterModel
}
type CoaGetResponse struct {
	Datas          []model.CoaEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type CoaGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.CoaEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type CoaGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CoaGetByIDResponse struct {
	model.CoaEntityModel
}
type CoaGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data CoaGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type CoaCreateRequest struct {
	model.CoaEntity
}
type CoaCreateResponse struct {
	model.CoaEntityModel
}
type CoaCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data CoaCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type CoaUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.CoaEntity
}
type CoaUpdateResponse struct {
	model.CoaEntityModel
}
type CoaUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data CoaUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type CoaDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CoaDeleteResponse struct {
	model.CoaEntityModel
}
type CoaDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data CoaDeleteResponse `json:"data"`
	} `json:"body"`
}

// export
type CoaExportRequest struct {
	model.CoaFilterModel
}
type CoaExportResponse struct {
	File string `json:"file"`
}
type CoaExportResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data CoaExportResponse `json:"data"`
	} `json:"body"`
}

// Import
type CoaImportRequest struct {
	Datas []model.CoaEntity
}
type CoaImportResponse struct {
	Datas          []model.CoaEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type CoaImportResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.CoaEntityModel `json:"data"`
	} `json:"body"`
}
