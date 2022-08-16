package routes

import (
	"github.com/gin-gonic/gin"
	"recipe-api/handlers"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// 従業員データ取得
	router.GET("/persons", handlers.ListPersonsHandler)
	// 従業員データの検索
	router.GET("/persons/search", handlers.SearchPersonsHandler)
	return router
}
