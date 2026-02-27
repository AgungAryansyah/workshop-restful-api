package app

import (
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/internal/usecase"
	"workshop-restful-api-backend/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := postgres.StartPostgres()
	app := gin.New()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	controller := rest.NewController(usecase)

	rest.NewRouter(app, controller)
}
