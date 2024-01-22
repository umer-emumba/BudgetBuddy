package models

import (
	"fmt"
	"log"

	"github.com/umer-emumba/BudgetBuddy/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func setupDB() (*gorm.DB, error) {

	appConfig := config.AppCfg
	databaseConfig := appConfig.DatabaseConfig

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database)

	// Open a connection and set up a connection pool
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(300)

	return db, nil
}

func migrate() {
	DB.AutoMigrate(&TransactionType{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Transaction{})
}

func InitDB() {

	db, err := setupDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	migrate()
}
