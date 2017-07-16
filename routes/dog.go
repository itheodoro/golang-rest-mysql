package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itheodoro/rest-mysql/controllers"
)

// InitRoutes init all routes
func InitRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("api/dog")
	{
		api.POST("/", controllers.CreateDog)
		api.GET("/", controllers.GetAllDogs)
	}

	return r
}
