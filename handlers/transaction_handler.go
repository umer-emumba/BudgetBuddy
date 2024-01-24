package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/services"
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
