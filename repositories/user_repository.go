package repositories

import (
	"github.com/umer-emumba/BudgetBuddy/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(userID uint) (*models.User, error)
	GetUserByEmail(email string) *models.User
	GetAllUsers() ([]*models.User, error)
	CountByEmail(email string) (int64, error)
	SaveUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{models.DB}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	return &user, err
}

func (r *userRepository) GetUserByEmail(email string) *models.User {
	var user models.User
	r.db.Where("email = ?", email).First(&user)
	return &user
}

func (r *userRepository) CountByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) SaveUser(user *models.User) error {
	return r.db.Save(&user).Error
}
