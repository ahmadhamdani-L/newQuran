package repository

import (
	"errors"
	"log"
	"quran/internal/abstraction"
	"quran/internal/model"
	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	GetUserById(ctx *abstraction.Context, id *int) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntity) (*model.UserEntityModel, error)
	checkTrx(ctx *abstraction.Context) *gorm.DB
	GetTemplateById(ctx *abstraction.Context, id uint, userId int) (model.Template, error)
	GetClientByIdUserId(ctx *abstraction.Context, clientId string, userId int) (model.Client, error)
	GetClientFields(ctx *abstraction.Context, id uint) ([]model.Field, error) 
	SaveAuditLog(ctx *abstraction.Context, audit model.Audit) error
	GetUserByName(ctx *abstraction.Context, name string) (*model.UserEntityModel, error)
	UpdateUserPassword(ctx *abstraction.Context, userId int, hashPwd string) error
	InvalidateAllSessions(ctx *abstraction.Context, userID int)
}


type user struct {
	abstraction.Repository
}



func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}


func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) GetUserById(ctx *abstraction.Context, id *int) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}


func (r *user) Create(ctx *abstraction.Context, e *model.UserEntity) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	data.UserEntity = *e
	err := conn.Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}

func (r *user)GetTemplateById(ctx *abstraction.Context, id uint, userId int) (model.Template, error) {
	conn := r.checkTrx(ctx)
	var template model.Template
	dbc := conn.Where("id = ? AND user_id = ?", id, userId).Find(&template)

	if dbc.Error != nil {
		log.Println("An error occurred: ", dbc.Error.Error())
		return model.Template{}, dbc.Error
	}

	if template.ID == 0 {
		return model.Template{}, errors.New("an error occurred - template details not found")
	}
	return template, nil
}

func (r *user) GetClientByIdUserId(ctx *abstraction.Context, clientId string, userId int) (model.Client, error) {
	conn := r.checkTrx(ctx)
	var client model.Client
	dbc := conn.Where("id = ? AND user_id = ?", clientId, userId).Find(&client)
	if dbc.Error != nil {
		log.Println("An error occurred while fetching client details :: ", dbc.Error.Error())
		return model.Client{}, dbc.Error
	}
	return client, nil
}

func (r *user) GetClientFields(ctx *abstraction.Context, id uint) ([]model.Field, error) {
	conn := r.checkTrx(ctx)
	var fieldDetails []model.Field
	conn.Where("client_id = ?", id).Find(&fieldDetails)
	if len(fieldDetails) == 0 {
		return []model.Field{}, errors.New("an error occurred - client fields not found")
	}
	return fieldDetails, nil
}

func (r *user) SaveAuditLog(ctx *abstraction.Context, audit model.Audit) error {
	conn := r.checkTrx(ctx)
	dbc := conn.Create(&audit)
	if dbc.Error != nil {
		log.Println("An error occurred while saving audit log :: ", dbc.Error.Error())
		return dbc.Error
	}
	return nil
}
func (r *user) GetUserByName(ctx *abstraction.Context, name string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)
	var user *model.UserEntityModel
	dbc := conn.Where("name = ?", name).Find(&user)
	if dbc.Error != nil {
		log.Println("An error occurred :: ", dbc.Error.Error())
		return &model.UserEntityModel{}, dbc.Error
	}
	return user, nil
}

// func (r *user)UpdateUserPassword(ctx *abstraction.Context, userId int, hashPwd string, e *model.UserEntity) error {
// 	conn := r.checkTrx(ctx)
// 	var data model.UserEntityModel

// 	err := conn.Where("id = ?", userId).First(&data).
// 		WithContext(ctx.Request().Context()).Error

// 		data.UserEntity = *e
// 	err = conn.Model(data).Where("id = ?", userId).Update("password", hashPwd).
// 	WithContext(ctx.Request().Context()).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *user)UpdateUserPassword(ctx *abstraction.Context, userId int, hashPwd string ) error {
	conn := r.checkTrx(ctx)
	
	var data model.UserEntityModel

		err := conn.Where("id = ?", userId).First(&data).
			WithContext(ctx.Request().Context()).Error
		if err != nil {
			return err
		}
		err = conn.Model(&data).UpdateColumn("password_hash", hashPwd ).
			WithContext(ctx.Request().Context()).Error
		if err != nil {
			return err
		}
		return  nil
}


// func (r *juz) Update(ctx *abstraction.Context, id *int, e *model.JuzEntity) (*model.JuzEntityModel, error) {
// 	conn := r.CheckTrx(ctx)

// 	var data model.UserEntityModel

// 	err := conn.Where("id = ?", id).First(&data).
// 		WithContext(ctx.Request().Context()).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	data.UserEntity = *e
// 	var Nama = ctx.Auth.Name
// 	data.ModifiedBy = Nama
// 	err = conn.Model(data).UpdateColumns(&data).
// 		WithContext(ctx.Request().Context()).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func (r *user)InvalidateAllSessions(ctx *abstraction.Context, userID int) {
	conn := r.checkTrx(ctx)
	dbc := conn.Where("user_id = ?", userID).Unscoped()
	if dbc.Error != nil {
		log.Println("An error occurred while deleting user sessions", dbc.Error.Error())
	}
}