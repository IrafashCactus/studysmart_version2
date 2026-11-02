package service

import (
	"study-smart-backend/internal/model"
	"study-smart-backend/internal/repository"

	"github.com/google/uuid"
)

type GamificationService struct {
	gameRepo *repository.GamificationRepository
}

func NewGamificationService(gameRepo *repository.GamificationRepository) *GamificationService {
	return &GamificationService{gameRepo: gameRepo}
}

type GamificationProfile struct {
	XP           *model.StudentXP      `json:"xp"`
	Streak       *model.Streak         `json:"streak"`
	Quests       []model.QuestProgress `json:"quests"`
	RecentBadges []string              `json:"recentBadges"`
}

func (s *GamificationService) GetProfile(userID uuid.UUID) (*GamificationProfile, error) {
	xp, err := s.gameRepo.GetXP(userID)
	if err != nil {
		return nil, err
	}

	streak, err := s.gameRepo.GetStreak(userID)
	if err != nil {
		return nil, err
	}

	quests, err := s.gameRepo.GetQuests(userID)
	if err != nil {
		return nil, err
	}

	badges := []string{"First Login", "Flame Keeper", "Math Pioneer"}
	if streak.CurrentStreak >= 7 {
		badges = append(badges, "7-Day Streak Master")
	}
	if xp.TotalXP >= 500 {
		badges = append(badges, "500 XP Club")
	}

	return &GamificationProfile{
		XP:           xp,
		Streak:       streak,
		Quests:       quests,
		RecentBadges: badges,
	}, nil
}

func (s *GamificationService) UseStreakFreeze(userID uuid.UUID) (*model.Streak, error) {
	return s.gameRepo.UseStreakFreeze(userID)
}

func (s *GamificationService) ClaimQuest(userID, questID uuid.UUID) (int, error) {
	return s.gameRepo.ClaimQuest(userID, questID)
}

func (s *GamificationService) GetLeaderboard(scope string) ([]repository.LeaderboardUser, error) {
	return s.gameRepo.GetLeaderboard(scope)
}

func (s *GamificationService) CompleteFocusSession(userID uuid.UUID, durationMinutes int) (int, bool, error) {
	// Pomodoro focus session: award 15 XP for 25 minutes, or proportional
	bonusXP := 15
	if durationMinutes >= 50 {
		bonusXP = 35
	} else if durationMinutes < 20 {
		bonusXP = 5
	}

	session := model.StudySession{
		UserID:        userID,
		SessionMode:   "focus",
		DurationSec:   durationMinutes * 60,
		IsFocusMode:   true,
		BonusXPEarned: bonusXP,
		CardsViewed:   0,
		QuizzesTaken:  0,
	}
	s.gameRepo.SaveStudySession(&session)

	_, leveledUp, err := s.gameRepo.AddXP(userID, bonusXP)
	s.gameRepo.UpdateStreakActivity(userID)

	return bonusXP, leveledUp, err
}
