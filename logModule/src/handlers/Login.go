package handlers

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type LoginParams struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginParams
	if err := c.ShouldBind(&req); err != nil {
		// 绑定出错
		return
	}
	// 绑定成功
	r := c.Request
	w := c.Writer
	r.ParseForm()

	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	password,err := models.QueryPassword(userdb, req.Username)
	if err != nil {
		// 错误处理
	}
	success := fmt.Sprintf("%v", req.Password==password)
	w.Write([]byte(success))
}
