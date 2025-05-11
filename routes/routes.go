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
}
