package handler

import (
	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuizHandler struct {
	quizSvc *service.QuizService
}

func NewQuizHandler(quizSvc *service.QuizService) *QuizHandler {
	return &QuizHandler{quizSvc: quizSvc}
}

func (h *QuizHandler) GetQuizzes(c *gin.Context) {
	topicID := c.Query("topicId")
	quizzes, err := h.quizSvc.GetQuizzes(topicID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, quizzes, "practice quizzes")
}

func (h *QuizHandler) GetQuiz(c *gin.Context) {
	quizIDStr := c.Param("id")
	quizID, err := uuid.Parse(quizIDStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid quiz ID")
		return
	}

	quiz, questions, err := h.quizSvc.GetQuizDetails(quizID)
	if err != nil {
		pkg.SendNotFound(c, "quiz not found")
		return
	}

	pkg.SendSuccess(c, gin.H{
		"quiz":      quiz,
		"questions": questions,
	}, "quiz loaded")
}

func (h *QuizHandler) SubmitQuiz(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	var req service.QuizSubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.SendBadRequest(c, "invalid submit payload")
		return
	}

	result, err := h.quizSvc.SubmitQuiz(userID, req)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, result, "quiz evaluated")
}

