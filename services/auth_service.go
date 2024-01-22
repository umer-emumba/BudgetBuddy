package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	config "github.com/umer-emumba/BudgetBuddy/configs"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

var wg sync.WaitGroup

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
	emailCount, err := service.userRepository.CountByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if emailCount > 0 {
		return nil, errors.New("email should be unique")
	}
	dto.Password, err = service.helper.CreateHash(dto.Password)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
	if createUserError := service.userRepository.CreateUser(user); createUserError != nil {
		return nil, createUserError
	}

	claims := types.JwtToken{
		Id:        int(user.ID),
		UserType:  types.User,
		TokenType: types.EmailVerification,
	}
	token, tokenErr := service.helper.CreateToken(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	mailOptions := types.MailOptions{
		To:      dto.Email,
		Subject: "Account Verification",
		Body:    fmt.Sprintf("Please click on following link to activate your account <a href='%s?%s'>Activate Account</a>", config.AppCfg.FrontendUrl, token),
	}

	queueService, queueErr := GetQueueService()
	if queueErr != nil {
		return nil, queueErr
	}
	queue := queueService.GetQueue("default")
	job := types.QueueJob{
		Name:     "send-email",
		MailData: mailOptions,
	}
	jobStr, jsonErr := json.Marshal(job)
	if jsonErr != nil {
		return nil, jsonErr
	}
	queue.Publish(string(jobStr))

	return types.Message{
		"message": "Account created successfully",
	}, nil
}
