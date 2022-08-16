package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"recipe-api/Util"
	"recipe-api/models"
)

func ListPersonsHandler(c *gin.Context) {
	p := &models.Person{}
	rows, err := Util.Db.Query("SELECT p.id, p.name, p.gender, p.birthday, p.phone, c.name AS company, d.name AS department, b.name AS branch " +
		"FROM golang.persons AS p " +
		"JOIN employees AS e on p.id = e.person_id " +
		"JOIN companies AS c on c.id = e.company_id " +
		"JOIN departments AS d on d.id = e.department_id " +
		"JOIN branches AS b on e.branch_id = b.id ")

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

	c.JSON(http.StatusOK, p)
}
