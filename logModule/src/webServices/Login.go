package webServices

import (
	"../dbUtils"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
	var username string
	username = r.Form["username"][0]
	fmt.Println(r.Form["password"][0]==dbUtils.QueryPassword(username))
}
