package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lvhungdev/gym/api"
	"github.com/lvhungdev/gym/domain/entity"
	"net/http"
	"time"
)

func HandleSignIn(c *gin.Context) {
	uc := api.GetDi(c).ResolveAuthUC()

	dto := signInReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	user, err := uc.SignIn(dto.Email, dto.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, signInResDto{
		Email: user.Email,
		Token: token,
	})
}

func HandleSignUp(c *gin.Context) {
	uc := api.GetDi(c).ResolveAuthUC()

	dto := signUpReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	user, err := uc.SignUp(dto.Email, dto.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, signUpResDto{
		Email: user.Email,
		Token: token,
	})
}

func generateJWT(user entity.User) (string, error) {
	// TODO move this to env
	secret := []byte("my top secret key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "", // TODO
		Subject:   user.Email,
		Audience:  nil,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        user.Id,
	})

	return claims.SignedString(secret)
}
