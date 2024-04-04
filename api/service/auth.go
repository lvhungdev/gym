package service

import (
	"errors"
	"github.com/lvhungdev/gym/domain/entity"
)

type AuthService struct {
	isSignedIn bool
	user       entity.User
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GetSignedInUser() (entity.User, error) {
	if !s.isSignedIn {
		return entity.User{}, errors.New("user not signed in")
	}
	return s.user, nil
}

func (s *AuthService) SetSignedInUser(user entity.User) {
	s.isSignedIn = true
	s.user = user
}
