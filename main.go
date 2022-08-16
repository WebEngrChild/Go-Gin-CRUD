// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"recipe-api/routes"
	"recipe-api/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer util.Db.Close()

	r := gin.Default()
	r.Run(":3000")
	router := routes.NewRoutes()
	router.Run()
}
