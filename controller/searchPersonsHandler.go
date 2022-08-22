package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-crud/models"
	"go-gin-crud/util"
)

func SearchPersonsHandler(c *gin.Context) {

	id := c.Query("id")

	rows, err := util.Db.Query("SELECT p.id AS id, p.name AS name, p.gender AS gender, p.birthday as birthday, p.phone AS phone, c.name AS company, d.name AS department, b.name AS branch "+
		"FROM golang.persons AS p "+
		"JOIN employees AS e on p.id = e.person_id "+
		"JOIN companies AS c on c.id = e.company_id "+
		"JOIN departments AS d on d.id = e.department_id "+
		"JOIN branches AS b on e.branch_id = b.id "+
		"WHERE p.id = ?", id)

	if err != nil {
		log.Fatalf("SearchPersonsHandler db.Query error err:%v", err)
	}
	defer rows.Close()

	persons := make([]*models.Person, 0)
	for rows.Next() {
		p := &models.Person{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Gender, &p.Birthday, &p.Phone, &p.Company, &p.Department, &p.Branch); err != nil {
			log.Fatalf("SearchPersonsHandler rows.Scan error err:%v", err)
		}
		persons = append(persons, p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("SearchPersonsHandler rows.Err error err:%v", err)
	}

	c.JSON(http.StatusOK, persons)
}
