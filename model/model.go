package model

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"bool"` // false -> male, true -> female
}
