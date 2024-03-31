package port

import "github.com/lvhungdev/gym/domain/entity"

type UserRepo interface {
	GetUserById(id string) (entity.User, bool)
	GetUserByEmail(email string) (entity.User, bool)
	CreateNewUser(user entity.User) error
}
