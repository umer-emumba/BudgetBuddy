package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/umer-emumba/BudgetBuddy/config"
	"github.com/umer-emumba/BudgetBuddy/types"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type Helper struct {
	appConfig *types.AppConfig
	validate  *validator.Validate
}

func NewHelper() *Helper {
	return &Helper{
		appConfig: config.AppCfg,
		validate:  validator.New(),
	}
}

func (helper *Helper) CreateToken(claims types.JwtToken) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(helper.appConfig.JWTConfig.Secret)
	return accessToken.SignedString(secretKey)
}

func (helper *Helper) VerifyToken(tokenString string) (*types.JwtToken, error) {
	secretKey := []byte(helper.appConfig.JWTConfig.Secret)
	var jwtClaims types.JwtToken

	token, err := jwt.ParseWithClaims(tokenString, &types.JwtToken{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {

		return &jwtClaims, err
	}

	if !token.Valid {
		return &jwtClaims, errors.New("token is not valid")
	}

	if claims, ok := token.Claims.(*types.JwtToken); ok {
		return claims, nil
	} else {
		return &jwtClaims, err
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

func (helper *Helper) SendMail(options types.MailOptions) error {

	smtp := helper.appConfig.SMTPConfig

	m := gomail.NewMessage()
	m.SetHeader("From", smtp.Sender)
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)
	m.SetBody("text/html", options.Body)

	d := gomail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Password)

	if err := d.DialAndSend(m); err != nil {

		fmt.Println("Failed to send email:", err)
		return err
	}

	fmt.Println("Email sent successfully")
	return nil
}

func (helper *Helper) CreateVerificationToken(id int) (string, error) {
	claims := types.JwtToken{
		Id:        id,
		UserType:  types.User,
		TokenType: types.EmailVerification,
	}
	token, tokenErr := helper.CreateToken(claims)
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}
