package service

import (
	"study-smart-backend/internal/model"
	"study-smart-backend/internal/repository"

	"github.com/google/uuid"
)

type FeedService struct {
	feedRepo *repository.FeedRepository
	gameRepo *repository.GamificationRepository
}

func NewFeedService(feedRepo *repository.FeedRepository, gameRepo *repository.GamificationRepository) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
		gameRepo: gameRepo,
	}
}

type FeedResponse struct {
	Cards      []ContentCardDTO `json:"cards"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	HasMore    bool             `json:"hasMore"`
	StreakDays int              `json:"streakDays"`
}

type ContentCardDTO struct {
	model.ContentCard
	IsBookmarked bool `json:"isBookmarked"`
}

func (s *FeedService) GetFeed(userID *uuid.UUID, subjectCode, topicID string, page, perPage int) (*FeedResponse, error) {
	if perPage <= 0 {
		perPage = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perPage

	cards, total, err := s.feedRepo.GetCards(subjectCode, topicID, perPage, offset)
	if err != nil {
		return nil, err
	}

	streakDays := 1
	if userID != nil {
		if streak, err := s.gameRepo.GetStreak(*userID); err == nil {
			streakDays = streak.CurrentStreak
		}
	}

	dtoCards := make([]ContentCardDTO, len(cards))
	for i, c := range cards {
		dtoCards[i] = ContentCardDTO{
			ContentCard:  c,
			IsBookmarked: userID != nil && s.feedRepo.IsBookmarked(*userID, c.ID),
		}
	}

	return &FeedResponse{
		Cards:      dtoCards,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		HasMore:    int64(page*perPage) < total,
		StreakDays: streakDays,
	}, nil
}

func (s *FeedService) RecordInteraction(userID uuid.UUID, cardID uuid.UUID, action string, durationSec int) (int, bool, error) {
	interaction := model.CardInteraction{
		UserID:      userID,
		CardID:      cardID,
		Action:      action,
		DurationSec: durationSec,
	}
	if err := s.feedRepo.RecordInteraction(&interaction); err != nil {
		return 0, false, err
	}

	// Update daily streak
	s.gameRepo.UpdateStreakActivity(userID)

	xpEarned := 0
	leveledUp := false

	// Award micro-learning XP
	if action == "complete" {
		xpEarned = 5
		_, leveledUp, _ = s.gameRepo.AddXP(userID, xpEarned)
	} else if action == "view" && durationSec >= 5 {
		xpEarned = 2
		_, leveledUp, _ = s.gameRepo.AddXP(userID, xpEarned)
	}

	return xpEarned, leveledUp, nil
}

func (s *FeedService) ToggleBookmark(userID, cardID uuid.UUID) (bool, error) {
	return s.feedRepo.ToggleBookmark(userID, cardID)
}

func (s *FeedService) GetBookmarks(userID uuid.UUID) ([]model.ContentCard, error) {
	return s.feedRepo.GetBookmarks(userID)
}

func (s *FeedService) GetSubjects() ([]model.Subject, error) {
	return s.feedRepo.GetSubjects()
}

