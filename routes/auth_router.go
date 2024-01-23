package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/handlers"
)

func setupAuthRoutes(router *gin.Engine) {
	handler := handlers.NewAuthHandler()
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/signup", handler.SignUp)
		authRoutes.POST("/verify", handler.VerifyAccount)
		authRoutes.POST("/signin", handler.SignIn)
	}

}
