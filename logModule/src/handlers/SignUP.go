package handlers

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SignUp(c *gin.Context)  {
	r := c.Request
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	var user models.User
	user.Username = username
	user.Password = password
	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	if err := models.InsertUser(userdb, user); err != nil {
		// 错误处理
		fmt.Println("err: %v", err)
	}
	// 正常
	c.Writer.Write([]byte("Signed up successfully."))
}
