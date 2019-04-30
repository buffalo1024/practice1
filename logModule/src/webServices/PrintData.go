package webServices

import (
	"../dbUtils"
	"../utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strings"
)

var DB *sql.DB
var users []dbUtils.User

func PrintData(c *gin.Context)  {
	//DB = dbUtils.InitDB()
	//defer DB.Close()
	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	users = dbUtils.QueryData(userdb)
	for _, user := range users {
		s := []string{"username: ",user.Username, " ,password: ",user.Password,"\n"}
		fmt.Fprintf(c.Writer,strings.Join(s,""))
	}
}
