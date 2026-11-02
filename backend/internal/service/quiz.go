package service

import (
	"time"

	"study-smart-backend/internal/model"
	"study-smart-backend/internal/repository"

	"github.com/google/uuid"
)

type QuizService struct {
	quizRepo *repository.QuizRepository
	gameRepo *repository.GamificationRepository
}

func NewQuizService(quizRepo *repository.QuizRepository, gameRepo *repository.GamificationRepository) *QuizService {
	return &QuizService{
		quizRepo: quizRepo,
		gameRepo: gameRepo,
	}
}

func (s *QuizService) GetQuizzes(topicID string) ([]model.Quiz, error) {
	return s.quizRepo.GetQuizzes(topicID)
}

func (s *QuizService) GetQuizDetails(quizID uuid.UUID) (*model.Quiz, []model.Question, error) {
	quiz, err := s.quizRepo.GetQuizByID(quizID)
	if err != nil {
		return nil, nil, err
	}

	var questions []model.Question
	if quiz.TopicID != nil {
		questions, _ = s.quizRepo.GetQuestionsForTopic(*quiz.TopicID, quiz.QuestionCount)
	}
	if len(questions) == 0 {
		questions, _ = s.quizRepo.GetAllQuestions(quiz.QuestionCount)
	}

	return quiz, questions, nil
}

type QuizSubmitRequest struct {
	QuizID       uuid.UUID         `json:"quizId"`
	TimeTakenSec int               `json:"timeTakenSec"`
	Answers      map[string]string `json:"answers"` // questionId -> selected answer
}

type QuizSubmitResult struct {
	Score          float64 `json:"score"`
	CorrectCount   int     `json:"correctCount"`
	TotalQuestions int     `json:"totalQuestions"`
	XPEarned       int     `json:"xpEarned"`
	LeveledUp      bool    `json:"leveledUp"`
	Feedback       string  `json:"feedback"`
}

func (s *QuizService) SubmitQuiz(userID uuid.UUID, req QuizSubmitRequest) (*QuizSubmitResult, error) {
	quiz, questions, err := s.GetQuizDetails(req.QuizID)
	if err != nil {
		return nil, err
	}

	correctCount := 0
	total := len(questions)
	if total == 0 {
		total = 1
	}

	for _, q := range questions {
		userAns := req.Answers[q.ID.String()]
		if userAns != "" && userAns == q.CorrectAnswer {
			correctCount++
		}
	}

	score := (float64(correctCount) / float64(total)) * 100.0
	now := time.Now()

	attempt := model.QuizAttempt{
		UserID:         userID,
		QuizID:         req.QuizID,
		Score:          score,
		CorrectCount:   correctCount,
		TotalQuestions: total,
		TimeTakenSec:   req.TimeTakenSec,
		CompletedAt:    &now,
	}
	s.quizRepo.SaveAttempt(&attempt)

	// Update topic mastery score
	if quiz.TopicID != nil {
		s.quizRepo.UpdateMastery(userID, *quiz.TopicID, score)
	}

	// Calculate XP Rewards
	xpEarned := 10 // Base XP for completing
	feedback := "Good effort! Keep reviewing the concepts."
	if score == 100.0 {
		xpEarned += 50
		feedback = "Outstanding! Perfect score! +50 Bonus XP awarded."
	} else if score >= 80.0 {
		xpEarned += 20
		feedback = "Great job! Mastery threshold reached. +20 Bonus XP awarded."
	} else if score < 50.0 {
		feedback = "Needs reinforcement. We have added review cards to your Mind-Feed."
	}

	_, leveledUp, _ := s.gameRepo.AddXP(userID, xpEarned)
	s.gameRepo.UpdateStreakActivity(userID)

	return &QuizSubmitResult{
		Score:          score,
		CorrectCount:   correctCount,
		TotalQuestions: total,
		XPEarned:       xpEarned,
		LeveledUp:      leveledUp,
		Feedback:       feedback,
	}, nil
}

