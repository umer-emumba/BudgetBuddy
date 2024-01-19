package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

func authMiddleware() gin.HandlerFunc {

	utilties := utils.NewHelper()

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := utilties.VerifyToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()

	}
}
