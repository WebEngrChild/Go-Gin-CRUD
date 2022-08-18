package models

type Person struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     bool   `json:"gender"`
	Birthday   string `json:"birthday"`
	Phone      string `json:"phone"`
	Company    string `json:"company"`
	Department string `json:"department"`
	Branch     string `json:"branch"`
}
