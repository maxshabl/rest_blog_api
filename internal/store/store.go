package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var instanse *sqlx.DB

func NewConnection(databaseURL string) {
	db, err := sqlx.Open("mysql", databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection open errors %s", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Ping errors %s", err)
	}

	fmt.Fprintln(os.Stdout, "Success connection")

	instanse = db
}

func GetDB() *sqlx.DB {
	return instanse
}
