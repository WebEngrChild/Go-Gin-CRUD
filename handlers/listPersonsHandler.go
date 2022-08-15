package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"recipe-api/models"

	"github.com/gin-gonic/gin"
)

func ListPersonsHandler(c *gin.Context) {
	persons := make([]models.Person, 10)

	c.JSON(http.StatusOK, persons)
}

func getRows(db *sql.DB, p *models.Person) {
	rows, err := db.Query("SELECT * FROM persons")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Gender, &p.Birthday, &p.Phone); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}
