package api

import (
	"github.com/gin-gonic/gin"
)

func NewApiError(c *gin.Context, httpStatusCode int, msg string) {
	err := errorDto{
		Message: msg,
	}

	c.JSON(httpStatusCode, err)
}

type errorDto struct {
	Message string `json:"message"`
}
