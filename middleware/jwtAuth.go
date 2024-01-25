package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	utilties := utils.NewHelper()
	repo := repositories.NewUserRepository()

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		token := strings.Split(authorization, " ")[1]
		if token == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		claims, err := utilties.VerifyToken(token)

		if err != nil || claims.UserType != types.User || claims.TokenType != types.Access {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		user, userErr := repo.GetUserByID(uint(claims.Id))
		if userErr != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()

	}
}
