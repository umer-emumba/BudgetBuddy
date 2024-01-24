package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/services"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler() TransactionHandler {
	return TransactionHandler{
		service: services.NewTransactionService(),
	}
}

func (h TransactionHandler) AddTransaction(c *gin.Context) {
	var dto dtos.CreateTransactionDto

	if err := c.ShouldBind(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}

	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := usr.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	data, err := h.service.AddTransaction(user, dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, data)
}

func (h TransactionHandler) GetTransactionTypes(c *gin.Context) {
	data, err := h.service.GetTransactionTypes()
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, data)
}

func (h TransactionHandler) GetCategories(c *gin.Context) {
	transactionTypeId := c.Param("transaction_type_id")
	if transactionTypeId == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "transaction_type_id is required")
		return
	}

	ID, convErr := strconv.Atoi(transactionTypeId)
	if convErr != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, convErr.Error())
		return
	}
	data, err := h.service.GetCategories(ID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, data)
}

func (h TransactionHandler) GetTransactions(c *gin.Context) {
	var dto dtos.PaginationDto

	if err := c.ShouldBindQuery(&dto); err != nil {
		message := utils.ConstructValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, message)
		return
	}

	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := usr.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	data, err := h.service.GetTransactions(user, dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, data)

}