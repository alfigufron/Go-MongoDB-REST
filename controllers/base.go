package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/helpers"
)

func Ping(c *gin.Context) {
	helpers.Response(c, "Ping Successfully!!", nil, 200)
}
