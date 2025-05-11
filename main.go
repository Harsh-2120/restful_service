package main

import (
	"github.com/Aakarsh-Kamboj/rest-service/db"
	"github.com/Aakarsh-Kamboj/rest-service/internal/domain"
	"github.com/Aakarsh-Kamboj/rest-service/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(
		&domain.Organization{},
		&domain.Framework{},
		&domain.Control{},
		&domain.FrameworkControl{},
		&domain.EvidenceTask{},
		&domain.Department{},
		&domain.Policy{},
	)
	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
