package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/services"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{
		authService: services.NewAuthService(),
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {

	var dto dtos.SignupDto
	if err := c.ShouldBind(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}
	data, error := h.authService.SignUp(dto)
	if error != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, data)
}

func (h *AuthHandler) VerifyAccount(c *gin.Context) {

	var dto dtos.AccountVerificationDto
	if err := c.ShouldBind(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}
	data, error := h.authService.VerifyAccount(dto)
	if error != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, data)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var dto dtos.SignInDto

	if err := c.ShouldBind(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}

	data, error := h.authService.SignIn(dto)
	if error != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, data)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, user)

}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var dto dtos.UpdateProfileDTO

	usr, exists := c.Get("user")
	if !exists {

		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, ok := usr.(*models.User)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := c.ShouldBindWith(&dto, binding.FormMultipart); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}

	data, error := h.authService.UpdateProfile(user, dto)
	if error != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, data)

}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "email is required")
		return
	}
	data, err := h.authService.ForgotPassword(email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, data)

}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var dto dtos.PasswordResetDto

	if err := c.ShouldBind(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}

	data, err := h.authService.ResetPassword(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, data)

}
