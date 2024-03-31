package infras

import "github.com/google/uuid"

type PasswordHasher struct{}

func NewPasswordHasher() PasswordHasher {
	return PasswordHasher{}
}

func (p PasswordHasher) HashPassword(password string) string {
	return password
}

func (p PasswordHasher) VerifyPassword(original string, hashed string) bool {
	return original == hashed
}

type IdGenerator struct{}

func NewIdGenerator() IdGenerator {
	return IdGenerator{}
}

func (g IdGenerator) Generate() string {
	id := uuid.New()
	return id.String()
}
