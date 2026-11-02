package handler

import (
	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TutorHandler struct {
	tutorSvc *service.TutorService
}

func NewTutorHandler(tutorSvc *service.TutorService) *TutorHandler {
	return &TutorHandler{tutorSvc: tutorSvc}
}

func (h *TutorHandler) ExplainProblem(c *gin.Context) {
	var req service.ExplainProblemRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.ProblemText == "" {
		pkg.SendBadRequest(c, "problemText is required")
		return
	}

	explanation, err := h.tutorSvc.ExplainProblem(req)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, explanation, "step-by-step AI problem breakdown")
}

func (h *TutorHandler) GetAdaptiveRecommendations(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	recs, err := h.tutorSvc.GetAdaptiveRecommendations(userID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, recs, "adaptive learning recommendations")
}

