package main

import (
	"../webServices"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/printdata", webServices.PrintData)
	http.HandleFunc("/", webServices.Login)
	http.HandleFunc("/signup", webServices.SignUp)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}