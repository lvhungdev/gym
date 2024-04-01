package entity

type User struct {
	Id       string
	Email    string
	Password string
	FullName string
}

func (u User) WithoutPassword() User {
	user := u
	user.Password = ""

	return user
}
