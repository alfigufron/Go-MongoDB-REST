package helpers

import (
	"github.com/Go-MongoDB-REST/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(c *gin.Context, m interface{}) int{
	if err := c.ShouldBindQuery(&m); err != nil {
		errs := err.(validator.ValidationErrors)
		out := []string{}
		for _, val := range errs.Translate(utils.Trans) {
			out = append(out, val)
		}
		Response(c, "Invalid Request", out, 400)
		return 0
	}
	return 1
}