package service

import (
	"errors"

	"study-smart-backend/internal/model"
	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	gameRepo *repository.GamificationRepository
}

func NewAuthService(userRepo *repository.UserRepository, gameRepo *repository.GamificationRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		gameRepo: gameRepo,
	}
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func (s *AuthService) Login(email, password string) (*AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is suspended. Please contact school administration")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := pkg.GenerateToken(user.ID, user.Email, user.FullName, user.Role)
	if err != nil {
		return nil, err
	}

	s.userRepo.UpdateLastLogin(user.ID)

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *AuthService) Register(email, password, fullName, role, lang string) (*AuthResponse, error) {
	if existing, _ := s.userRepo.GetByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if role == "" {
		role = "student"
	}
	if lang == "" {
		lang = "en"
	}

	user := model.User{
		Email:        email,
		PasswordHash: string(hashed),
		FullName:     fullName,
		Role:         role,
		LanguagePref: lang,
		IsActive:     true,
		AvatarURL:    "https://api.dicebear.com/7.x/bottts/svg?seed=" + email,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	// Initialize gamification profile if student
	if role == "student" {
		s.gameRepo.GetXP(user.ID)
		s.gameRepo.GetStreak(user.ID)
	}

	token, err := pkg.GenerateToken(user.ID, user.Email, user.FullName, user.Role)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

func (s *AuthService) GetUser(userID uuid.UUID) (*model.User, error) {
	return s.userRepo.GetByID(userID)
}

func (s *AuthService) DemoLogin(role string) (*AuthResponse, error) {
	var email string
	switch role {
	case "student":
		email = "student@studysmart.rw"
	case "teacher":
		email = "teacher@studysmart.rw"
	case "parent":
		email = "parent@studysmart.rw"
	case "admin", "super_admin":
		email = "admin@studysmart.rw"
	default:
		email = "student@studysmart.rw"
	}

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	token, err := pkg.GenerateToken(user.ID, user.Email, user.FullName, user.Role)
	if err != nil {
		return nil, err
	}

	s.userRepo.UpdateLastLogin(user.ID)

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

