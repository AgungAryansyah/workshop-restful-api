package app

import (
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/pkg/postgres"
)

func Run() {
	db := postgres.StartPostgres()

	_ = repository.NewRepository(db)
}
