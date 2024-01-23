package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/handlers"
	"github.com/umer-emumba/BudgetBuddy/middleware"
)

func setupAuthRoutes(router *gin.Engine) {
	handler := handlers.NewAuthHandler()
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/signup", handler.SignUp)
		authRoutes.POST("/verify", handler.VerifyAccount)
		authRoutes.POST("/signin", handler.SignIn)
		authRoutes.GET("/profile", middleware.AuthMiddleware(), handler.GetProfile)
		authRoutes.PUT("/update_profile", middleware.AuthMiddleware(), handler.UpdateProfile)
	}

}
