package domain

import (
	"errors"

	"github.com/lvhungdev/gym/domain/entity"
	"github.com/lvhungdev/gym/domain/port"
)

type AuthUC struct {
	repo        port.UserRepo
	hasher      port.PasswordHasher
	idGenerator port.IdGenerator
}

func NewAuthUC(repo port.UserRepo, hasher port.PasswordHasher, idGenerator port.IdGenerator) AuthUC {
	return AuthUC{
		repo:        repo,
		hasher:      hasher,
		idGenerator: idGenerator,
	}
}

func (u AuthUC) SignIn(email string, password string) (entity.User, error) {
	user, found := u.repo.GetUserByEmail(email)
	if !found {
		return entity.User{}, errors.New("credentials are incorrect")
	}

	if !u.hasher.VerifyPassword(password, user.Password) {
		return entity.User{}, errors.New("credentials are incorrect")
	}

	return user, nil
}

func (u AuthUC) SignUp(email string, password string) (entity.User, error) {
	_, found := u.repo.GetUserByEmail(email)
	if found {
		return entity.User{}, errors.New("email already exists")
	}

	user := entity.User{
		Id:       u.idGenerator.Generate(),
		Email:    email,
		Password: u.hasher.HashPassword(password),
		FullName: "",
	}

	err := u.repo.CreateNewUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
