package webServices

import (
	"../dbUtils"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var DB *sql.DB
var users []dbUtils.User

func PrintData(w http.ResponseWriter, r *http.Request)  {
	DB = dbUtils.InitDB()
	defer DB.Close()
	users = dbUtils.QueryData(DB)
	for _, user := range users {
		s := []string{"username: ",user.Username, " ,password: ",user.Password,"\n"}
		fmt.Fprintf(w,strings.Join(s,""))
	}
}
