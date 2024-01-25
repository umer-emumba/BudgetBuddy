package repositories

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/types/dtos"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	UpdateTransaction(ID int, dto dtos.UpdateTransactionDto) error
	GetTransactionTypes() ([]*models.TransactionType, error)
	GetCategories(transactionTypeId int) ([]*models.Category, error)
	Count(userID uint, dto dtos.PaginationDto) (int, error)
	FindAll(userID uint, dto dtos.PaginationDto) ([]*models.Transaction, error)
	FindDetails(userID uint, ID int) (*models.Transaction, error)
	FindOne(userID uint, ID int) (*models.Transaction, error)
	DeleteTransaction(userID uint, ID int) error
	GetReportByInterval(userID uint, interval string) ([]*types.IntervalReport, error)
	GetReportByCategory(userID uint) ([]*types.CategoryReport, error)
	MonthlyReport(userID uint) ([]types.MonthlyReport, error)
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

func (r *transactionRepository) FindOne(userID uint, ID int) (*models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Where("transactions.user_id =?", userID).Where("transactions.id=?", ID).First(&trans).Error
	return &trans, err
}

func (r *transactionRepository) FindDetails(userID uint, ID int) (*models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Joins("TransactionType").Joins("Category").Where("transactions.user_id =?", userID).Where("transactions.id=?", ID).First(&trans).Error
	return &trans, err
}

func (r *transactionRepository) UpdateTransaction(ID int, dto dtos.UpdateTransactionDto) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", ID).Updates(dto).Error
}

func (r *transactionRepository) DeleteTransaction(userID uint, ID int) error {
	return r.db.Where("user_id = ?", userID).Where("id = ?", ID).Delete(&models.Transaction{}).Error
}

func (r *transactionRepository) GetReportByInterval(userID uint, interval string) ([]*types.IntervalReport, error) {
	var result []*types.IntervalReport
	var err error
	if interval == "monthly" {
		err = r.db.Model(&models.Transaction{}).
			Select("CONCAT(MONTHNAME(transaction_date), ' ', YEAR(transaction_date)) AS `interval`, TransactionType.name as transaction_type, SUM(amount) as total_amount").
			Joins("TransactionType").
			Where("user_id=?", userID).
			Group("`interval`, transaction_type").
			Scan(&result).Error
	} else {
		err = r.db.Model(&models.Transaction{}).
			Select("YEAR(transaction_date) AS `interval`, TransactionType.name as transaction_type, SUM(amount) as total_amount").
			Joins("TransactionType").
			Where("user_id=?", userID).
			Group("`interval`, transaction_type").
			Scan(&result).Error
	}

	return result, err
}

func (r *transactionRepository) GetReportByCategory(userID uint) ([]*types.CategoryReport, error) {
	var result []*types.CategoryReport

	err := r.db.Model(&models.Transaction{}).
		Select(" Category.name as category, TransactionType.name as transaction_type, SUM(amount) as total_amount").
		Joins("TransactionType").
		Joins("Category").
		Where("user_id=?", userID).
		Group("`category`, transaction_type").
		Scan(&result).Error

	return result, err
}

func (r *transactionRepository) MonthlyReport(userID uint) ([]types.MonthlyReport, error) {
	var result []types.MonthlyReport

	err := r.db.Model(&models.Transaction{}).
		Select("TransactionType.name as transaction_type, SUM(amount) as total_amount").
		Joins("TransactionType").
		Joins("Category").
		Where("user_id=?", userID).
		Where("YEAR(transaction_date) = YEAR(CURRENT_DATE - INTERVAL 1 MONTH)").
		Where("MONTH(transaction_date) = MONTH(CURRENT_DATE - INTERVAL 1 MONTH)").
		Group("transaction_type").
		Scan(&result).Error

	return result, err
}
