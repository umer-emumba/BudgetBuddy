package handlers

import (
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
		utils.ErrorResponse(c, 400, message)
		return
	}
	data, error := h.authService.SignUp(dto)
	if error != nil {
		utils.ErrorResponse(c, 400, error.Error())
		return
	}
	utils.SuccessResponse(c, 201, data)
}
