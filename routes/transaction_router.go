package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/handlers"
	"github.com/umer-emumba/BudgetBuddy/middleware"
)

func setupTransactionRoutes(router *gin.Engine) {
	handler := handlers.NewTransactionHandler()
	tranRoutes := router.Group("/api/transactions")

	tranRoutes.POST("/", middleware.AuthMiddleware(), handler.AddTransaction)
	tranRoutes.GET("/", middleware.AuthMiddleware(), handler.GetTransactions)
	tranRoutes.GET("/transaction_types", middleware.AuthMiddleware(), handler.GetTransactionTypes)
	tranRoutes.GET("/categories/:transaction_type_id", middleware.AuthMiddleware(), handler.GetCategories)
	tranRoutes.GET("/:id", middleware.AuthMiddleware(), handler.GetTransactionDetails)
}
