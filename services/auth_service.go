package services

import (
	"errors"

	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
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

	return map[string]string{
		"message": "Account created successfully",
	}, nil
}
