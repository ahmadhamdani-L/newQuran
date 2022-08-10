package dto

import (
	"quran/internal/abstraction"
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Get
type SurahGetRequest struct {
	abstraction.Pagination
	model.SurahFilterModel
}
type SurahGetResponse struct {
	Datas          []model.SurahEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type SurahGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.SurahEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type SurahGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SurahGetByIDResponse struct {
	model.SurahEntityModel
}
type SurahGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data SurahGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type SurahCreateRequest struct {
	model.SurahEntity
	// juz_id string `json:"juz_id"`
}
type SurahCreateResponse struct {
	model.SurahEntityModel
}
type SurahCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data SurahCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type SurahUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.SurahEntity
}
type SurahUpdateResponse struct {
	model.SurahEntityModel
}
type SurahUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data SurahUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type SurahDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SurahDeleteResponse struct {
	model.SurahEntityModel
}
type SurahDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data SurahDeleteResponse `json:"data"`
	} `json:"body"`
}
