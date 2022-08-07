// main.go
package main

import "recipe-api/routes"

func main() {
	router := routes.NewRoutes()
	router.Run()
}
