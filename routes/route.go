package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/controllers"
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
	}

	return
}
