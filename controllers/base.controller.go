package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/helpers"
)

func Ping(ctx *gin.Context) {
	helpers.Response(ctx, "Ping Successfully", nil, 200)
}
