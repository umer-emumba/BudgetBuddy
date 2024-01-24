package repositories

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionTypes() ([]*models.TransactionType, error)
	GetCategories(transactionTypeId int) ([]*models.Category, error)
	Count(userID uint, dto dtos.PaginationDto) (int, error)
	FindAll(userID uint, dto dtos.PaginationDto) ([]*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{models.DB}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
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

func (r *transactionRepository) Count(userID uint, dto dtos.PaginationDto) (int, error) {
	var count int64

	query := r.db.Model(&models.Transaction{})
	query = query.Where("user_id =?", userID)
	if dto.TransactionTypeId != 0 {
		query = query.Where("transaction_type_id =?", dto.TransactionTypeId)
	}
	result := query.Count(&count)
	return int(count), result.Error

}

func (r *transactionRepository) FindAll(userID uint, dto dtos.PaginationDto) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	query := r.db.Model(&models.Transaction{})
	query = query.Where("user_id =?", userID)
	if dto.TransactionTypeId != 0 {
		query = query.Where("transaction_type_id =?", dto.TransactionTypeId)
	}
	query = query.Limit(dto.Limit).Offset(dto.Offset * dto.Limit).Order("id " + dto.Order)
	err := query.Find(&transactions).Error
	return transactions, err

}