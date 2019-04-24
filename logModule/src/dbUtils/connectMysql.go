package dbUtils

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip = "127.0.0.1"
	port = "3306"
	dbName = "users"
	)

var DB *sql.DB

func InitDB() * sql.DB {
	path := strings.Join([]string{userName, ":", password,  "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, _ = sql.Open("mysql", path)

	DB.SetConnMaxLifetime(100)

	DB.SetMaxIdleConns(10)

	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	return DB
}
