package models

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	Username string `db:"username" json:"name"`
	Password string `db:"password" json:"password"`
}

func QueryData(db *sqlx.DB) ([]User,error){
	users := make([]User,0)
	err := db.Select(&users,"select * from users")
	return users,err
}
