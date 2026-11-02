package handler

import (
	"net/http"

	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health-check HTTP requests.
type HealthHandler struct {
	svc *service.HealthService
}

// NewHealthHandler creates a new HealthHandler.
func NewHealthHandler(svc *service.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

// Check godoc
// @Summary     Health check
// @Description Returns the health status of the API and database.
// @Tags        health
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Failure     503 {object} map[string]interface{}
// @Router      /api/v1/health [get]
func (h *HealthHandler) Check(c *gin.Context) {
	status, err := h.svc.Check(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"data":    status,
			"message": "service unhealthy",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    status,
		"message": "ok",
	})
}
