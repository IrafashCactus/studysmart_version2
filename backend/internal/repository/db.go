package repository

import (
	"context"
	"fmt"

	"study-smart-backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresConnection establishes a connection to PostgreSQL using GORM.
func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// HealthRepository provides database health check capabilities.
type HealthRepository struct {
	db *gorm.DB
}

// NewHealthRepository creates a new HealthRepository.
func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{db: db}
}

// Ping verifies the database connection is alive.
func (r *HealthRepository) Ping(ctx context.Context) error {
	sqlDB, err := r.db.WithContext(ctx).DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	return sqlDB.Ping()
}
