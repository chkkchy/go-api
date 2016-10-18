package user

import (
	userDao "../../dao"
	m "../../model"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
}

func (s *User) findUsers(sex string) []m.User {
	// TODO: filter by sex
	res := userDao.FindAll(true)
	return res
}

func (s *User) findUser(id int) m.User {
	res := userDao.FindOne(id)
	return res
}

func (s *User) createUser(user *m.User) *m.User {
	res := userDao.Insert(user)
	return res
}

func (s *User) updateUser(id int, name string, age int) *m.User {
	res := userDao.Update(id, name, age)
	return res
}

func (s *User) deleteUser(id int) bool {
	res := userDao.Delete(id)
	return res
}
