package controllers

import (
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
	helpers.Response(ctx, "Ping Successfully", nil, 200)
}
