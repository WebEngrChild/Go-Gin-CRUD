package handlers

import (
	"net/http"
	"recipe-api/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func SearchRecipesHandler(c *gin.Context) {
	// リクエストからtagを取得する
	// ?tag=XXXXの様に指定するだけで値を取得可能
	tag := c.Query("tag")

	listOfRecipes := make([]models.Recipe, 0)

	// レシピをループ
	for i := 0; i < len(recipes); i++ {
		found := false
		// レシピ内のタグを抽出
		for _, t := range recipes[i].Tags {
			// タグに合致しているか判定
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}
	c.JSON(http.StatusOK, listOfRecipes)
}
