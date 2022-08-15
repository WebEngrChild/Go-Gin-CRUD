package models

type Person struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Gender     bool   `gorm:"column:gender"`
	Birthday   string `gorm:"column:birthday"`
	Phone      string `gorm:"column:phone"`
	Company    string `gorm:"column:company"`
	Department string `gorm:"column:department"`
	Branch     string `gorm:"column:branch"`
}
