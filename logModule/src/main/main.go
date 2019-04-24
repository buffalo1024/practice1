package main

import (
	"../dbUtils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {
	DB := dbUtils.InitDB()
	dbUtils.QueryData(DB)
}