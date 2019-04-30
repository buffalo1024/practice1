package main

import (
	"../config"
	"../middlewares"
	"../utils"
	"../webServices"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var (
	tomlFile = flag.String("config", "E:/LogAuthenticion/logModule/src/docs/doc.toml", "config file")
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	flag.Parse()
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {

	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.Set(utils.Config, tomlConfig))
	router.Use(middlewares.Mysql(utils.UserDB, tomlConfig))

	//router := gin.Default()
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
	if err := router.Run(tomlConfig.GetListenAddr()); err != nil {
		log.Fatalf("run err:%v\n", err)
	}
}