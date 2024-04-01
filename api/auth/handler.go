package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lvhungdev/gym/api"
	"github.com/lvhungdev/gym/domain/entity"
	"github.com/lvhungdev/gym/infras"
)

func HandleSignIn(c *gin.Context) {
	uc := infras.ResolveAuthUC()

	dto := signInReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := uc.SignIn(dto.Email, dto.Password)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	jwt, err := generateJWT(user)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, signInResDto{
		Email: user.Email,
		Token: jwt,
	})
}

func HandleSignUp(c *gin.Context) {
	uc := infras.ResolveAuthUC()

	dto := signUpReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := uc.SignUp(dto.Email, dto.Password)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	jwt, err := generateJWT(user)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResDto{
		Email: user.Email,
		Token: jwt,
	})
}

func generateJWT(user entity.User) (string, error) {
	// TODO move this to env
	secret := []byte("my top secret key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		// "aud": getRole(username),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	return claims.SignedString(secret)
}
