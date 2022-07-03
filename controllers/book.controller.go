package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexa/go-boilerplate-restapi/helpers"
)

type BookController interface {
	FindAll(ctx *gin.Context)
}

type BookControllerImp struct {

}

func NewBookController() BookController {
	return &BookControllerImp{}
}

func (controller BookControllerImp) FindAll(ctx *gin.Context) {
	res := helpers.ResponseBody("success", "Ping Successfully", nil, 200)

	ctx.JSON(http.StatusOK, res)
}
