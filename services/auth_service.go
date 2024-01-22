package services

import (
	"errors"
	"fmt"

	config "github.com/umer-emumba/BudgetBuddy/configs"
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

func (service AuthService) SignUp(dto dtos.SignupDto) (interface{}, error) {
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
	createUserError := service.userRepository.CreateUser(user)
	if createUserError != nil {
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

	service.helper.SendMail(mailOptions)

	return map[string]string{
		"message": "Account created successfully",
	}, nil
}
