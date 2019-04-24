package dbUtils

import (
	"database/sql"
	"fmt"
)

func QueryData(DB *sql.DB) {
	rows, _ := DB.Query("select * from users")
	defer rows.Close()
	var username,pwd string
	for rows.Next() {
		rows.Scan(&username,&pwd)
		fmt.Print("username:",username,"; password:",pwd,"\n")
	}
}
