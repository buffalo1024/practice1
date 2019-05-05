package models

import (
	"github.com/jmoiron/sqlx"
)

func QueryPassword(db *sqlx.DB, username string) (string,error) {
	psql := "select password from users where username =  ? "
	var pwd string
	err := db.Get(&pwd, psql, username)
	return pwd,err
}
