// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/routes"
	"go-gin-crud/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	SetupServer()
}

func SetupServer() *gin.Engine {
	defer util.Db.Close()

	r := gin.Default()
	r.Run(":3000")
	router := routes.NewRoutes()
	router.Run()

	return r
}
