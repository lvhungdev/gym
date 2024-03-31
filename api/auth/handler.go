package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api"
	"github.com/lvhungdev/gym/domain"
	"github.com/lvhungdev/gym/infras"
)

func HandleSignIn(c *gin.Context) {
	repo := infras.NewUserRepo()
	passwordHasher := infras.NewPasswordHasher()
	idGenerator := infras.NewIdGenerator()

	useCases := domain.NewAuthUC(repo, passwordHasher, idGenerator)

	dto := signInReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := useCases.SignIn(dto.Email, dto.Password)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, signInResDto{
		Email: user.Email,
		Token: "TODO",
	})
}

func HandleSignUp(c *gin.Context) {
	repo := infras.NewUserRepo()
	passwordHasher := infras.NewPasswordHasher()
	idGenerator := infras.NewIdGenerator()

	useCases := domain.NewAuthUC(repo, passwordHasher, idGenerator)

	dto := signUpReqDto{}
	if err := c.BindJSON(&dto); err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := useCases.SignUp(dto.Email, dto.Password)
	if err != nil {
		api.NewApiError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResDto{
		Email: user.Email,
		Token: "TODO",
	})
}
