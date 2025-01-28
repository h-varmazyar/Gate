package httpext

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data      any    `json:"data"`
	Error     string `json:"error"`
	IsSuccess bool   `json:"is_success"`
}

func SendGinError(c *gin.Context, err error) {
	resp := Response{
		Error:     err.Error(),
		IsSuccess: false,
	}

	c.JSON(400, resp)
}

func SendGinModel(c *gin.Context, code int, model any) {
	resp := Response{
		Data:      model,
		IsSuccess: true,
	}

	c.JSON(code, resp)
}

func SendGinData(c *gin.Context, code int, mime string, data []byte) {}

func SendGinCode(c *gin.Context, code int) {}
