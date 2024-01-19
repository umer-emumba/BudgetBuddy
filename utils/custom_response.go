package utils

import (
	"github.com/gin-gonic/gin"
)

type CustomResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	response := CustomResponse{
		Success: true,
		Data:    data,
	}
	c.JSON(status, response)
}

func ErrorResponse(c *gin.Context, status int, message string) {
	response := CustomResponse{
		Success: false,
		Error:   message,
	}
	c.JSON(status, response)
}
