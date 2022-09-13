package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"quran/internal/abstraction"
	"quran/internal/dto"
	"quran/internal/factory"
	"quran/internal/model"
	"quran/internal/repository"
	res "quran/pkg/util/response"
	"quran/pkg/util/trxmanager"
	"quran/utils"
	uuid "github.com/satori/go.uuid"
	kafka "github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var KafkaConn *kafka.Conn

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
	ForgotPassword(ctx *abstraction.Context, payload *dto.ForgotPasswordRequest)
	GetNewPassword(ctx *abstraction.Context, payload *dto.NewPassword ) (string, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	token, err := data.GenerateToken()

	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var data *model.UserEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.UserEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AuthRegisterResponse{
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) ForgotPassword(ctx *abstraction.Context, payload *dto.ForgotPasswordRequest) (*dto.ForgotPasswordPayload, error) {
	var result *dto.ForgotPasswordPayload

	user, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		log.Println("Invalid User .. User doesn't exists")
		return result ,err
	}

	cacheKey := uuid.NewV4().String()
	cacheValue := user.Name
	Cache.setDetails(cacheKey, cacheValue)
	Cache.setExpiry(cacheKey, 10000)

	content := fmt.Sprintf("Post new password on given link: http://localhost:3030/auth/getNewPassword/%s to update your password. Request is valid for 5 minutes.", cacheKey)
	kafkaPayload := dto.ForgotPasswordPayload{
		To:      user.Email,
		From:    os.Getenv("EMAIL"),
		Text:    content,
		Subject: "Password Reset",
	}
	log.Println("kafka Payload :: ", kafkaPayload)
	payloa, err := json.Marshal(kafkaPayload)
	if err != nil {
		return result, err
	}
	_, err = KafkaConn.WriteMessages(
		kafka.Message{Value: payloa},
	)
	if err != nil {
		log.Println("failed to write messages:", err)
	}
	return result, err
}

var KafkaConn *kafka.Conn


// func (s *service) SendMail(ctx *abstraction.Context , payload *dto.FindById)(*model.KafkaPayload, error) {
// 	var user *model.UserEntityModel
// 	var input []model.SendMailInput
// 	var processedIds []string
// 	var result *model.KafkaPayload

// 	userIdentifier := int(ctx.Auth.ID)
// 	log.Println("user id :: ", userIdentifier)

// 	err := ctx.Bind(&input)
// 	if err != nil {
// 		return result, err
// 	}

// 	user, err = s.Repository.GetUserById(ctx, &payload.ID)
// 	if err != nil {
// 		return result,err
// 	}

// 	// if (user.NotificationCounter == 0) || (len(input) > user.NotificationCounter) {
// 	// 	return err
// 	// }

// 	for _, detail := range input {

// 		var template model.Template
// 		var clientRecord model.Client
// 		var clientDetails []model.Field
// 		var err error

// 		log.Println("User Id :: ", userIdentifier)
// 		log.Println("Template Id :: ", detail.TemplateId)
// 		template, err = s.Repository.GetTemplateById(ctx,detail.TemplateId, userIdentifier)
// 		if err != nil {
// 			// msg := fmt.Sprintf("Template with templateId %s not found !!", strconv.FormatUint(uint64(detail.TemplateId), 10))
// 			// resp := echo.Map{"status": "Ok", "message": msg}
// 			return result,err
// 		}
// 		log.Println("template details :: ", template)

// 		clientRecord, err = s.Repository.GetClientByIdUserId(ctx,strconv.FormatUint(uint64(detail.ClientId), 10), userIdentifier)
// 		if err != nil {
// 			// msg := fmt.Sprintf("Client with clientId %s not found !!", strconv.FormatUint(uint64(detail.ClientId), 10))
// 			// resp := echo.Map{"status": "Ok", "message": msg}
// 			return result,err
// 		}
// 		log.Println("ClientDetails :: ", clientRecord)

// 		clientDetails, _ = s.Repository.GetClientFields(ctx,detail.ClientId)
// 		log.Println("Client Fields :: ", clientDetails)

// 		subject := template.Subject
// 		bodyContent := template.Body

// 		if len(clientDetails) > 0 {
// 			for _, field := range clientDetails {
// 				oldVal := fmt.Sprintf("{{ %s }}", field.Key)
// 				newVal := field.Value
// 				subject = strings.Replace(subject, oldVal, newVal, -1)
// 				bodyContent = strings.Replace(bodyContent, oldVal, newVal, -1)
// 			}
// 		}

// 		kafkaPayload := model.KafkaPayload{
// 			To:      clientRecord.MailId,
// 			From:    user.Name,
// 			Subject: subject,
// 			Text:    bodyContent,
// 		}
// 		auditLog := model.Audit{
// 			To:           clientRecord.MailId,
// 			FromUser:     user.ID,
// 			TemplateID:   template.ID,
// 			TemplateName: template.Name,
// 		}

// 		_ = s.Repository.SaveAuditLog(ctx, auditLog)

// 		log.Println("kafka Payload :: ", kafkaPayload)
// 		payload, err := json.Marshal(kafkaPayload)
// 		if err != nil {
// 			return result,err
// 		}
// 		_, err = KafkaConn.WriteMessages(
// 			kafka.Message{Value: payload},
// 		)
// 		if err != nil {
// 			log.Println("failed to write messages:", err)
// 		}
// 		processedIds = append(processedIds, strconv.FormatUint(uint64(detail.ClientId), 10))
// 	}
// 	// conn := s.checkTrx(ctx)
// 	// count := len(processedIds)
// 	// user.NotificationCounter = user.NotificationCounter - count
// 	ctx.Trx.Db.Save(&user)
// 	// msg := fmt.Sprintf("Processed clientIds: %s", strings.Join(processedIds, ","))
// 	// resp := echo.Map{"status": "OK", "message": msg}
// 	return result, err
// }


func (s *service) GetNewPassword(ctx *abstraction.Context, payload *dto.NewPassword) (string, error) {
	cacheKey := ctx.Param("id")
	mail, _ := Cache.getDetails(cacheKey)
	// var data *model.UserEntity
	if mail != "" {
	
		user, err := s.Repository.GetUserByName(ctx,mail)
		if err != nil {
			return "error",err
		}
		hashPassword, _ := utils.HashPassword(payload.Password)
		err = s.Repository.UpdateUserPassword(ctx, int(user.ID),  hashPassword)
		if err != nil {
			return "error",err
		}
		s.Repository.InvalidateAllSessions(ctx,int(user.ID))
		return "error",err
	}
	return "Request to change password has been expired" , err
}
