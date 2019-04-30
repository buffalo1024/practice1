package webServices

import (
	"../dbUtils"
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
	//fmt.Println("method: ", r.Method)
	//fmt.Println("url: ",r.URL)
	//fmt.Println("header: ",r.Header)
	//fmt.Println("postform: ", r.PostForm)
	//fmt.Println("form: ", r.Form)
	//fmt.Println("close: ", r.Close)
	//fmt.Println("body: ", r.Body)
	//fmt.Println("contentlength: ", r.ContentLength)
	//fmt.Println("host: ", r.Host)
	//fmt.Println("proto: ", r.Proto)
	//fmt.Println("protomajor: ", r.ProtoMajor)
	//fmt.Println("protominor: ", r.ProtoMinor)
	//fmt.Println("remoteaddr: ", r.RemoteAddr)
	//fmt.Println("requesturi: ", r.RequestURI)
	//fmt.Println("username: ", r.Form["username"])
	//fmt.Println("password: ", r.Form["password"])
	//fmt.Println(r.Form["password"][0]==dbUtils.QueryPassword(username))
	userdb := c.MustGet(utils.UserDB).(*sqlx.DB)
	password,err := dbUtils.QueryPassword(userdb, req.Username)
	if err != nil {
		// 错误处理
	}
	success := fmt.Sprintf("%v", req.Password==password)
	w.Write([]byte(success))
}
