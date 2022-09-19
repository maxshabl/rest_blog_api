package entity

import "database/sql"

type User struct {
	Id                 int64
	Login              string
	Email              string
	PasswordHash       string `db:"password_hash"`
	PasswordResetToken string `db:"password_reset_token"`
	Status             int
	Role               int
	CreatedAt          sql.NullString `db:"created_at"`
	UpdatedAt          sql.NullString `db:"updated_at"`
}

type UserInterface interface {
	FindUserByName(name string) (User, error)
	IsLoginEmailExist(login string, email string) bool
	AddUser(login string, passwordHash string, email string, status int, role int) (int64, error)
}
