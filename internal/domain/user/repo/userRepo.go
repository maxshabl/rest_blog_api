package repo

import (
	"blog/internal/domain/user/entity"
	"blog/internal/errors"

	"github.com/jmoiron/sqlx"
)

var tableName = "users"

var UserRepo Repo

type Repo struct {
	db *sqlx.DB
}

func Init(instanse *sqlx.DB) Repo {
	UserRepo.db = instanse
	return UserRepo
}

func (r Repo) FindUserByName(name string) (entity.User, error) {

	var user entity.User

	err := r.db.Get(&user, "SELECT * FROM "+tableName+" WHERE login = ? or email = ?", name, name)
	if err != nil {
		return user, errors.New(err.Error())
	}
	return user, nil
}

func (r Repo) IsLoginEmailExist(login string, email string) bool {

	var userId int64

	err := r.db.Get(&userId, "SELECT id userId FROM "+tableName+" WHERE login = ? or email = ?", login, email)
	
	return err != nil
}

func (r Repo) AddUser(login string, passwordHash string, email string, status int, role int) (int64, error) {
	sql := "INSERT INTO users (login, password_hash, password_reset_token, email, status, role) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(sql, login, passwordHash, passwordHash, email, status, role)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	userId, _ := result.LastInsertId()
	return userId, nil
}
