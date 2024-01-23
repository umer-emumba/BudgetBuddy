package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
