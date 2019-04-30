package dbUtils

import "github.com/jmoiron/sqlx"

func InsertUser(db *sqlx.DB, user User) (err error)  {
	//DB := InitDB()
	_,err = db.Exec("insert into users values (?,?)",user.Username,user.Password)
	return
}
