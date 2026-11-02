package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseEnvelope matches the AGENTS.md requirement: { "success": true, "data": {}, "message": "ok" }
type ResponseEnvelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func SendSuccess(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "ok"
	}
	c.JSON(http.StatusOK, ResponseEnvelope{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func SendCreated(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "created successfully"
	}
	c.JSON(http.StatusCreated, ResponseEnvelope{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func SendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ResponseEnvelope{
		Success: false,
		Data:    nil,
		Message: message,
	})
}

func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message)
}

func SendUnauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "unauthorized"
	}
	SendError(c, http.StatusUnauthorized, message)
}

func SendNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "resource not found"
	}
	SendError(c, http.StatusNotFound, message)
}

func SendServerError(c *gin.Context, message string) {
	if message == "" {
		message = "internal server error"
	}
	SendError(c, http.StatusInternalServerError, message)
}

