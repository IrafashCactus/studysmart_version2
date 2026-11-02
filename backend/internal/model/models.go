package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all UUID entities
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

// User represents any system user (student, teacher, parent, admin)
type User struct {
	Base
	Email        string     `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string     `gorm:"not null" json:"-"`
	FullName     string     `gorm:"not null" json:"fullName"`
	AvatarURL    string     `json:"avatarUrl"`
	Phone        string     `json:"phone"`
	LanguagePref string     `gorm:"default:'en'" json:"languagePref"` // 'en', 'fr', 'rw'
	Role         string     `gorm:"not null;default:'student'" json:"role"` // 'student', 'teacher', 'parent', 'school_admin', 'super_admin'
	IsActive     bool       `gorm:"default:true" json:"isActive"`
	LastLoginAt  *time.Time `json:"lastLoginAt"`
}

// School represents an educational institution
type School struct {
	Base
	Name       string `gorm:"not null" json:"name"`
	Code       string `gorm:"uniqueIndex" json:"code"`
	District   string `json:"district"`
	Sector     string `json:"sector"`
	Province   string `json:"province"`
	SchoolType string `json:"schoolType"` // 'primary', 'secondary', 'tvet'
}

// Classroom represents a class/grade stream
type Classroom struct {
	Base
	SchoolID     uuid.UUID `gorm:"type:uuid;not null;index" json:"schoolId"`
	School       School    `gorm:"foreignKey:SchoolID" json:"school,omitempty"`
	GradeLevel   string    `gorm:"not null" json:"gradeLevel"` // 'S1', 'S2', 'S3', 'S4', 'S5', 'S6'
	Stream       string    `json:"stream"`                     // 'MCB', 'PCB', 'MPC', 'General'
	AcademicYear int       `json:"academicYear"`
}

// Subject represents a curriculum subject
type Subject struct {
	Base
	Name        string `gorm:"not null" json:"name"`
	NameRW      string `json:"nameRw"`
	NameFR      string `json:"nameFr"`
	Code        string `gorm:"uniqueIndex;not null" json:"code"` // 'MATH', 'PHY', 'CHEM', 'BIO', 'ENT'
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	GradeLevels string `json:"gradeLevels"` // comma separated e.g. "S1,S2,S3,S4,S5,S6"
}

// Topic represents a specific unit in a subject
type Topic struct {
	Base
	SubjectID      uuid.UUID `gorm:"type:uuid;not null;index" json:"subjectId"`
	Subject        Subject   `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	GradeLevel     string    `gorm:"not null" json:"gradeLevel"`
	UnitNumber     int       `json:"unitNumber"`
	Title          string    `gorm:"not null" json:"title"`
	TitleRW        string    `json:"titleRw"`
	TitleFR        string    `json:"titleFr"`
	Description    string    `json:"description"`
	CompetencyTags string    `json:"competencyTags"` // comma separated REB codes
	SortOrder      int       `gorm:"default:0" json:"sortOrder"`
}

// ContentCard represents a bite-sized piece of learning for the Mind-Feed
type ContentCard struct {
	Base
	TopicID     uuid.UUID `gorm:"type:uuid;not null;index" json:"topicId"`
	Topic       Topic     `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
	CardType    string    `gorm:"not null;default:'micro_concept'" json:"cardType"` // 'flashcard', 'micro_concept', 'summary', 'video_30s'
	Title       string    `gorm:"not null" json:"title"`
	ContentText string    `gorm:"type:text;not null" json:"contentText"`
	Front       string    `gorm:"type:text" json:"front,omitempty"`
	Back        string    `gorm:"type:text" json:"back,omitempty"`
	KeyTakeaway string    `json:"keyTakeaway,omitempty"`
	MediaURL    string    `json:"mediaUrl,omitempty"`
	Difficulty  int       `gorm:"default:1" json:"difficulty"`
	Language    string    `gorm:"default:'en'" json:"language"`
	Status      string    `gorm:"default:'approved'" json:"status"` // 'draft', 'review', 'approved'
	Version     int64     `gorm:"default:1" json:"version"`

	// Embedded Interactive Check (prompt within the feed card)
	QuestionPrompt string `json:"questionPrompt,omitempty"`
	QuestionType   string `json:"questionType,omitempty"` // 'single_choice', 'multi_choice', 'matching'
	OptionsJSON    string `gorm:"type:text" json:"optionsJson,omitempty"` // JSON string array of choices
	CorrectAnswer  string `json:"correctAnswer,omitempty"`
	Explanation    string `json:"explanation,omitempty"`
}

// Bookmark allows students to save difficult cards
type Bookmark struct {
	UserID    uuid.UUID   `gorm:"type:uuid;primary_key" json:"userId"`
	CardID    uuid.UUID   `gorm:"type:uuid;primary_key" json:"cardId"`
	Card      ContentCard `gorm:"foreignKey:CardID" json:"card,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
}

// CardInteraction tracks engagement telemetry
type CardInteraction struct {
	Base
	UserID      uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	CardID      uuid.UUID `gorm:"type:uuid;not null;index" json:"cardId"`
	Action      string    `gorm:"not null" json:"action"` // 'view', 'complete', 'bookmark', 'skip'
	DurationSec int       `json:"durationSec"`
}

// Question is a reusable assessment item
type Question struct {
	Base
	TopicID       uuid.UUID `gorm:"type:uuid;not null;index" json:"topicId"`
	Topic         Topic     `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
	QuestionType  string    `gorm:"not null;default:'mcq'" json:"questionType"` // 'mcq', 'short_answer', 'fill_blank', 'matching'
	Prompt        string    `gorm:"type:text;not null" json:"prompt"`
	PromptRW      string    `gorm:"type:text" json:"promptRw"`
	PromptFR      string    `gorm:"type:text" json:"promptFr"`
	OptionsJSON   string    `gorm:"type:text" json:"optionsJson"` // JSON encoded choices
	CorrectAnswer string    `gorm:"type:text;not null" json:"correctAnswer"`
	Explanation   string    `gorm:"type:text" json:"explanation"`
	Difficulty    int       `gorm:"default:1" json:"difficulty"`
	Status        string    `gorm:"default:'approved'" json:"status"`
}

// Quiz groups questions into an assessment
type Quiz struct {
	Base
	Title          string    `gorm:"not null" json:"title"`
	TopicID        *uuid.UUID `gorm:"type:uuid;index" json:"topicId,omitempty"`
	Topic          *Topic    `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
	QuizType       string    `gorm:"default:'practice'" json:"quizType"` // 'practice', 'daily_challenge', 'homework', 'exam_prep'
	TimeLimitSec   int       `gorm:"default:300" json:"timeLimitSec"`
	QuestionCount  int       `gorm:"default:5" json:"questionCount"`
	ClassroomID    *uuid.UUID `gorm:"type:uuid" json:"classroomId,omitempty"`
	IsPublished    bool      `gorm:"default:true" json:"isPublished"`
}

// QuizAttempt represents a completed or ongoing student quiz
type QuizAttempt struct {
	Base
	UserID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"userId"`
	QuizID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"quizId"`
	Quiz           Quiz       `gorm:"foreignKey:QuizID" json:"quiz,omitempty"`
	Score          float64    `json:"score"` // 0-100%
	CorrectCount   int        `json:"correctCount"`
	TotalQuestions int        `json:"totalQuestions"`
	TimeTakenSec   int        `json:"timeTakenSec"`
	CompletedAt    *time.Time `json:"completedAt"`
}

// Gamification Models
type StudentXP struct {
	UserID       uuid.UUID `gorm:"type:uuid;primary_key" json:"userId"`
	TotalXP      int64     `gorm:"default:0" json:"totalXp"`
	CurrentLevel int       `gorm:"default:1" json:"currentLevel"`
	XPToNext     int       `gorm:"default:100" json:"xpToNext"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Streak struct {
	UserID            uuid.UUID  `gorm:"type:uuid;primary_key" json:"userId"`
	CurrentStreak     int        `gorm:"default:1" json:"currentStreak"`
	LongestStreak     int        `gorm:"default:1" json:"longestStreak"`
	LastActivityDate  string     `json:"lastActivityDate"` // YYYY-MM-DD
	FreezePasses      int        `gorm:"default:2" json:"freezePasses"`
	LastFreezeUsedDate string    `json:"lastFreezeUsedDate"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

type Quest struct {
	Base
	QuestType       string `gorm:"not null" json:"questType"` // 'daily', 'weekly'
	Title           string `gorm:"not null" json:"title"`
	Description     string `json:"description"`
	RequirementType string `gorm:"not null" json:"requirementType"` // 'cards_viewed', 'quizzes_completed', 'accuracy_target', 'focus_minutes'
	TargetCount     int    `gorm:"not null" json:"targetCount"`
	XPReward        int    `gorm:"not null" json:"xpReward"`
	SubjectCode     string `json:"subjectCode,omitempty"`
}

type QuestProgress struct {
	Base
	UserID       uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	QuestID      uuid.UUID `gorm:"type:uuid;not null;index" json:"questId"`
	Quest        Quest     `gorm:"foreignKey:QuestID" json:"quest,omitempty"`
	CurrentCount int       `gorm:"default:0" json:"currentCount"`
	IsCompleted  bool      `gorm:"default:false" json:"isCompleted"`
	Claimed      bool      `gorm:"default:false" json:"claimed"`
}

type StudySession struct {
	Base
	UserID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"userId"`
	SessionMode   string     `gorm:"not null" json:"sessionMode"` // 'feed', 'focus', 'quiz', 'tutor'
	DurationSec   int        `gorm:"default:0" json:"durationSec"`
	CardsViewed   int        `gorm:"default:0" json:"cardsViewed"`
	QuizzesTaken  int        `gorm:"default:0" json:"quizzesTaken"`
	IsFocusMode   bool       `gorm:"default:false" json:"isFocusMode"`
	BonusXPEarned int        `gorm:"default:0" json:"bonusXpEarned"`
}

// MasteryScore tracks student proficiency per topic
type MasteryScore struct {
	UserID          uuid.UUID `gorm:"type:uuid;primary_key" json:"userId"`
	TopicID         uuid.UUID `gorm:"type:uuid;primary_key" json:"topicId"`
	Topic           Topic     `gorm:"foreignKey:TopicID" json:"topic,omitempty"`
	MasteryPct      float64   `gorm:"default:0" json:"masteryPct"`
	TotalAttempts   int       `gorm:"default:0" json:"totalAttempts"`
	CorrectAttempts int       `gorm:"default:0" json:"correctAttempts"`
	LastPracticedAt time.Time `json:"lastPracticedAt"`
}

// RAG and Content Verification Models
type Document struct {
	Base
	Filename         string     `gorm:"not null" json:"filename"`
	FileURL          string     `json:"fileUrl"`
	FileType         string     `json:"fileType"` // 'pdf', 'docx'
	FileSizeBytes    int64      `json:"fileSizeBytes"`
	SubjectID        *uuid.UUID `gorm:"type:uuid" json:"subjectId,omitempty"`
	Subject          *Subject   `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	GradeLevel       string     `json:"gradeLevel"`
	DocCategory      string     `json:"docCategory"` // 'textbook', 'past_exam', 'syllabus', 'teacher_notes'
	ProcessingStatus string     `gorm:"default:'completed'" json:"processingStatus"` // 'pending', 'processing', 'completed', 'failed'
	ChunkCount       int        `gorm:"default:0" json:"chunkCount"`
	UploadedBy       *uuid.UUID `gorm:"type:uuid" json:"uploadedBy,omitempty"`
}

type GeneratedContent struct {
	Base
	DocumentID  uuid.UUID `gorm:"type:uuid;not null;index" json:"documentId"`
	Document    Document  `gorm:"foreignKey:DocumentID" json:"document,omitempty"`
	ContentType string    `gorm:"not null" json:"contentType"` // 'question', 'flashcard', 'summary'
	Title       string    `json:"title"`
	ContentJSON string    `gorm:"type:text;not null" json:"contentJson"`
	Status      string    `gorm:"default:'draft'" json:"status"` // 'draft', 'approved', 'rejected'
	ReviewNotes string    `json:"reviewNotes"`
}

type ContentFlag struct {
	Base
	ReporterID     uuid.UUID `gorm:"type:uuid;not null;index" json:"reporterId"`
	Reporter       User      `gorm:"foreignKey:ReporterID" json:"reporter,omitempty"`
	ContentType    string    `gorm:"not null" json:"contentType"` // 'content_card', 'question'
	ContentID      uuid.UUID `gorm:"type:uuid;not null" json:"contentId"`
	Reason         string    `gorm:"not null" json:"reason"` // 'incorrect', 'unclear', 'inappropriate', 'typo'
	Details        string    `json:"details"`
	Status         string    `gorm:"default:'pending'" json:"status"` // 'pending', 'resolved', 'dismissed'
	ResolutionNote string    `json:"resolutionNote"`
}

