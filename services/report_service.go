package services

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type ReportService struct {
	repo   repositories.TransactionRepository
	helper *utils.Helper
}

func NewReportService() ReportService {
	return ReportService{
		repo:   repositories.NewTransactionRepository(),
		helper: utils.NewHelper(),
	}
}

func (service ReportService) GetReportByInterval(user *models.User, interval string) ([]*types.IntervalReport, error) {
	return service.repo.GetReportByInterval(user.ID, interval)
}

func (service ReportService) GetReportByCategory(user *models.User) ([]*types.CategoryReport, error) {
	return service.repo.GetReportByCategory(user.ID)
}
