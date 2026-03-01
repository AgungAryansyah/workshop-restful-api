package rest

import (
	"net/http"
	"workshop-restful-api-backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *V1) GetRestaurantItems(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	items, err := r.usecase.ItemUsecase.GetRestaurantItems(ctx, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (r *V1) CreateItem(c *gin.Context) {
	var create model.CreateItem
	err := c.ShouldBindBodyWithJSON(&create)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	item, err := r.usecase.ItemUsecase.CreateItem(ctx, create)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, item)
}
