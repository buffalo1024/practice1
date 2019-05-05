package handlers

import (
	"../models"
	"../utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strings"
)

var DB *sql.DB
var users []models.User

func PrintData(c *gin.Context)  {
	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	users = models.QueryData(userdb)
	for _, user := range users {
		s := []string{"username: ",user.Username, " ,password: ",user.Password,"\n"}
		fmt.Fprintf(c.Writer,strings.Join(s,""))
	}
}
