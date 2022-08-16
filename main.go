// main.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"recipe-api/Util"
	"recipe-api/routes"

	_ "github.com/go-sql-driver/mysql"
	"recipe-api/models"
)

func main() {
	defer Util.Db.Close()

	r := gin.Default()
	r.Run(":3000")
	router := routes.NewRoutes()
	router.Run()
}

func getSingleRow(db *sql.DB, p *models.Person, id int) {
	rows, err := db.Query("SELECT p.id, p.name, p.gender, p.birthday, p.phone, c.name AS company, d.name AS department, b.name AS branch "+
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
