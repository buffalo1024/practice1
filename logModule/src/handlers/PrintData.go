package handlers

import (
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strings"
)

func PrintData(c *gin.Context)  {
	userdb := c.MustGet(UserDB).(*sqlx.DB)
	users,err := models.QueryData(userdb)
	if err != nil {
		c.Writer.Write([]byte(fmt.Sprintf("error: %v\n", err)))
	}
	for _, user := range users {
		s := []string{"username: ",user.Username, " ,password: ",user.Password,"\n"}
		fmt.Fprintf(c.Writer,strings.Join(s,""))
	}
}
