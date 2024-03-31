package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api/auth"
)

func main() {
	router := gin.Default()
	router.GET("/", indexHandler)

	router.POST("/auth/sign-in", auth.HandleSignIn)
    router.POST("/auth/sign-up", auth.HandleSignUp)

	router.Run("localhost:8080")
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, world!",
	})
}
