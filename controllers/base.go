package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/helpers"
)

func Ping(c *gin.Context) {
	payload := helpers.ResponseBody("success", "Ping Successfully", nil, 200)

	c.JSON(http.StatusOK, payload)
}
