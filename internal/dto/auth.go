package dto

import (
	"quran/internal/model"
	res "quran/pkg/util/response"
)

// Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	model.UserEntityModel
}
type AuthLoginResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data AuthLoginResponse `json:"data"`
	} `json:"body"`
}

// Register
type AuthRegisterRequest struct {
	model.UserEntity
}
type AuthRegisterResponse struct {
	model.UserEntityModel
}
type AuthRegisterResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data AuthRegisterResponse `json:"data"`
	} `json:"body"`
}

//forgot passsword
type ForgotPasswordRequest struct {
	Email    string `json:"email" validate:"required,email"`
}

type FindById struct {
	ID    int `json:"id" validate:"required,id"`
}

type ForgotPasswordPayload struct {
	To      string `json:"to" `
	From    string `json:"from" `
	Text    string `json:"text" `
	Subject string `json:"subject" `
}

type NewPassword struct {
	Password string `json:"password"`
}
