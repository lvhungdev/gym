package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api/auth"
)

func main() {
	router := gin.Default()

	router.POST("/auth/sign-in", auth.HandleSignIn)
    router.POST("/auth/sign-up", auth.HandleSignUp)

	router.Run("localhost:8080")
}
