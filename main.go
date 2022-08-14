// main.go
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	dotnet "github.com/joho/godotenv"
	"os"
)

func main() {

	type Persons struct {
		Id       int    `gorm:"column:id"`
		Name     string `gorm:"column:name"`
		Gender   bool   `gorm:"column:gender"`
		Birthday string `gorm:"column:birthday"`
		Phone    string `gorm:"column:phone"`
	}

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

	var result []*Persons
	dbErr := db.Find(&result).Error
	if dbErr != nil || len(result) == 0 {
		return
	}
	for _, user := range result {
		fmt.Println(user.Name)
		fmt.Println(user.Gender)
		fmt.Println(user.Birthday)
		fmt.Println(user.Phone)
	}

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

func sqlConnect() (database *gorm.DB, err error) {
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
	return gorm.Open(DBMS, CONNECT)
}
