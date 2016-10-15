package model

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"` // false -> female, true -> male
}
