package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, controller *Controller) {
	api := app.Group("/api/v1")
	{
		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("/", controller.GetRestaurant)
			restaurants.POST("/", controller.CreateRestaurant)
		}
	}

	app.Run()
}
