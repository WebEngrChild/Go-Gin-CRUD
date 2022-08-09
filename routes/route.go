package routes

import (
	"github.com/gin-gonic/gin"
	"recipe-api/handlers"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// レシピ作成
	router.POST("/recipes", handlers.NewRecipeHandler)
	// レシピ一覧の取得
	router.GET("/recipes", handlers.ListRecipesHandler)
	// レシピの更新
	router.PUT("/recipes/:id", handlers.UpdateRecipeHandler)
	// レシピの削除
	router.DELETE("/recipes/:id", handlers.DeleteRecipeHandler)
	// レシピの検索
	router.GET("/recipes/search", handlers.SearchRecipesHandler)
	return router
}
