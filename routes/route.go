package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/controllers"
)

func InitRouter() (engine *gin.Engine) {
	debug := os.Getenv("APP_MODE")

	switch debug {
	case "STAGING":
		gin.SetMode(gin.TestMode)
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	engine = gin.New()

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/ping", controllers.Ping)
	}

	return
}
