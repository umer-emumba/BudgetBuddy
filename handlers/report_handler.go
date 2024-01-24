package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umer-emumba/BudgetBuddy/models"
	"github.com/umer-emumba/BudgetBuddy/services"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type ReportHandler struct {
	service services.ReportService
}

func NewReportHandler() ReportHandler {
	return ReportHandler{
		service: services.NewReportService(),
	}
}

func (h ReportHandler) GetReportByInterval(c *gin.Context) {
	interval := c.Param("interval")
	if interval == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "interval is required")
		return
	}
	if interval != "monthly" && interval != "yearly" {
		utils.ErrorResponse(c, http.StatusBadRequest, "interval should be monthly or yearly ")
		return
	}

	usr, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, ok := usr.(*models.User)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	data, err := h.service.GetReportByInterval(user, interval)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, data)

}
