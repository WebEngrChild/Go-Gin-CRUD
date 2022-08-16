package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"recipe-api/models"
	"recipe-api/util"
)

func SearchPersonsHandler(c *gin.Context) {
	p := &models.Person{}
	rows, err := util.Db.Query("SELECT p.id, p.name, p.gender, p.birthday, p.phone, c.name AS company, d.name AS department, b.name AS branch "+
		"FROM golang.persons AS p "+
		"JOIN employees AS e on p.id = e.person_id "+
		"JOIN companies AS c on c.id = e.company_id "+
		"JOIN departments AS d on d.id = e.department_id "+
		"JOIN branches AS b on e.branch_id = b.id "+
		"WHERE p.id = ?", id)

	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Gender, &p.Birthday, &p.Phone, &p.Company, &p.Department, &p.Branch); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}

//func SearchPersonsHandler(c *gin.Context) {
//	// リクエストからtagを取得する
//	// ?tag=XXXXの様に指定するだけで値を取得可能
//	tag := c.Query("tag")
//
//	listOfRecipes := make([]models.Recipe, 0)
//
//	// レシピをループ
//	for i := 0; i < len(recipes); i++ {
//		found := false
//		// レシピ内のタグを抽出
//		for _, t := range recipes[i].Tags {
//			// タグに合致しているか判定
//			if strings.EqualFold(t, tag) {
//				found = true
//			}
//		}
//		if found {
//			listOfRecipes = append(listOfRecipes, recipes[i])
//		}
//	}
//	c.JSON(http.StatusOK, listOfRecipes)
//}
