package v1

import (
	"time"

	"core/dao"
)

func GetUsers(input *GetUsersInput) (output *UserOutputList) {
	output = &UserOutputList{}
	// TODO: filter by sex
	ents := dao.FindUsers(true)
	for _, ent := range ents {
		user := eToM(ent)
		output.List = append(output.List, user)
	}
	return
}

func GetUser(id int) (output *UserOutput) {
	ent := dao.FindUser(id)
	output = eToM(ent)
	return
}

func CreateUser(input *CreateUserInput) (output *UserOutput) {
	ent := dao.InsertUser(input.Name, input.Age, input.Sex)
	output = eToM(ent)
	return
}

func UpdateUser(input *UpdateUserInput) (output *UserOutput) {
	ent := dao.UpdateUser(input.Id, input.Name, input.Age)
	output = eToM(ent)
	return
}

func DeleteUser(id int) (output bool) {
	output = dao.DeleteUser(id)
	return
}

func eToM(e *dao.Users) (m *UserOutput) {
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
