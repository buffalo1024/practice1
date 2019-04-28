package webServices

import (
	"../dbUtils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context)  {
	r := c.Request
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	var user dbUtils.User
	user.Username = username
	user.Password = password
	dbUtils.InsertUser(user)
}
