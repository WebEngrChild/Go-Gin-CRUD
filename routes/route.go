package routes

import (
	"github.com/gin-gonic/gin"
	"recipe-api/handlers"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// 従業員データ登録
	router.POST("/recipes", handlers.NewRecipeHandler)
	// 従業員データ取得
	router.GET("/recipes", handlers.ListRecipesHandler)
	// 従業員データ更新
	router.PUT("/recipes/:id", handlers.UpdateRecipeHandler)
	// 従業員データ削除
	router.DELETE("/recipes/:id", handlers.DeleteRecipeHandler)
	// 従業員データの検索
	router.GET("/recipes/search", handlers.SearchRecipesHandler)
	return router
}
