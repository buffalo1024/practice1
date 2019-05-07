package handlers

import (
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SignUpParams struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"'`
}

func SignUp(c *gin.Context)  {
	var req SignUpParams
	if err := c.ShouldBind(&req); err != nil {
		//绑定出错
		c.Writer.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	var user models.User
	user.Username = req.Username
	user.Password = req.Password
	userdb := c.MustGet(UserDB).(*sqlx.DB)
	if err := models.InsertUser(userdb, user); err != nil {
		// 错误处理
		c.Writer.Write([]byte(fmt.Sprintf("error: %v\n", err)))
	}
	// 正常
	c.Writer.Write([]byte("Signed up successfully."))
}
