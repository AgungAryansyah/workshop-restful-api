package rest

import "workshop-restful-api-backend/internal/usecase"

type Controller struct {
	usecase *usecase.Usecase
}

func NewController(usecase *usecase.Usecase) *Controller {
	return &Controller{usecase}
}
