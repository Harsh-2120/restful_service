package handlers

import (
	"net/http"

	"github.com/Aakarsh-Kamboj/rest-service/db"
	"github.com/Aakarsh-Kamboj/rest-service/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var controlValidate = validator.New()

type CreateControlInput struct {
	ControlCode     string `json:"control_code" validate:"required"`
	ControlName     string `json:"control_name" validate:"required"`
	ControlDomain   string `json:"control_domain"`
	Status          string `json:"status"`
	Assignee        string `json:"assignee"`
	Description     string `json:"description"`
	ControlQuestion string `json:"control_question"`
	OrganizationID  string `json:"organization_id" validate:"required"`
}

func CreateControl(c echo.Context) error {
	var input CreateControlInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := controlValidate.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	control := domain.Control{
		ControlCode:     input.ControlCode,
		ControlName:     input.ControlName,
		ControlDomain:   input.ControlDomain,
		Status:          input.Status,
		Assignee:        input.Assignee,
		Description:     input.Description,
		ControlQuestion: input.ControlQuestion,
		OrganizationID:  input.OrganizationID,
	}

	if err := db.DB.Create(&control).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create control"})
	}

	return c.JSON(http.StatusCreated, control)
}

func GetControlByID(c echo.Context) error {
	id := c.Param("id")

	var control domain.Control
	if err := db.DB.Preload("Organization").First(&control, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Control not found"})
	}

	return c.JSON(http.StatusOK, control)
}

func ListControls(c echo.Context) error {
	var controls []domain.Control
	if err := db.DB.Preload("Organization").Find(&controls).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch controls"})
	}

	return c.JSON(http.StatusOK, controls)
}
