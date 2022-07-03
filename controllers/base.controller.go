package controllers

import (
	"github.com/Go-MongoDB-REST/helpers"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	helpers.Response(ctx, "Ping Successfully", nil, 200)
}
