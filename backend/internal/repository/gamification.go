package repository

import (
	"math"
	"time"

	"study-smart-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GamificationRepository struct {
	db *gorm.DB
}

func NewGamificationRepository(db *gorm.DB) *GamificationRepository {
	return &GamificationRepository{db: db}
}

func (r *GamificationRepository) GetXP(userID uuid.UUID) (*model.StudentXP, error) {
	var xp model.StudentXP
	err := r.db.Where("user_id = ?", userID).First(&xp).Error
	if err != nil {
		// Create default if missing
		xp = model.StudentXP{
			UserID:       userID,
			TotalXP:      0,
			CurrentLevel: 1,
			XPToNext:     100,
			UpdatedAt:    time.Now(),
		}
		r.db.Create(&xp)
	}
	return &xp, nil
}

func (r *GamificationRepository) AddXP(userID uuid.UUID, amount int) (*model.StudentXP, bool, error) {
	xp, err := r.GetXP(userID)
	if err != nil {
		return nil, false, err
	}

	xp.TotalXP += int64(amount)
	oldLevel := xp.CurrentLevel

	// Level Formula: Level = max(1, floor((TotalXP / 100) ^ (1/1.5)))
	newLevel := int(math.Floor(math.Pow(float64(xp.TotalXP)/100.0, 1.0/1.5))) + 1
	if newLevel < 1 {
		newLevel = 1
	}

	leveledUp := newLevel > oldLevel
	xp.CurrentLevel = newLevel

	// Next level threshold calculation
	nextThreshold := int64(100 * math.Pow(float64(newLevel), 1.5))
	xp.XPToNext = int(nextThreshold - xp.TotalXP)
	if xp.XPToNext < 10 {
		xp.XPToNext = 50
	}
	xp.UpdatedAt = time.Now()

	err = r.db.Save(xp).Error
	return xp, leveledUp, err
}

func (r *GamificationRepository) GetStreak(userID uuid.UUID) (*model.Streak, error) {
	var streak model.Streak
	err := r.db.Where("user_id = ?", userID).First(&streak).Error
	if err != nil {
		streak = model.Streak{
			UserID:           userID,
			CurrentStreak:    1,
			LongestStreak:    1,
			LastActivityDate: time.Now().Format("2006-01-02"),
			FreezePasses:     2,
			UpdatedAt:        time.Now(),
		}
		r.db.Create(&streak)
	}
	return &streak, nil
}

func (r *GamificationRepository) UpdateStreakActivity(userID uuid.UUID) (*model.Streak, error) {
	streak, err := r.GetStreak(userID)
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	if streak.LastActivityDate == today {
		// Already updated today
		return streak, nil
	}

	if streak.LastActivityDate == yesterday {
		// Consecutive day!
		streak.CurrentStreak++
		if streak.CurrentStreak > streak.LongestStreak {
			streak.LongestStreak = streak.CurrentStreak
		}
	} else if streak.LastActivityDate != "" {
		// Missed a day! Check if freeze pass is available
		if streak.FreezePasses > 0 && streak.LastFreezeUsedDate != yesterday {
			streak.FreezePasses--
			streak.LastFreezeUsedDate = yesterday
			// Streak is saved by freeze pass!
		} else {
			streak.CurrentStreak = 1
		}
	} else {
		streak.CurrentStreak = 1
	}

	streak.LastActivityDate = today
	streak.UpdatedAt = time.Now()
	err = r.db.Save(streak).Error
	return streak, err
}

func (r *GamificationRepository) UseStreakFreeze(userID uuid.UUID) (*model.Streak, error) {
	streak, err := r.GetStreak(userID)
	if err != nil {
		return nil, err
	}

	if streak.FreezePasses > 0 {
		streak.FreezePasses--
		streak.LastFreezeUsedDate = time.Now().Format("2006-01-02")
		streak.UpdatedAt = time.Now()
		err = r.db.Save(streak).Error
	}
	return streak, err
}

func (r *GamificationRepository) GetQuests(userID uuid.UUID) ([]model.QuestProgress, error) {
	var progresses []model.QuestProgress
	err := r.db.Preload("Quest").Where("user_id = ?", userID).Find(&progresses).Error
	if err != nil || len(progresses) == 0 {
		// If no progress records, populate from available quests
		var allQuests []model.Quest
		r.db.Find(&allQuests)
		for _, q := range allQuests {
			prog := model.QuestProgress{
				UserID:       userID,
				QuestID:      q.ID,
				Quest:        q,
				CurrentCount: 0,
				IsCompleted:  false,
				Claimed:      false,
			}
			r.db.Create(&prog)
			progresses = append(progresses, prog)
		}
	}
	return progresses, nil
}

func (r *GamificationRepository) ClaimQuest(userID, questID uuid.UUID) (int, error) {
	var prog model.QuestProgress
	err := r.db.Preload("Quest").Where("user_id = ? AND quest_id = ?", userID, questID).First(&prog).Error
	if err != nil {
		return 0, err
	}

	if prog.Claimed {
		return 0, nil
	}

	prog.Claimed = true
	r.db.Save(&prog)

	reward := prog.Quest.XPReward
	r.AddXP(userID, reward)
	return reward, nil
}

type LeaderboardUser struct {
	UserID       uuid.UUID `json:"userId"`
	FullName     string    `json:"fullName"`
	AvatarURL    string    `json:"avatarUrl"`
	TotalXP      int64     `json:"totalXp"`
	CurrentLevel int       `json:"currentLevel"`
	Streak       int       `json:"streak"`
	Rank         int       `json:"rank"`
}

func (r *GamificationRepository) GetLeaderboard(scope string) ([]LeaderboardUser, error) {
	// Query top students joined with XP and streaks
	rows, err := r.db.Raw(`
		SELECT u.id as user_id, u.full_name, u.avatar_url, 
		       COALESCE(x.total_xp, 0) as total_xp, 
		       COALESCE(x.current_level, 1) as current_level,
		       COALESCE(s.current_streak, 0) as streak
		FROM users u
		LEFT JOIN student_xps x ON x.user_id = u.id
		LEFT JOIN streaks s ON s.user_id = u.id
		WHERE u.role = 'student'
		ORDER BY total_xp DESC
		LIMIT 20
	`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []LeaderboardUser
	rank := 1
	for rows.Next() {
		var lu LeaderboardUser
		if err := rows.Scan(&lu.UserID, &lu.FullName, &lu.AvatarURL, &lu.TotalXP, &lu.CurrentLevel, &lu.Streak); err == nil {
			lu.Rank = rank
			results = append(results, lu)
			rank++
		}
	}
	return results, nil
}

func (r *GamificationRepository) SaveStudySession(session *model.StudySession) error {
	return r.db.Create(session).Error
}

