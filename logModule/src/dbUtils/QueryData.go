package dbUtils

import (
	"database/sql"
)

type User struct {
	Username string
	Password string
}



func QueryData(DB *sql.DB) []User{
	users := make([]User,0)
	rows, _ := DB.Query("select * from users")
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
