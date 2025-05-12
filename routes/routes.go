package routes

import (
	"github.com/Aakarsh-Kamboj/rest-service/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/frameworks", handlers.GetFrameworkSummaries)
	e.GET("/framework-summaries", handlers.GetFrameworkSummaries)
	e.POST("/frameworks", handlers.CreateFramework)
	e.POST("/organizations", handlers.CreateOrganization)
	e.POST("/controls", handlers.CreateControl)
	e.GET("/controls/:id", handlers.GetControlByID)
	e.GET("/controls", handlers.ListControls)
	e.POST("/evidence-tasks", handlers.CreateEvidenceTask)
	e.GET("/evidence-tasks/:id", handlers.GetEvidenceTaskByID)
	e.GET("/evidence-tasks", handlers.ListEvidenceTasks)
	e.PUT("/evidence-tasks/:id", handlers.UpdateEvidenceTask)
	e.GET("/evidence-tasks/summary", handlers.GetEvidenceSummary)

}
