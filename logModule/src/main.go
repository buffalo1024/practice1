package main

import (
	"./config"
	"./handlers"
	"./middlewares"
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
		log.Println("err: %v", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.Set(handlers.Config, tomlConfig))
	router.Use(middlewares.Mysql(handlers.UserDB, tomlConfig))

	router.GET("/printdata", handlers.PrintData)
	router.GET("/", handlers.Login)
	router.POST("/", handlers.Login)
	router.POST("signup", handlers.SignUp)

	if err := router.Run(tomlConfig.GetListenAddr()); err != nil {
		log.Fatalf("run err:%v\n", err)
	}
}