package repository

import (
	"context"
	"fmt"
	"log/slog"

	"study-smart-backend/internal/config"
	"study-smart-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresConnection establishes a connection to PostgreSQL and auto-migrates models.
func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate all models
	err = db.AutoMigrate(
		&model.User{},
		&model.School{},
		&model.Classroom{},
		&model.Subject{},
		&model.Topic{},
		&model.ContentCard{},
		&model.Bookmark{},
		&model.CardInteraction{},
		&model.Question{},
		&model.Quiz{},
		&model.QuizAttempt{},
		&model.StudentXP{},
		&model.Streak{},
		&model.Quest{},
		&model.QuestProgress{},
		&model.StudySession{},
		&model.MasteryScore{},
		&model.Document{},
		&model.GeneratedContent{},
		&model.ContentFlag{},
	)
	if err != nil {
		slog.Error("failed to automigrate models", "error", err)
		return nil, fmt.Errorf("automigrate failed: %w", err)
	}
	slog.Info("database schema migrated successfully")

	// Seed sample curriculum and accounts if empty
	if err := SeedDatabase(db); err != nil {
		slog.Warn("database seeding warning", "error", err)
	}

	return db, nil
}

// HealthRepository provides database health check capabilities.
type HealthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{db: db}
}

func (r *HealthRepository) Ping(ctx context.Context) error {
	sqlDB, err := r.db.WithContext(ctx).DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	return sqlDB.Ping()
}
