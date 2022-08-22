package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/controller"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// 従業員データ取得
	router.GET("/persons", controller.ListPersonsHandler)
	// 従業員データの検索
	router.GET("/persons/search", controller.SearchPersonsHandler)
	return router
}
