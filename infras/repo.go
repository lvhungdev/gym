package infras

import (
	"database/sql"

	"github.com/lvhungdev/gym/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo() UserRepo {
	db, err := sql.Open("sqlite3", "./gym.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users
        (
            id        TEXT PRIMARY KEY,
            email     TEXT NOT NULL,
            password  TEXT NOT NULL,
            full_name TEXT
        )`,
	)
	if err != nil {
		panic(err)
	}

	return UserRepo{
		db: db,
	}
}

func (r UserRepo) GetUserById(id string) (entity.User, bool) {
	// TODO implement this
	return entity.User{}, false
}

func (r UserRepo) GetUserByEmail(email string) (entity.User, bool) {
	user := entity.User{}

	row := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FullName)
	if err != nil {
		return user, false
	}

	return user, true
}

func (r UserRepo) CreateNewUser(user entity.User) error {
	_, err := r.db.Exec(`
        INSERT INTO users
        (
            id, email, password, full_name
        )
        VALUES (?, ?, ?, ?)
    `, user.Id, user.Email, user.Password, user.FullName)

	return err
}
