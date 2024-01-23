package services

import (
	"errors"
	"fmt"
	"log"
	"time"

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
	user := &models.User{
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

func (service AuthService) VerifyAccount(dto dtos.AccountVerificationDto) (types.Message, error) {
	msg := types.Message{}
	claims, jwtErr := service.helper.VerifyToken(dto.Token)
	if jwtErr != nil {
		return msg, jwtErr
	}

	if claims.UserType != types.User {
		return msg, errors.New("invalid token")
	}
	if claims.TokenType != types.EmailVerification {
		return msg, errors.New("invalid token")
	}

	user, userErr := service.userRepository.GetUserByID(uint(claims.Id))
	if userErr != nil {
		return msg, userErr
	}
	user.EmailVerifiedAt = time.Now()
	saveErr := service.userRepository.SaveUser(user)
	if saveErr != nil {
		return msg, saveErr
	}

	msg.Message = "Account Verified Successfully"
	return msg, nil
}

func (service AuthService) SignIn(dto dtos.SignInDto) (types.Login, error) {
	login := types.Login{}
	user := service.userRepository.GetUserByEmail(dto.Email)
	if user.ID == 0 {
		return login, errors.New("invalid credentials")
	}
	passwordVerification := service.helper.VerifyPassword(user.Password, dto.Password)
	if !passwordVerification {
		return login, errors.New("invalid credentials")
	}

	if user.EmailVerifiedAt.IsZero() {
		return login, errors.New("please verify your account first")
	}

	accessToken, accessTokenErr := service.helper.CreateAccessToken(int(user.ID))
	if accessTokenErr != nil {
		return login, accessTokenErr
	}

	refreshToken, refreshTokenErr := service.helper.CreateRefreshToken(int(user.ID))
	if refreshTokenErr != nil {
		return login, refreshTokenErr
	}

	login.AccessToken = accessToken
	login.RefreshToken = refreshToken

	return login, nil

}

func (service AuthService) UpdateProfile(user *models.User, dto dtos.UpdateProfileDTO) (types.Message, error) {
	msg := types.Message{}

	if dto.Image != nil {
		isValidImage := service.helper.IsImage(dto.Image)
		if !isValidImage {
			return msg, errors.New("image should be a valid image")
		}
		imageUrl, uploadErr := service.helper.UploadFile(dto.Image, "uploads")
		if uploadErr != nil {
			return msg, uploadErr
		}

		user.ImageUrl = imageUrl
	}

	user.Name = dto.Name
	service.userRepository.SaveUser(user)
	msg.Message = "Profile Updated"

	return msg, nil

}
