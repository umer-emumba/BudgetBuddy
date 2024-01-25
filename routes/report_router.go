package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/handlers"
	"github.com/umer-emumba/BudgetBuddy/middleware"
)

func setupReportRoutes(router *gin.Engine) {
	handler := handlers.NewReportHandler()
	reportRoutes := router.Group("/api/reports")

	reportRoutes.GET("/by_interval/:interval", middleware.AuthMiddleware(), handler.GetReportByInterval)
	reportRoutes.GET("/by_category", middleware.AuthMiddleware(), handler.GetReportByCategory)
}
