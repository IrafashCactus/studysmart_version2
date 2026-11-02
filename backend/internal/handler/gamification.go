package handler

import (
	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GamificationHandler struct {
	gameSvc *service.GamificationService
}

func NewGamificationHandler(gameSvc *service.GamificationService) *GamificationHandler {
	return &GamificationHandler{gameSvc: gameSvc}
}

func (h *GamificationHandler) GetProfile(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	profile, err := h.gameSvc.GetProfile(userID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, profile, "gamification profile")
}

func (h *GamificationHandler) FreezeStreak(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	streak, err := h.gameSvc.UseStreakFreeze(userID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, streak, "streak freeze applied")
}

func (h *GamificationHandler) ClaimQuest(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	questIDStr := c.Param("questId")
	questID, err := uuid.Parse(questIDStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid questId")
		return
	}

	rewardXP, err := h.gameSvc.ClaimQuest(userID, questID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{
		"xpReward": rewardXP,
	}, "quest reward claimed")
}

func (h *GamificationHandler) GetLeaderboard(c *gin.Context) {
	scope := c.DefaultQuery("scope", "national")
	board, err := h.gameSvc.GetLeaderboard(scope)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, board, "leaderboard rankings")
}

type FocusSessionInput struct {
	DurationMinutes int `json:"durationMinutes" binding:"required,min=1"`
}

func (h *GamificationHandler) CompleteFocusSession(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	var input FocusSessionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.SendBadRequest(c, "durationMinutes required (min 1)")
		return
	}

	bonusXP, leveledUp, err := h.gameSvc.CompleteFocusSession(userID, input.DurationMinutes)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{
		"bonusXp":   bonusXP,
		"leveledUp": leveledUp,
	}, "pomodoro focus session completed! bonus XP awarded")
}

