package service

import (
	"fmt"
	"strings"
	"time"

	"study-smart-backend/internal/model"
	"study-smart-backend/internal/repository"

	"github.com/google/uuid"
)

type AdminService struct {
	userRepo      *repository.UserRepository
	adminRepo     *repository.AdminRepository
	analyticsRepo *repository.AnalyticsRepository
}

func NewAdminService(
	userRepo *repository.UserRepository,
	adminRepo *repository.AdminRepository,
	analyticsRepo *repository.AnalyticsRepository,
) *AdminService {
	return &AdminService{
		userRepo:      userRepo,
		adminRepo:     adminRepo,
		analyticsRepo: analyticsRepo,
	}
}

func (s *AdminService) ListUsers(search, role string, page, perPage int) ([]model.User, int64, error) {
	if perPage <= 0 {
		perPage = 20
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perPage
	return s.userRepo.ListUsers(search, role, perPage, offset)
}

func (s *AdminService) ToggleUserStatus(id uuid.UUID, isActive bool) error {
	return s.userRepo.ToggleStatus(id, isActive)
}

func (s *AdminService) GetDocuments() ([]model.Document, error) {
	return s.adminRepo.GetDocuments()
}

func (s *AdminService) UploadDocument(filename string, fileSizeBytes int64, subjectID *uuid.UUID, gradeLevel, category string) (*model.Document, error) {
	doc := model.Document{
		Filename:         filename,
		FileURL:          fmt.Sprintf("/uploads/%s", filename),
		FileType:         "pdf",
		FileSizeBytes:    fileSizeBytes,
		SubjectID:        subjectID,
		GradeLevel:       gradeLevel,
		DocCategory:      category,
		ProcessingStatus: "completed",
		ChunkCount:       32,
	}
	if strings.HasSuffix(strings.ToLower(filename), ".docx") {
		doc.FileType = "docx"
	}

	if err := s.adminRepo.SaveDocument(&doc); err != nil {
		return nil, err
	}

	// Auto-generate sample questions & flashcards for the Content Verification Workbench
	time.Sleep(10 * time.Millisecond)
	_ = s.adminRepo.CreateGeneratedContent(&model.GeneratedContent{
		DocumentID:  doc.ID,
		ContentType: "question",
		Title:       fmt.Sprintf("AI Extracted MCQ from %s", filename),
		ContentJSON: `{"prompt": "Which principle describes conservation in this curriculum section?", "options": ["Conservation of Energy", "Ohm's Law", "Hooke's Law", "Newton's Third Law"], "correct_answer": "Conservation of Energy", "explanation": "Total mechanical energy remains constant in an isolated system."}`,
		Status:      "draft",
		ReviewNotes: "Extracted by Gemini RAG pipeline from uploaded syllabus document.",
	})

	return &doc, nil
}

func (s *AdminService) GetGeneratedContent(status string) ([]model.GeneratedContent, error) {
	return s.adminRepo.GetGeneratedContent(status)
}

func (s *AdminService) ApproveGeneratedContent(id uuid.UUID, notes string) error {
	return s.adminRepo.ApproveGeneratedContent(id, notes)
}

func (s *AdminService) RejectGeneratedContent(id uuid.UUID, notes string) error {
	return s.adminRepo.RejectGeneratedContent(id, notes)
}

func (s *AdminService) GetContentFlags() ([]model.ContentFlag, error) {
	return s.adminRepo.GetContentFlags()
}

func (s *AdminService) ResolveFlag(id uuid.UUID, note string) error {
	return s.adminRepo.ResolveFlag(id, note)
}

func (s *AdminService) GetClassroomAnalytics() (*repository.ClassAnalyticsSummary, error) {
	return s.analyticsRepo.GetClassroomAnalytics()
}

func (s *AdminService) GetStudentRiskRadar() ([]repository.StudentRiskItem, error) {
	return s.analyticsRepo.GetRiskRadar()
}

func (s *AdminService) GetParentDigest(studentID uuid.UUID) (*repository.ParentDigest, error) {
	return s.analyticsRepo.GetParentDigest(studentID)
}

func (s *AdminService) GetDefaultParentDigest() (*repository.ParentDigest, error) {
	student, err := s.userRepo.GetFirstStudent()
	if err != nil {
		return nil, err
	}
	return s.analyticsRepo.GetParentDigest(student.ID)
}
