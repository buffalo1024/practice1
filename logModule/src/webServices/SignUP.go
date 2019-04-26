package webServices

import (
	"../dbUtils"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	var user dbUtils.User
	user.Username = username
	user.Password = password
	dbUtils.InsertUser(user)
}
