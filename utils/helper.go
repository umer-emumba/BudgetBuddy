package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(helper.appConfig.JWTConfig.Secret)
	return token.SignedString(secretKey)
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

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMessage := ""

		for _, validationErr := range validationErrors {
			errorMessage += fmt.Sprintf("Rule '%s' failed for field '%s' , ", validationErr.ActualTag(), validationErr.Field())
		}
		return errorMessage
	} else {
		return "Request data binding failed"
	}
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

func (helper *Helper) CreateAccessToken(id int) (string, error) {
	expiryDuration := time.Duration(helper.appConfig.JWTConfig.AccessTokenExpiry)
	expirationTime := time.Now().Add(expiryDuration * time.Second)
	fmt.Println(expirationTime)
	claims := types.JwtToken{
		Id:        id,
		UserType:  types.User,
		TokenType: types.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token, tokenErr := helper.CreateToken(claims)
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}

func (helper *Helper) CreateRefreshToken(id int) (string, error) {
	expiryDuration := time.Duration(helper.appConfig.JWTConfig.RefreshTokenExpiry)
	expirationTime := time.Now().Add(expiryDuration * time.Second)
	claims := types.JwtToken{
		Id:        id,
		UserType:  types.User,
		TokenType: types.Refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token, tokenErr := helper.CreateToken(claims)
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}

func (helper *Helper) CreatePasswordResetToken(id int) (string, error) {
	expiryDuration := time.Duration(helper.appConfig.JWTConfig.AccessTokenExpiry)
	expirationTime := time.Now().Add(expiryDuration * time.Second)
	claims := types.JwtToken{
		Id:        id,
		UserType:  types.User,
		TokenType: types.PasswordReset,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token, tokenErr := helper.CreateToken(claims)
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}

func (helper *Helper) IsImage(file *multipart.FileHeader) bool {

	if file.Size == 0 {
		return false
	}

	// Check if the file type is allowed (you can customize this based on your requirements)
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	validImage := false
	for _, allowedType := range allowedTypes {
		if file.Header.Get("Content-Type") == allowedType {
			validImage = true
			break
		}
	}

	return validImage
}

func (helper *Helper) UploadFile(file *multipart.FileHeader, destinationDir string) (string, error) {
	fmt.Println(file.Header.Get("Content-Type"))
	// Ensure the destination directory exists
	if err := os.MkdirAll("public/"+destinationDir, 0755); err != nil {
		return "", err
	}

	// Generate a unique filename using a timestamp and the original filename
	uniqueFilename := generateUniqueFilename(file.Filename)

	// Create the destination file
	destinationPath := filepath.Join(destinationDir, uniqueFilename)
	newFile, err := os.Create("public/" + destinationPath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(newFile, uploadedFile)
	if err != nil {
		return "", err
	}

	return destinationPath, nil
}

func generateUniqueFilename(originalFilename string) string {
	// Use the current timestamp to generate a unique part of the filename
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Combine the timestamp and file extension to create a unique filename
	uniquePart := fmt.Sprintf("%d", timestamp)

	// Combine the unique part and the original filename
	uniqueFilename := fmt.Sprintf("%s_%s", uniquePart, originalFilename)

	// Replace any spaces with underscores in the filename
	uniqueFilename = filepath.Join(filepath.SplitList(uniqueFilename)...)
	return uniqueFilename
}
