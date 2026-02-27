package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetRestaurant(ctx *gin.Context) {
	restaurants, err := c.usecase.RestaurantUsecase.GetRestaurants()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, restaurants)
}
