package infras

import "github.com/lvhungdev/gym/domain"

func ResolveAuthUC() domain.AuthUC {
	repo := NewUserRepo()
	passwordHasher := NewPasswordHasher()
	idGenerator := NewIdGenerator()

	return domain.NewAuthUC(repo, passwordHasher, idGenerator)
}
