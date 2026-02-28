package app

import (
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/internal/usecase"
	httpserver "workshop-restful-api-backend/pkg/gin"
	"workshop-restful-api-backend/pkg/postgres"
)

func Run() {
	db := postgres.StartPostgres()
	app := httpserver.Start()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	v1 := rest.NewV1(usecase)

	rest.NewRouter(app, v1)
}
