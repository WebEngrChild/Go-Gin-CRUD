package handlers

import (
	"net/http"
	"recipe-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// 一時的なtmpDB
// 将来的にはDatabaseに格納したい
var recipes []models.Recipe

func init() {
	recipes = make([]models.Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe models.Recipe
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// ユニークなIDを生成
	recipe.ID = xid.New().String()
	// 現在時刻を追加
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}
