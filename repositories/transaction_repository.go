package repositories

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionTypes() ([]*models.TransactionType, error)
	GetCategories(transactionTypeId int) ([]*models.Category, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{models.DB}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return nil
}

func (r *transactionRepository) GetTransactionTypes() ([]*models.TransactionType, error) {
	var transactionTypes []*models.TransactionType
	err := r.db.Find(&transactionTypes).Error
	return transactionTypes, err
}

func (r *transactionRepository) GetCategories(transactionTypeId int) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.Where("transaction_type_id = ?", transactionTypeId).Find(&categories).Error
	return categories, err
}
