package handlers

import (
	"../models"
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
	w := c.Writer
	if err := c.ShouldBind(&req); err != nil {
		// 绑定出错
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	// 绑定成功
	userdb := c.MustGet(UserDB).(*sqlx.DB)
	password,err := models.QueryPassword(userdb, req.Username)
	if err != nil {
		// 错误处理
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	success := fmt.Sprintf("%v", req.Password==password)
	w.Write([]byte(success))
}
