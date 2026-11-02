package repository

import (
	"time"

	"study-smart-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (r *QuizRepository) GetQuizzes(topicID string) ([]model.Quiz, error) {
	var quizzes []model.Quiz
	q := r.db.Preload("Topic").Preload("Topic.Subject").Where("is_published = ?", true)

	if topicID != "" {
		if uid, err := uuid.Parse(topicID); err == nil {
			q = q.Where("topic_id = ?", uid)
		}
	}

	err := q.Order("created_at DESC").Find(&quizzes).Error
	return quizzes, err
}

func (r *QuizRepository) GetQuizByID(id uuid.UUID) (*model.Quiz, error) {
	var quiz model.Quiz
	err := r.db.Preload("Topic").Preload("Topic.Subject").First(&quiz, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *QuizRepository) GetQuestionsForTopic(topicID uuid.UUID, limit int) ([]model.Question, error) {
	var questions []model.Question
	err := r.db.Where("topic_id = ? AND status = ?", topicID, "approved").
		Limit(limit).
		Find(&questions).Error
	return questions, err
}

func (r *QuizRepository) GetAllQuestions(limit int) ([]model.Question, error) {
	var questions []model.Question
	err := r.db.Preload("Topic").Preload("Topic.Subject").
		Where("status = ?", "approved").
		Limit(limit).
		Find(&questions).Error
	return questions, err
}

func (r *QuizRepository) SaveAttempt(attempt *model.QuizAttempt) error {
	return r.db.Create(attempt).Error
}

func (r *QuizRepository) UpdateMastery(userID, topicID uuid.UUID, score float64) error {
	var mastery model.MasteryScore
	err := r.db.Where("user_id = ? AND topic_id = ?", userID, topicID).First(&mastery).Error
	now := time.Now()

	if err != nil {
		mastery = model.MasteryScore{
			UserID:          userID,
			TopicID:         topicID,
			MasteryPct:      score,
			TotalAttempts:   1,
			CorrectAttempts: int(score / 100.0 * 5),
			LastPracticedAt: now,
		}
		return r.db.Create(&mastery).Error
	}

	// Exponential Moving Average or Weighted Average for mastery
	mastery.TotalAttempts++
	mastery.MasteryPct = (mastery.MasteryPct*0.6) + (score*0.4)
	mastery.LastPracticedAt = now
	return r.db.Save(&mastery).Error
}

func (r *QuizRepository) GetUserMastery(userID uuid.UUID) ([]model.MasteryScore, error) {
	var scores []model.MasteryScore
	err := r.db.Preload("Topic").Preload("Topic.Subject").
		Where("user_id = ?", userID).
		Find(&scores).Error
	return scores, err
}

