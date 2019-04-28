package main

import (
	"../webServices"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	router := gin.Default()
	router.GET("/printdata", webServices.PrintData)
	//http.HandleFunc("/printdata", webServices.PrintData)

	//http.HandleFunc("/", webServices.Login)
	router.GET("/", webServices.Login)
	router.POST("/", webServices.Login)

	//http.HandleFunc("/signup", webServices.SignUp)
	//router.GET("/signup",webServices.SignUp)
	router.POST("signup", webServices.SignUp)

	//err := http.ListenAndServe(":9090",nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	router.Run(":9090")
}