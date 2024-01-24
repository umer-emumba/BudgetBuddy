package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/config"
	"github.com/umer-emumba/BudgetBuddy/middleware"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/routes"
	"go.uber.org/zap"
)

func main() {

	//initliaze config and database
	config.LoadConfig()
	models.InitDB()
	config.InitLogger()

	// Initialize Gin router
	router := gin.Default()

	router.Use(middleware.GinZapLogger(config.Logger, time.RFC3339, true))

	// Setup routes
	router.Static("/uploads", "./public/uploads")
	routes.SetupRoutes(router)

	//setup asynq server
	routes.SetupAsynqServeMux()

	// Start the server
	serverAddr := fmt.Sprintf(":%d", config.AppCfg.Port)

	err := router.Run(serverAddr)
	if err != nil {
		config.Logger.Fatal("Failed to start the server: ", zap.Error(err))
	} else {
		config.Logger.Info("Server is running on http://localhost%s", zap.String("message", serverAddr))
	}

}
