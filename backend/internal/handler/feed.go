package handler

import (
	"strconv"

	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FeedHandler struct {
	feedSvc *service.FeedService
}

func NewFeedHandler(feedSvc *service.FeedService) *FeedHandler {
	return &FeedHandler{feedSvc: feedSvc}
}

func (h *FeedHandler) GetFeed(c *gin.Context) {
	subject := c.Query("subject")
	topicID := c.Query("topicId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	var userUID *uuid.UUID
	if val, exists := c.Get("userId"); exists {
		u := val.(uuid.UUID)
		userUID = &u
	}

	feed, err := h.feedSvc.GetFeed(userUID, subject, topicID, page, perPage)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, feed, "mind-feed loaded")
}

type InteractInput struct {
	CardID      string `json:"cardId" binding:"required"`
	Action      string `json:"action" binding:"required"` // 'view', 'complete', 'skip'
	DurationSec int    `json:"durationSec"`
}

func (h *FeedHandler) Interact(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	var input InteractInput
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.SendBadRequest(c, "cardId and action required")
		return
	}

	cardID, err := uuid.Parse(input.CardID)
	if err != nil {
		pkg.SendBadRequest(c, "invalid cardId format")
		return
	}

	xp, leveledUp, err := h.feedSvc.RecordInteraction(userID, cardID, input.Action, input.DurationSec)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{
		"xpEarned":  xp,
		"leveledUp": leveledUp,
	}, "interaction logged")
}

func (h *FeedHandler) ToggleBookmark(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	cardIDStr := c.Param("cardId")
	cardID, err := uuid.Parse(cardIDStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid cardId")
		return
	}

	isBookmarked, err := h.feedSvc.ToggleBookmark(userID, cardID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	msg := "card removed from bookmarks"
	if isBookmarked {
		msg = "card bookmarked for review"
	}

	pkg.SendSuccess(c, gin.H{
		"isBookmarked": isBookmarked,
	}, msg)
}

func (h *FeedHandler) GetBookmarks(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "authentication required")
		return
	}
	userID := val.(uuid.UUID)

	bookmarks, err := h.feedSvc.GetBookmarks(userID)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, bookmarks, "bookmarked cards")
}

func (h *FeedHandler) GetSubjects(c *gin.Context) {
	subjects, err := h.feedSvc.GetSubjects()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, subjects, "curriculum subjects")
}

