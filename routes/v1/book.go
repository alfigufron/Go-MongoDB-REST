package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/controllers"
)

func BookRoutes(route *gin.Engine) {
	bookController := controllers.NewBookController()

	bookRoute := route.Group("/book")
	{
		bookRoute.GET("/", bookController.FindAll)
	}
}