package services

import (
	"time"

	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type TransactionService struct {
	repo   repositories.TransactionRepository
	helper *utils.Helper
}

func NewTransactionService() TransactionService {
	return TransactionService{
		repo:   repositories.NewTransactionRepository(),
		helper: utils.NewHelper(),
	}
}

func (service TransactionService) AddTransaction(user *models.User, dto dtos.CreateTransactionDto) (types.Message, error) {
	msg := types.Message{}

	dateTime, parseErr := time.Parse("2006-01-02T15:04:05", dto.TransactionDate)
	if parseErr != nil {
		return msg, parseErr
	}
	transaction := &models.Transaction{
		UserID:            int(user.ID),
		Amount:            dto.Amount,
		TransactionTypeID: dto.TransactionTypeID,
		CategoryID:        dto.CategoryID,
		TransactionDate:   dateTime,
	}
	err := service.repo.CreateTransaction(transaction)
	if err != nil {
		return msg, err
	}

	msg.Message = "Transaction added successfully"
	return msg, nil
}

func (service TransactionService) GetTransactionTypes() ([]*models.TransactionType, error) {
	return service.repo.GetTransactionTypes()
}

func (service TransactionService) GetCategories(transactionTypeId int) ([]*models.Category, error) {
	return service.repo.GetCategories(transactionTypeId)
}
