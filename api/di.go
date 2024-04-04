package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lvhungdev/gym/api/service"
	"github.com/lvhungdev/gym/domain"
	"github.com/lvhungdev/gym/domain/port"
	"github.com/lvhungdev/gym/infras"
)

type ScopedDI struct {
	authUC *domain.AuthUC

	userRepo port.UserRepo

	authService port.AuthService

	idGenerator    port.IdGenerator
	passwordHasher port.PasswordHasher
}

func GetDi(c *gin.Context) *ScopedDI {
	diValue, _ := c.Get("di")
	di := diValue.(*ScopedDI)
	return di
}

func NewScopedDI() *ScopedDI {
	return &ScopedDI{}
}

func (s *ScopedDI) ResolveAuthUC() domain.AuthUC {
	if s.authUC == nil {
		authUC := domain.NewAuthUC(s.ResolveUserRepo(), s.ResolvePasswordHasher(), s.ResolveIdGenerator())
		s.authUC = &authUC
	}
	return *s.authUC
}

func (s *ScopedDI) ResolveUserRepo() port.UserRepo {
	if s.userRepo == nil {
		s.userRepo = infras.NewUserRepo()
	}
	return s.userRepo
}

func (s *ScopedDI) ResolveAuthService() port.AuthService {
	if s.authService == nil {
		s.authService = service.NewAuthService()
	}
	return s.authService
}

func (s *ScopedDI) ResolvePasswordHasher() port.PasswordHasher {
	if s.passwordHasher == nil {
		s.passwordHasher = infras.NewPasswordHasher()
	}
	return s.passwordHasher
}

func (s *ScopedDI) ResolveIdGenerator() port.IdGenerator {
	if s.idGenerator == nil {
		s.idGenerator = infras.NewIdGenerator()
	}
	return s.idGenerator
}
