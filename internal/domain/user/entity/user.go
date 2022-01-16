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
