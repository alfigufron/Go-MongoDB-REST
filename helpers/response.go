package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseMeta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

func Response(c *gin.Context, message string, data interface{}, code int) {
	var status string

	slice := strconv.Itoa(code)

	switch slice[:1] {
	case "2":
		status = "success"
	case "3":
		status = "redirected"
	case "4":
		status = "client error"
	case "5":
		status = "server error"
	default:
		status = "unknown"
	}

	body := &APIResponse{
		Meta: ResponseMeta{
			Status:  status,
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	c.JSON(code, body)
}
