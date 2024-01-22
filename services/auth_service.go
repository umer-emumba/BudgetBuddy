package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/umer-emumba/BudgetBuddy/config"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type AuthService struct {
	userRepository repositories.UserRepository
	helper         *utils.Helper
}

func NewAuthService() AuthService {
	return AuthService{
		userRepository: repositories.NewUserRepository(),
		helper:         utils.NewHelper(),
	}
}

func (service AuthService) SignUp(dto dtos.SignupDto) (types.Message, error) {
	msg := types.Message{}
	emailCount, err := service.userRepository.CountByEmail(dto.Email)
	if err != nil {
		return msg, err
	}
	if emailCount > 0 {
		return msg, errors.New("email should be unique")
	}
	dto.Password, err = service.helper.CreateHash(dto.Password)
	if err != nil {
		return msg, err
	}
	user := models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
	createUserError := service.userRepository.CreateUser(user)
	if createUserError != nil {
		return msg, createUserError
	}

	token, tokenErr := service.helper.CreateVerificationToken(int(user.ID))
	if tokenErr != nil {
		return msg, tokenErr
	}

	mailOptions := types.MailOptions{
		To:      dto.Email,
		Subject: "Account Verification",
		Body:    fmt.Sprintf("Please click on following link to activate your account <a href='%s?%s'>Activate Account</a>", config.AppCfg.FrontendUrl, token),
	}

	client := config.CreateAsynqClient()
	defer client.Close()
	task, err := utils.NewEmailDeliveryTask(mailOptions)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	msg.Message = "Account Created Successfully"
	return msg, nil
}
