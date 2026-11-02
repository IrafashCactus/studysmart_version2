package handler

import (
	"strconv"

	"study-smart-backend/internal/pkg"
	"study-smart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminHandler struct {
	adminSvc *service.AdminService
}

func NewAdminHandler(adminSvc *service.AdminService) *AdminHandler {
	return &AdminHandler{adminSvc: adminSvc}
}

// User Management
func (h *AdminHandler) GetUsers(c *gin.Context) {
	search := c.Query("search")
	role := c.Query("role")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "20"))

	users, total, err := h.adminSvc.ListUsers(search, role, page, perPage)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{
		"users":   users,
		"total":   total,
		"page":    page,
		"perPage": perPage,
	}, "user directory")
}

func (h *AdminHandler) ToggleUserStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid user id")
		return
	}

	var body struct {
		IsActive bool `json:"isActive"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		pkg.SendBadRequest(c, "isActive boolean required")
		return
	}

	if err := h.adminSvc.ToggleUserStatus(id, body.IsActive); err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	status := "activated"
	if !body.IsActive {
		status = "suspended"
	}
	pkg.SendSuccess(c, gin.H{"id": id, "isActive": body.IsActive}, "user account "+status)
}

// RAG Documents
func (h *AdminHandler) GetDocuments(c *gin.Context) {
	docs, err := h.adminSvc.GetDocuments()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, docs, "curriculum documents")
}

type UploadDocInput struct {
	Filename      string `json:"filename" binding:"required"`
	FileSizeBytes int64  `json:"fileSizeBytes"`
	GradeLevel    string `json:"gradeLevel"`
	Category      string `json:"category"`
	SubjectID     string `json:"subjectId"`
}

func (h *AdminHandler) UploadDocument(c *gin.Context) {
	var input UploadDocInput
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.SendBadRequest(c, "filename required")
		return
	}

	var subUID *uuid.UUID
	if input.SubjectID != "" {
		if u, err := uuid.Parse(input.SubjectID); err == nil {
			subUID = &u
		}
	}

	doc, err := h.adminSvc.UploadDocument(input.Filename, input.FileSizeBytes, subUID, input.GradeLevel, input.Category)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendCreated(c, doc, "document ingested and queued for RAG extraction")
}

// Content Verification Workbench
func (h *AdminHandler) GetGeneratedContent(c *gin.Context) {
	status := c.Query("status")
	content, err := h.adminSvc.GetGeneratedContent(status)
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, content, "generated content queue")
}

type ReviewContentInput struct {
	Notes string `json:"notes"`
}

func (h *AdminHandler) ApproveContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid content id")
		return
	}

	var input ReviewContentInput
	c.ShouldBindJSON(&input)

	if err := h.adminSvc.ApproveGeneratedContent(id, input.Notes); err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{"id": id, "status": "approved"}, "content verified and published to curriculum bank")
}

func (h *AdminHandler) RejectContent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid content id")
		return
	}

	var input ReviewContentInput
	c.ShouldBindJSON(&input)

	if err := h.adminSvc.RejectGeneratedContent(id, input.Notes); err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{"id": id, "status": "rejected"}, "content marked as rejected")
}

// Moderation Flags
func (h *AdminHandler) GetFlags(c *gin.Context) {
	flags, err := h.adminSvc.GetContentFlags()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, flags, "moderation flag queue")
}

func (h *AdminHandler) ResolveFlag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		pkg.SendBadRequest(c, "invalid flag id")
		return
	}

	var input struct {
		ResolutionNote string `json:"resolutionNote"`
	}
	c.ShouldBindJSON(&input)

	if err := h.adminSvc.ResolveFlag(id, input.ResolutionNote); err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}

	pkg.SendSuccess(c, gin.H{"id": id, "status": "resolved"}, "moderation flag resolved")
}

// Educator Endpoints
func (h *AdminHandler) GetEducatorAnalytics(c *gin.Context) {
	stats, err := h.adminSvc.GetClassroomAnalytics()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, stats, "classroom analytics overview")
}

func (h *AdminHandler) GetRiskRadar(c *gin.Context) {
	radar, err := h.adminSvc.GetStudentRiskRadar()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, radar, "student risk radar alerts")
}

// Parent Endpoints
func (h *AdminHandler) GetParentDigest(c *gin.Context) {
	studentIDStr := c.Query("studentId")
	if studentIDStr != "" {
		if studentID, err := uuid.Parse(studentIDStr); err == nil {
			digest, err := h.adminSvc.GetParentDigest(studentID)
			if err == nil {
				pkg.SendSuccess(c, digest, "weekly parent accountability digest")
				return
			}
		}
	}

	// Fallback to default student demo digest
	digest, err := h.adminSvc.GetDefaultParentDigest()
	if err != nil {
		pkg.SendServerError(c, err.Error())
		return
	}
	pkg.SendSuccess(c, digest, "weekly parent accountability digest")
}

