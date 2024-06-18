package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	})
}

func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  http.StatusText(statusCode),
		Message: message,
	})
}
