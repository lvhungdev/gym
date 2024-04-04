package port

import "github.com/lvhungdev/gym/domain/entity"

type AuthService interface {
	GetSignedInUser() (entity.User, error)
	SetSignedInUser(user entity.User)
}
