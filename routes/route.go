package routes

import (
	"os"

	"github.com/Go-MongoDB-REST/controllers"
	"github.com/Go-MongoDB-REST/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() (route *gin.Engine) {
	debug := os.Getenv("APP_MODE")

	switch debug {
	case "STAGING":
		gin.SetMode(gin.TestMode)
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	route = gin.New()

	apiGroup := route.Group("/api")
	{
		apiGroup.GET("/ping", controllers.Ping)

		v1Group:=apiGroup.Group("/v1")
		{
			routes.BookRoutes(v1Group)
		}
	}

	return
}
