package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	config "github.com/umer-emumba/BudgetBuddy/configs"
	"golang.org/x/crypto/bcrypt"
)

type Helper struct {
	appConfig *config.AppConfig
	validate  *validator.Validate
}

func NewHelper() *Helper {
	return &Helper{
		appConfig: config.AppCfg,
		validate:  validator.New(),
	}
}

func (helper *Helper) VerifyToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return helper.appConfig.JWTConfig.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, err
	}
}

func (helper *Helper) CreateHash(data string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (helper *Helper) VerifyPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

func ConstructValidationError(err error) string {
	errorMessage := ""
	for _, err := range err.(validator.ValidationErrors) {

		errorMessage += fmt.Sprintf("Rule '%s' failed for field '%s' , ", err.ActualTag(), err.Field())
	}
	return errorMessage
}
