package dbUtils

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	Username string `db:"username" json:"name"`
	Password string `db:"password" json:"password"`
}

func QueryData(db *sqlx.DB) []User{
	users := make([]User,0)
	rows, _ := db.Query("select * from users")
	defer rows.Close()
	var username,pwd string
	for rows.Next() {
		rows.Scan(&username,&pwd)
		var u User
		u.Username = username
		u.Password = pwd
		//fmt.Print("username:",username,"; password:",pwd,"\n")
		users = append(users, u)
	}
	return users
}
