package repository

import (
	"encoding/json"
	"time"

	"study-smart-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) GetDocuments() ([]model.Document, error) {
	var docs []model.Document
	err := r.db.Preload("Subject").Order("created_at DESC").Find(&docs).Error
	return docs, err
}

func (r *AdminRepository) SaveDocument(doc *model.Document) error {
	return r.db.Create(doc).Error
}

func (r *AdminRepository) CreateGeneratedContent(gc *model.GeneratedContent) error {
	return r.db.Create(gc).Error
}

func (r *AdminRepository) GetGeneratedContent(status string) ([]model.GeneratedContent, error) {
	var contents []model.GeneratedContent
	q := r.db.Preload("Document").Preload("Document.Subject")
	if status != "" {
		q = q.Where("status = ?", status)
	}
	err := q.Order("created_at DESC").Find(&contents).Error
	return contents, err
}

func (r *AdminRepository) ApproveGeneratedContent(id uuid.UUID, notes string) error {
	var gc model.GeneratedContent
	if err := r.db.Preload("Document").First(&gc, "id = ?", id).Error; err != nil {
		return err
	}

	gc.Status = "approved"
	gc.ReviewNotes = notes
	if err := r.db.Save(&gc).Error; err != nil {
		return err
	}

	// Publish into live questions or content_cards
	if gc.ContentType == "question" {
		var qData struct {
			Prompt        string   `json:"prompt"`
			Options       []string `json:"options"`
			CorrectAnswer string   `json:"correct_answer"`
			Explanation   string   `json:"explanation"`
		}
		if err := json.Unmarshal([]byte(gc.ContentJSON), &qData); err == nil {
			optsJSON, _ := json.Marshal(qData.Options)
			// Find a topic for this document's subject
			var topic model.Topic
			r.db.Where("subject_id = ?", gc.Document.SubjectID).First(&topic)

			newQ := model.Question{
				TopicID:       topic.ID,
				QuestionType:  "mcq",
				Prompt:        qData.Prompt,
				OptionsJSON:   string(optsJSON),
				CorrectAnswer: qData.CorrectAnswer,
				Explanation:   qData.Explanation,
				Difficulty:    2,
				Status:        "approved",
			}
			r.db.Create(&newQ)
		}
	} else if gc.ContentType == "flashcard" {
		var fData struct {
			Front       string `json:"front"`
			Back        string `json:"back"`
			KeyTakeaway string `json:"keyTakeaway"`
		}
		if err := json.Unmarshal([]byte(gc.ContentJSON), &fData); err == nil {
			var topic model.Topic
			r.db.Where("subject_id = ?", gc.Document.SubjectID).First(&topic)

			newCard := model.ContentCard{
				TopicID:     topic.ID,
				CardType:    "flashcard",
				Title:       gc.Title,
				ContentText: fData.Front + "\n\n" + fData.Back,
				Front:       fData.Front,
				Back:        fData.Back,
				KeyTakeaway: fData.KeyTakeaway,
				Difficulty:  2,
				Status:      "approved",
			}
			r.db.Create(&newCard)
		}
	}

	return nil
}

func (r *AdminRepository) RejectGeneratedContent(id uuid.UUID, notes string) error {
	return r.db.Model(&model.GeneratedContent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":       "rejected",
			"review_notes": notes,
		}).Error
}

func (r *AdminRepository) GetContentFlags() ([]model.ContentFlag, error) {
	var flags []model.ContentFlag
	err := r.db.Preload("Reporter").Order("created_at DESC").Find(&flags).Error
	return flags, err
}

func (r *AdminRepository) CreateContentFlag(flag *model.ContentFlag) error {
	return r.db.Create(flag).Error
}

func (r *AdminRepository) ResolveFlag(id uuid.UUID, resolutionNote string) error {
	return r.db.Model(&model.ContentFlag{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":          "resolved",
			"resolution_note": resolutionNote,
			"updated_at":      time.Now(),
		}).Error
}
