package dbUtils

func QueryPassword(username string) string {
	DB := InitDB()
	row, _ := DB.Query("select password from users where username=?",username)
	var pwd string
	for row.Next() {
		row.Scan(&pwd)
	}
	//row.Scan(&pwd)
	//fmt.Println(pwd)
	return pwd
}
