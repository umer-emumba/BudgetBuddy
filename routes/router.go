package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/umer-emumba/BudgetBuddy/config"
	"github.com/umer-emumba/BudgetBuddy/handlers"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

func SetupRoutes(router *gin.Engine) {
	setupAuthRoutes(router)
	setupTransactionRoutes(router)
	router.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, 404, "Requested resource not found")
	})
}

func SetupAsynqServeMux() {
	server := config.CreateAsyncQServer()
	mux := asynq.NewServeMux()
	handler := handlers.NewQueueTaskHandler()
	mux.HandleFunc(types.TypeEmailDelivery, handler.HandleEmailDeliveryTask)

	go func() {
		if err := server.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

}
