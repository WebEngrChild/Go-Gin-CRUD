package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRecipeHandler(c *gin.Context) {
	// リクエストからIDを取得
	id := c.Param("id")

	// インデックスが0の場合があるため-1
	recipeIndex := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			recipeIndex = i
		}
	}

	// レシピが見つからない場合
	if recipeIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "recipe not found"})
		return
	}

	// idが一致したレシピ"以外"をrecipes変数に詰め直す
	recipes = append(recipes[:recipeIndex], recipes[recipeIndex+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})
}
