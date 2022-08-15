// main.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	dotnet "github.com/joho/godotenv"
	"log"
	"os"
	"recipe-api/models"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
}

func main() {
	defer Db.Close()
	p := &models.Person{}
	getRows(Db, p)
	getSingleRow(Db, p, 1)

	// Ginの内容
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "OK",
	//	})
	//})

	//r.Run(":3000")
	//router := routes.NewRoutes()
	//router.Run()
}

func sqlConnect() (database *sql.DB, err error) {
	err = dotnet.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	return db, nil
}

func getRows(db *sql.DB, p *models.Person) {
	rows, err := db.Query("SELECT p.id, p.name, p.gender, p.birthday, p.phone, c.name AS company, d.name AS department, b.name AS branch " +
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
