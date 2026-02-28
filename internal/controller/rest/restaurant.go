package rest

import (
	"net/http"
	"workshop-restful-api-backend/internal/model"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetRestaurant(ctx *gin.Context) {
	restaurants, err := c.usecase.RestaurantUsecase.GetRestaurants()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, restaurants)
}

func (c *Controller) CreateRestaurant(ctx *gin.Context) {
	var createRestaurant model.CreateRestaurant

	err := ctx.ShouldBindBodyWithJSON(&createRestaurant)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	restaurant, err := c.usecase.RestaurantUsecase.CreateRestaurant(createRestaurant)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusCreated, restaurant)
}
