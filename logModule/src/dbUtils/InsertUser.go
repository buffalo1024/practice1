package dbUtils

func InsertUser(user User)  {
	DB := InitDB()
	DB.Exec("insert into users values (?,?)",user.Username,user.Password)
}
