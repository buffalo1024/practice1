package webServices

import (
	"../dbUtils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var DB *sql.DB
var users []dbUtils.User

func PrintData(c *gin.Context)  {
	DB = dbUtils.InitDB()
	defer DB.Close()
	users = dbUtils.QueryData(DB)
	for _, user := range users {
		s := []string{"username: ",user.Username, " ,password: ",user.Password,"\n"}
		fmt.Fprintf(c.Writer,strings.Join(s,""))
	}
}
