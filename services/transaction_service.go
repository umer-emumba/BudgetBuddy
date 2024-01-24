package services

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
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

func (service TransactionService) AddTransaction() {

}

func (service TransactionService) GetTransactionTypes() ([]*models.TransactionType, error) {
	return service.repo.GetTransactionTypes()
}

func (service TransactionService) GetCategories(transactionTypeId int) ([]*models.Category, error) {
	return service.repo.GetCategories(transactionTypeId)
}
