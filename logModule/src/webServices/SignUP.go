package webServices

import (
	"../dbUtils"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SignUp(c *gin.Context)  {
	r := c.Request
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	var user dbUtils.User
	user.Username = username
	user.Password = password
	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	if err := dbUtils.InsertUser(userdb, user); err != nil {
		// 错误处理
	}
	// 正常
}
