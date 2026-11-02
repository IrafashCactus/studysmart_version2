package repository

import (
	"time"

	"study-smart-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *FeedRepository {
	return &FeedRepository{db: db}
}

func (r *FeedRepository) GetCards(subjectCode, topicID string, limit, offset int) ([]model.ContentCard, int64, error) {
	var cards []model.ContentCard
	var total int64

	q := r.db.Model(&model.ContentCard{}).
		Preload("Topic").
		Preload("Topic.Subject").
		Where("status = ?", "approved")

	if topicID != "" {
		if uid, err := uuid.Parse(topicID); err == nil {
			q = q.Where("topic_id = ?", uid)
		}
	} else if subjectCode != "" && subjectCode != "ALL" {
		q = q.Joins("JOIN topics ON topics.id = content_cards.topic_id").
			Joins("JOIN subjects ON subjects.id = topics.subject_id").
			Where("subjects.code = ?", subjectCode)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := q.Order("content_cards.created_at DESC").Limit(limit).Offset(offset).Find(&cards).Error
	return cards, total, err
}

func (r *FeedRepository) GetCardByID(id uuid.UUID) (*model.ContentCard, error) {
	var card model.ContentCard
	err := r.db.Preload("Topic").Preload("Topic.Subject").First(&card, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *FeedRepository) RecordInteraction(interaction *model.CardInteraction) error {
	return r.db.Create(interaction).Error
}

func (r *FeedRepository) GetBookmarks(userID uuid.UUID) ([]model.ContentCard, error) {
	var bookmarks []model.Bookmark
	err := r.db.Preload("Card").
		Preload("Card.Topic").
		Preload("Card.Topic.Subject").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&bookmarks).Error

	if err != nil {
		return nil, err
	}

	cards := make([]model.ContentCard, len(bookmarks))
	for i, b := range bookmarks {
		cards[i] = b.Card
	}
	return cards, nil
}

func (r *FeedRepository) ToggleBookmark(userID, cardID uuid.UUID) (bool, error) {
	var existing model.Bookmark
	err := r.db.Where("user_id = ? AND card_id = ?", userID, cardID).First(&existing).Error
	if err == nil {
		// Already bookmarked, remove it
		err = r.db.Delete(&existing).Error
		return false, err
	}

	// Create new bookmark
	newBm := model.Bookmark{
		UserID:    userID,
		CardID:    cardID,
		CreatedAt: time.Now(),
	}
	err = r.db.Create(&newBm).Error
	return true, err
}

func (r *FeedRepository) IsBookmarked(userID, cardID uuid.UUID) bool {
	var count int64
	r.db.Model(&model.Bookmark{}).Where("user_id = ? AND card_id = ?", userID, cardID).Count(&count)
	return count > 0
}

func (r *FeedRepository) GetSubjects() ([]model.Subject, error) {
	var subjects []model.Subject
	err := r.db.Order("name ASC").Find(&subjects).Error
	return subjects, err
}

func (r *FeedRepository) GetTopicsBySubject(subjectID uuid.UUID) ([]model.Topic, error) {
	var topics []model.Topic
	err := r.db.Where("subject_id = ?", subjectID).Order("unit_number ASC").Find(&topics).Error
	return topics, err
}

