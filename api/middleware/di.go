package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api"
)

func DIMiddleware(c *gin.Context) {
	di := api.NewScopedDI()
	c.Set("di", di)
}
