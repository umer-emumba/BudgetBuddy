package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/umer-emumba/BudgetBuddy/configs"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

func SetupRoutes(router *gin.Engine) {
	setupAuthRoutes(router)
	router.GET("/job_stats", func(ctx *gin.Context) {
		queues, err := config.RmqConnection.GetOpenQueues()
		if err != nil {
			panic(err)
		}

		stats, err := config.RmqConnection.CollectStats(queues)
		if err != nil {
			panic(err)
		}
		fmt.Println(stats)
		html := stats.GetHtml("", "")

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})
	router.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, http.StatusNotFound, "Requested resource not found")
	})
}
