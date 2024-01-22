package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/umer-emumba/BudgetBuddy/types"
)

var AppCfg *types.AppConfig

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("FAILED TO LOAD .env FILE")
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	accessTokenExpiry, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRY"))
	refreshTokenExpiry, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRY"))

	AppCfg = &types.AppConfig{
		Port:        port,
		FrontendUrl: os.Getenv("FRONTEND_URL"),
		DatabaseConfig: types.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     dbPort,
			Database: os.Getenv("DB_DATABASE"),
		},
		JWTConfig: types.JWTConfig{
			Secret:             os.Getenv("JWT_SECRET"),
			AccessTokenExpiry:  accessTokenExpiry,
			RefreshTokenExpiry: refreshTokenExpiry,
		},
		SMTPConfig: types.SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			User:     os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Port:     smtpPort,
			Sender:   os.Getenv("SMTP_SENDER"),
		},
		RedisConfig: types.RedisConfig{
			Host: os.Getenv("REDIS_HOST"),
		},
	}
}
