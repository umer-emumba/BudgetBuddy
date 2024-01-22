package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/umer-emumba/BudgetBuddy/configs"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/routes"
)

func main() {

	//initliaze config and database
	config.LoadConfig()
	models.InitDB()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	router.Static("/public", "./public")
	routes.SetupRoutes(router)

	// Start the server
	serverAddr := fmt.Sprintf(":%d", config.AppCfg.Port)

	err := router.Run(serverAddr)
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	} else {
		log.Printf("Server is running on http://localhost%s", serverAddr)
	}
}
