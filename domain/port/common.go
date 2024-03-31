package port

type PasswordHasher interface {
	HashPassword(password string) string
	VerifyPassword(original string, hashed string) bool
}

type IdGenerator interface {
	Generate() string
}
