package app

import (
	"log"
	"os"
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/internal/usecase"
	httpserver "workshop-restful-api-backend/pkg/gin"
	"workshop-restful-api-backend/pkg/postgres"
)

func Run() {
	db := postgres.StartPostgres()
	app := httpserver.Start()

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)
	v1 := rest.NewV1(uc)

	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
