package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/helpers"
)

func Ping(ctx *gin.Context) {
	res := helpers.ResponseBody("success", "Ping Successfully", nil, 200)

	ctx.JSON(http.StatusOK, res)
}
