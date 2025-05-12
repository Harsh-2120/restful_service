package handlers

import (
	"net/http"
	"time"

	"github.com/Aakarsh-Kamboj/rest-service/db"
	"github.com/Aakarsh-Kamboj/rest-service/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var evidenceValidate = validator.New()

type CreateEvidenceInput struct {
	EvidenceName   string    `json:"evidence_name" validate:"required"`
	Status         string    `json:"status" validate:"required,oneof='Pending' 'In Progress' 'Completed' 'Reviewed'"`
	Assignee       string    `json:"assignee"`
	Department     string    `json:"department"`
	DueDate        time.Time `json:"due_date" validate:"required"`
	UploadedDate   time.Time `json:"uploaded_date" validate:"required"`
	FrameworkID    string    `json:"framework_id" validate:"required"`
	OrganizationID string    `json:"organization_id" validate:"required"`
}

type UpdateEvidenceInput struct {
	EvidenceName string    `json:"evidence_name"`
	Status       string    `json:"status" validate:"omitempty,oneof='Pending' 'In Progress' 'Completed' 'Reviewed'"`
	Assignee     string    `json:"assignee"`
	Department   string    `json:"department"`
	DueDate      time.Time `json:"due_date"`
	UploadedDate time.Time `json:"uploaded_date"`
}

func CreateEvidenceTask(c echo.Context) error {
	var input CreateEvidenceInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := evidenceValidate.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	evidence := domain.EvidenceTask{
		EvidenceName:   input.EvidenceName,
		Status:         input.Status,
		Assignee:       input.Assignee,
		Department:     input.Department,
		DueDate:        input.DueDate,
		UploadedDate:   input.UploadedDate,
		FrameworkID:    input.FrameworkID,
		OrganizationID: input.OrganizationID,
	}

	if err := db.DB.Create(&evidence).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create evidence task"})
	}

	return c.JSON(http.StatusCreated, evidence)
}

func GetEvidenceTaskByID(c echo.Context) error {
	id := c.Param("id")

	var evidence domain.EvidenceTask
	if err := db.DB.Preload("Framework").Preload("Organization").First(&evidence, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Evidence task not found"})
	}

	return c.JSON(http.StatusOK, evidence)
}

func ListEvidenceTasks(c echo.Context) error {
	var evidences []domain.EvidenceTask
	if err := db.DB.Preload("Framework").Preload("Organization").Find(&evidences).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch evidence tasks"})
	}
	return c.JSON(http.StatusOK, evidences)
}

func UpdateEvidenceTask(c echo.Context) error {
	id := c.Param("id")
	var input UpdateEvidenceInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := evidenceValidate.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var evidence domain.EvidenceTask
	if err := db.DB.First(&evidence, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Evidence task not found"})
	}

	db.DB.Model(&evidence).Updates(domain.EvidenceTask{
		EvidenceName: input.EvidenceName,
		Status:       input.Status,
		Assignee:     input.Assignee,
		Department:   input.Department,
		DueDate:      input.DueDate,
		UploadedDate: input.UploadedDate,
	})

	return c.JSON(http.StatusOK, evidence)
}
