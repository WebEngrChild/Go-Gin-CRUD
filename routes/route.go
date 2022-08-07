package routes

import (
	"recipe-api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// レシピ作成
	router.POST("/recipes", handlers.NewRecipeHandler)
	return router
}
