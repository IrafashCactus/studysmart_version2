package handler

import (
	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func NewAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.SendBadRequest(c, "email and password are required")
		return
	}

	resp, err := h.authSvc.Login(input.Email, input.Password)
	if err != nil {
		pkg.SendUnauthorized(c, err.Error())
		return
	}

	pkg.SendSuccess(c, resp, "logged in successfully")
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
	Role     string `json:"role"`
	Language string `json:"language"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.SendBadRequest(c, "valid email, password (min 6 chars), and full name required")
		return
	}

	resp, err := h.authSvc.Register(input.Email, input.Password, input.FullName, input.Role, input.Language)
	if err != nil {
		pkg.SendBadRequest(c, err.Error())
		return
	}

	pkg.SendCreated(c, resp, "registered successfully")
}

func (h *AuthHandler) Me(c *gin.Context) {
	val, exists := c.Get("userId")
	if !exists {
		pkg.SendUnauthorized(c, "not authenticated")
		return
	}

	userID := val.(uuid.UUID)
	user, err := h.authSvc.GetUser(userID)
	if err != nil {
		pkg.SendNotFound(c, "user not found")
		return
	}

	pkg.SendSuccess(c, user, "user profile")
}

func (h *AuthHandler) DemoLogin(c *gin.Context) {
	role := c.Query("role")
	if role == "" {
		role = "student"
	}

	resp, err := h.authSvc.DemoLogin(role)
	if err != nil {
		pkg.SendServerError(c, "failed to authenticate demo account")
		return
	}

	pkg.SendSuccess(c, resp, "demo account authenticated")
}

