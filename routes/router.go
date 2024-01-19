package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

func SetupRoutes(router *gin.Engine) {
	setupAuthRoutes(router)
	router.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, 404, "Requested resource not found")
	})
}
