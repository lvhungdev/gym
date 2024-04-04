package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lvhungdev/gym/api"
	"github.com/lvhungdev/gym/domain/entity"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	authService := api.GetDi(c).ResolveAuthService()

	tokenString := c.GetHeader("Authorization")
	if len(tokenString) < 7 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	tokenString = tokenString[7:]

	token, err := verifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	email, err := claims.GetSubject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	authService.SetSignedInUser(entity.User{
		Id:    "1",
		Email: email,
	})

	c.Next()
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	// TODO move this to env
	secret := []byte("my top secret key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
