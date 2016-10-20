package v1

import (
	"time"

	userdao "core/dao/user"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(input *GetUsersInput) (output *UserOutputList) {
	output = &UserOutputList{}
	// TODO: filter by sex
	ents := userdao.FindAll(true)
	for _, ent := range ents {
		user := eToM(ent)
		output.List = append(output.List, user)
	}
	return
}

func GetUser(id int) (output *UserOutput) {
	ent := userdao.FindOne(id)
	output = eToM(ent)
	return
}

func CreateUser(input *CreateUserInput) (output *UserOutput) {
	ent := userdao.Insert(input.Name, input.Age, input.Sex)
	output = eToM(ent)
	return
}

func UpdateUser(input *UpdateUserInput) (output *UserOutput) {
	ent := userdao.Update(input.Id, input.Name, input.Age)
	output = eToM(ent)
	return
}

func DeleteUser(id int) (output bool) {
	output = userdao.Delete(id)
	return
}

func eToM(e *userdao.Users) (m *UserOutput) {
	m = &UserOutput{}
	m.Id = e.Id
	m.Name = e.Name
	m.Age = e.Age
	m.Sex = e.Sex
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	return
}

type GetUsersInput struct {
	Sex bool `json:"sex"` // false -> female, true -> male
	_   int
	_   bool
}

type CreateUserInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"` // false -> female, true -> male
}

type UpdateUserInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserOutput struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Sex       bool      `json:"sex"` // false -> female, true -> male
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOutputList struct {
	List []*UserOutput `json:"list"`
}
