package middlewares

import (
	"../config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Mysql(dbName string, tomlConfig *config.Config) gin.HandlerFunc {
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	db, err := sqlx.Open("mysql", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("sqlx.Open: err:%v", err))
	}

	db.SetMaxIdleConns(15)

	return func(c *gin.Context) {
		c.Set(dbName, db)
		c.Next()
	}
}
