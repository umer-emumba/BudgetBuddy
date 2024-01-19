package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	Database string
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type AppConfig struct {
	Port        int
	FrontendUrl string
	DatabaseConfig
	JWTConfig
	SMTPConfig
}

var AppCfg *AppConfig

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

	AppCfg = &AppConfig{
		Port:        port,
		FrontendUrl: os.Getenv("FRONTEND_URL"),
		DatabaseConfig: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     dbPort,
			Database: os.Getenv("DB_DATABASE"),
		},
		JWTConfig: JWTConfig{
			Secret:             os.Getenv("JWT_SECRET"),
			AccessTokenExpiry:  accessTokenExpiry,
			RefreshTokenExpiry: refreshTokenExpiry,
		},
		SMTPConfig: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Port:     smtpPort,
		},
	}
}
