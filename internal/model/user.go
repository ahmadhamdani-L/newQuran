package model

import (
	"os"
	"quran/internal/abstraction"
	"quran/pkg/constant"
	"quran/pkg/util/date"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Name         string `json:"name" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	Email        string `json:"email" validate:"required,email" gorm:"index:idx_user_email,unique"`
	PasswordHash string `json:"-"`
	Password     string `json:"password" validate:"required" gorm:"-"`
	IsActive     bool   `json:"is_active" validate:"required"`
}

type UserEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	UserEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}



func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	m.hashPassword()
	m.Password = ""
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = *date.DateTodayLocal()
	m.ModifiedBy = m.Context.Auth.Name
	return
}

func (m *UserEntityModel) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserEntityModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    m.ID,
		"email": m.Email,
		"name":  m.Name,
		"phone": m.Phone,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}

type SendMailInput struct {
	TemplateId uint `json:"templateId"`
	ClientId   uint `json:"clientId"`
}

type Template struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(50);unique;not null"`
	Subject string `json:"subject" gorm:"type:varchar(50)"`
	Body    string `json:"body" gorm:"type:varchar(10000)"`
	UserID  uint
	User    UserEntityModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Client struct {
	gorm.Model
	Name       string `json:"name" gorm:"type:varchar(50)"`
	MailId     string `json:"mailID" valid:"email,required" gorm:"type:varchar(50);not null;unique_index:idx_first_second"`
	Phone      int    `json:"phone" valid:"required" gorm:"type:bigint;not null"`
	Preference string `json:"preference" gorm:"type:varchar(50)"`
	UserID     uint   `gorm:"unique_index:idx_first_second"`
	User       UserEntityModel   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Field struct {
	gorm.Model
	Key      string `json:"key" valid:"required" gorm:"type:varchar(50);not null;unique_index:idx_field"`
	Value    string `json:"value" valid:"required" gorm:"type:varchar(50);not null"`
	ClientID uint   `gorm:"unique_index:idx_field"`
	Client   Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type KafkaPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Text    string `json:"text"`
	Subject string `json:"subject"`
}

type Audit struct {
	gorm.Model
	To           string
	FromUser     uint `gorm:"not null"`
	TemplateName string
	TemplateID   uint
}

func (UserEntityModel) TableName() string {
	return "users"
}

