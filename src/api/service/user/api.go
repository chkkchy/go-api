package user

import (
	"time"

	userDao "api/dao"

	_ "github.com/go-sql-driver/mysql"
)

func (s *Service) GetUsers(input *GetUsersInput) (output *UserOutputList) {
	output = &UserOutputList{}
	// TODO: filter by sex
	ents := userDao.FindAll(true)
	for _, ent := range ents {
		user := eToM(ent)
		output.List = append(output.List, user)
	}
	return
}

func (s *Service) GetUser(id int) (output *UserOutput) {
	ent := userDao.FindOne(id)
	output = eToM(ent)
	return
}

func (s *Service) CreateUser(input *CreateUserInput) (output *UserOutput) {
	ent := userDao.Insert(input.Name, input.Age, input.Sex)
	output = eToM(ent)
	return
}

func (s *Service) UpdateUser(input *UpdateUserInput) (output *UserOutput) {
	ent := userDao.Update(input.Id, input.Name, input.Age)
	output = eToM(ent)
	return
}

func (s *Service) DeleteUser(id int) (output bool) {
	output = userDao.Delete(id)
	return
}

func eToM(e *userDao.Users) (m *UserOutput) {
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
