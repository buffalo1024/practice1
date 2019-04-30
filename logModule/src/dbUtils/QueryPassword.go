package dbUtils

import (
	"github.com/jmoiron/sqlx"
)

func QueryPassword(db *sqlx.DB, username string) (string,error) {
	//DB := InitDB()
	//row, _ := DB.Query("select password from users where username=?",username)
	//var pwd string
	//for row.Next() {
	//	row.Scan(&pwd)
	//}
	//return pwd

	psql := "select password from users where username =  ? "
	var pwd string
	err := db.Get(&pwd, psql, username)
	return pwd,err
}
