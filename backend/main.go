package main

import (
	"log/slog"
	"os"

	"study-smart-backend/internal/config"
	"study-smart-backend/internal/handler"
	"study-smart-backend/internal/pkg"
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

	// Connect to database & auto-migrate & seed
	db, err := repository.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Info("database connected and initialized successfully")

	// Initialize Repositories
	healthRepo := repository.NewHealthRepository(db)
	userRepo := repository.NewUserRepository(db)
	feedRepo := repository.NewFeedRepository(db)
	gameRepo := repository.NewGamificationRepository(db)
	quizRepo := repository.NewQuizRepository(db)
	adminRepo := repository.NewAdminRepository(db)
	analyticsRepo := repository.NewAnalyticsRepository(db)

	// Initialize Services
	healthSvc := service.NewHealthService(healthRepo)
	authSvc := service.NewAuthService(userRepo, gameRepo)
	feedSvc := service.NewFeedService(feedRepo, gameRepo)
	gameSvc := service.NewGamificationService(gameRepo)
	quizSvc := service.NewQuizService(quizRepo, gameRepo)
	tutorSvc := service.NewTutorService(quizRepo, feedRepo)
	adminSvc := service.NewAdminService(userRepo, adminRepo, analyticsRepo)

	// Initialize Handlers
	healthHandler := handler.NewHealthHandler(healthSvc)
	authHandler := handler.NewAuthHandler(authSvc)
	feedHandler := handler.NewFeedHandler(feedSvc)
	gameHandler := handler.NewGamificationHandler(gameSvc)
	quizHandler := handler.NewQuizHandler(quizSvc)
	tutorHandler := handler.NewTutorHandler(tutorSvc)
	adminHandler := handler.NewAdminHandler(adminSvc)

	// Gin Engine
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.CORSOrigin, "http://localhost:5173", "http://127.0.0.1:5173", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Health
		v1.GET("/health", healthHandler.Check)

		// Auth
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/demo", authHandler.DemoLogin)
			authGroup.GET("/me", pkg.AuthMiddleware(), authHandler.Me)
		}

		// Feed (Infinite Student Mind-Feed)
		feedGroup := v1.Group("/feed")
		{
			feedGroup.GET("", pkg.OptionalAuthMiddleware(), feedHandler.GetFeed)
			feedGroup.GET("/subjects", feedHandler.GetSubjects)
			feedGroup.POST("/interact", pkg.AuthMiddleware(), feedHandler.Interact)
			feedGroup.GET("/bookmarks", pkg.AuthMiddleware(), feedHandler.GetBookmarks)
			feedGroup.POST("/bookmarks/:cardId", pkg.AuthMiddleware(), feedHandler.ToggleBookmark)
		}

		// Gamification Engine
		gameGroup := v1.Group("/gamification")
		{
			gameGroup.GET("/leaderboard", gameHandler.GetLeaderboard)
			gameGroup.GET("/profile", pkg.AuthMiddleware(), gameHandler.GetProfile)
			gameGroup.POST("/streak/freeze", pkg.AuthMiddleware(), gameHandler.FreezeStreak)
			gameGroup.POST("/quests/:questId/claim", pkg.AuthMiddleware(), gameHandler.ClaimQuest)
			gameGroup.POST("/focus", pkg.AuthMiddleware(), gameHandler.CompleteFocusSession)
		}

		// Quizzes & Assessment
		quizGroup := v1.Group("/quizzes")
		{
			quizGroup.GET("", quizHandler.GetQuizzes)
			quizGroup.GET("/:id", quizHandler.GetQuiz)
			quizGroup.POST("/submit", pkg.AuthMiddleware(), quizHandler.SubmitQuiz)
		}

		// AI Tutor & Problem Explainer
		tutorGroup := v1.Group("/tutor")
		{
			tutorGroup.POST("/explain", tutorHandler.ExplainProblem)
			tutorGroup.GET("/adaptive", pkg.AuthMiddleware(), tutorHandler.GetAdaptiveRecommendations)
		}

		// Educator Module
		educatorGroup := v1.Group("/educator")
		{
			educatorGroup.GET("/analytics", adminHandler.GetEducatorAnalytics)
			educatorGroup.GET("/risk-radar", adminHandler.GetRiskRadar)
		}

		// Parent Module
		parentGroup := v1.Group("/parent")
		{
			parentGroup.GET("/digest", adminHandler.GetParentDigest)
		}

		// Admin & RAG Workbench
		adminGroup := v1.Group("/admin")
		{
			adminGroup.GET("/users", adminHandler.GetUsers)
			adminGroup.PATCH("/users/:id/status", adminHandler.ToggleUserStatus)
			adminGroup.GET("/documents", adminHandler.GetDocuments)
			adminGroup.POST("/documents", adminHandler.UploadDocument)
			adminGroup.GET("/generated", adminHandler.GetGeneratedContent)
			adminGroup.POST("/generated/:id/approve", adminHandler.ApproveContent)
			adminGroup.POST("/generated/:id/reject", adminHandler.RejectContent)
			adminGroup.GET("/flags", adminHandler.GetFlags)
			adminGroup.PATCH("/flags/:id", adminHandler.ResolveFlag)
		}
	}

	// Start server
	addr := ":" + cfg.AppPort
	slog.Info("server starting", "address", addr)
	if err := router.Run(addr); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
