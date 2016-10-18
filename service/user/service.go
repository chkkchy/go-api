package user

import (
	userDao "../../dao"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"` // false -> female, true -> male
}

func (s *User) findUsers(sex string) []User {
	// TODO: filter by sex
	ents := userDao.FindAll(true)
	var users []User
	for _, ent := range ents {
		user := eToM(ent)
		users = append(users, user)
	}
	return users
}

func (s *User) findUser(id int) User {
	ent := userDao.FindOne(id)
	user := eToM(ent)
	return user
}

func (s *User) createUser(u *User) *User {
	ent := userDao.Insert(u.Name, u.Age, u.Sex)
	user := eToM(*ent)
	return &user
}

func (s *User) updateUser(id int, name string, age int) *User {
	ent := userDao.Update(id, name, age)
	user := eToM(*ent)
	return &user
}

func (s *User) deleteUser(id int) bool {
	res := userDao.Delete(id)
	return res
}

func eToM(e userDao.Users) User {
	m := User{}
	m.Id = e.Id
	m.Name = e.Name
	m.Age = e.Age
	m.Sex = e.Sex
	return m
}
