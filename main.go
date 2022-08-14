// main.go
package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	dotnet "github.com/joho/godotenv"
	"log"
	"os"
	"recipe-api/models"
)

func main() {

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

	p := &models.Person{}

	getRows(db, p)
	getSingleRow(db, p, 1)

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

func getSingleRow(db *sql.DB, p *models.Person, id int) {
	err := db.QueryRow("SELECT * FROM persons WHERE id = ?", id).
		Scan(&p.Id, &p.Name, &p.Gender, &p.Birthday, &p.Phone)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("getSingleRow no records.")
		return
	}
	if err != nil {
		log.Fatalf("getSingleRow db.QueryRow error err:%v", err)
	}
	fmt.Println(p)
}
