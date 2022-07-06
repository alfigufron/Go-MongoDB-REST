package routes

import (
	"github.com/Go-MongoDB-REST/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(route *gin.RouterGroup) {
	bookController := controllers.NewBookController()

	bookRoute := route.Group("/books")
	{
		bookRoute.GET("/", bookController.FindAll)
		bookRoute.POST("/", bookController.Store)
	}
}