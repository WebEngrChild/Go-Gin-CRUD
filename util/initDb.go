package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	dotnet "github.com/joho/godotenv"
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
