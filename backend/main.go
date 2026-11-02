package main

import (
	"log/slog"
	"os"

	"study-smart-backend/internal/config"
	"study-smart-backend/internal/handler"
	"study-smart-backend/internal/repository"
	"study-smart-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := repository.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Info("database connected successfully")

	// Initialize layers
	healthRepo := repository.NewHealthRepository(db)
	healthSvc := service.NewHealthService(healthRepo)
	healthHandler := handler.NewHealthHandler(healthSvc)

	// Set up Gin router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", healthHandler.Check)
	}

	// Start server
	addr := ":" + cfg.AppPort
	slog.Info("server starting", "address", addr)
	if err := router.Run(addr); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
