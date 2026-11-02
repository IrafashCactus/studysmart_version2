package repository

import (
	"errors"
	"time"

	"study-smart-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateLastLogin(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("last_login_at", &now).Error
}

func (r *UserRepository) ListUsers(search, role string, limit, offset int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	q := r.db.Model(&model.User{})
	if search != "" {
		s := "%" + search + "%"
		q = q.Where("full_name ILIKE ? OR email ILIKE ?", s, s)
	}
	if role != "" {
		q = q.Where("role = ?", role)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := q.Order("created_at DESC").Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err
}

func (r *UserRepository) ToggleStatus(id uuid.UUID, isActive bool) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("is_active", isActive).Error
}

func (r *UserRepository) GetStudents() ([]model.User, error) {
	var students []model.User
	err := r.db.Where("role = ?", "student").Find(&students).Error
	return students, err
}

func (r *UserRepository) GetFirstStudent() (*model.User, error) {
	var student model.User
	err := r.db.Where("role = ?", "student").First(&student).Error
	if err != nil {
		return nil, errors.New("no student found")
	}
	return &student, nil
}

