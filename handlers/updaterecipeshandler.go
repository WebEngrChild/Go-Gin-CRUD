package handlers

import (
	"net/http"
	"recipe-api/models"

	"github.com/gin-gonic/gin"
)

func UpdateRecipeHandler(c *gin.Context) {
	// urlに指定するrecipes/:idからidを取得できる
	id := c.Param("id")
	var recipe models.Recipe

	// データバインディングを行う
	if err := c.ShouldBindJSON(&recipe); err != nil {
		// エラーの場合は、400エラーを返す
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// 更新するレシピのIndexをrecipesから探す
	// インデックスが0の場合があるため-1
	targetRecipeIndex := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			targetRecipeIndex = i
		}
	}

	// もし更新するレシピがない場合
	if targetRecipeIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	// 更新する
	recipes[targetRecipeIndex] = recipe
	c.JSON(http.StatusOK, recipe)
}
