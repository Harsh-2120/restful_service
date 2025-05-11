package handlers

import (
	"net/http"

	"github.com/Aakarsh-Kamboj/rest-service/db"
	"github.com/Aakarsh-Kamboj/rest-service/internal/domain"
	"github.com/Aakarsh-Kamboj/rest-service/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type CreateOrganizationInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func CreateOrganization(c echo.Context) error {
	var input CreateOrganizationInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := validate.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	org := domain.Organization{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := db.DB.Create(&org).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create organization"})
	}

	return c.JSON(http.StatusCreated, org)
}

func CreateFramework(c echo.Context) error {
	var framework domain.Framework
	if err := c.Bind(&framework); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	framework.Organization = domain.Organization{}
	if err := db.DB.Create(&framework).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create framework",
		})
	}

	return c.JSON(http.StatusCreated, framework)
}

func GetFrameworkSummaries(c echo.Context) error {
	summaries, err := services.GetFrameworkSummaries()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch frameworks" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, summaries)
}
