package controllers

import (
	"net/http"

	"github.com/Go-MongoDB-REST/helpers"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	res := helpers.ResponseBody("success", "Ping Successfully", nil, 200)

	ctx.JSON(http.StatusOK, res)
}
