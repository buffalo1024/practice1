package models

import "github.com/jmoiron/sqlx"

func InsertUser(db *sqlx.DB, user User) (err error)  {
	_,err = db.Exec("insert into users values (?,?)",user.Username,user.Password)
	return
}
