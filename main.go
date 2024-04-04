package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api/auth"
	"github.com/lvhungdev/gym/api/class"
	"github.com/lvhungdev/gym/api/middleware"
)

func main() {
	router := gin.Default()

	router.Use(middleware.DIMiddleware)

	router.POST("/auth/sign-in", auth.HandleSignIn)
	router.POST("/auth/sign-up", auth.HandleSignUp)

	router.GET("/class", middleware.AuthMiddleware, class.HandleGetClasses)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
