package user

import (
	userDao "../dao"
	m "../model"
	_ "github.com/go-sql-driver/mysql"
)

func FindUsers(sex string) []m.User {
	// TODO: filter by sex
	res := userDao.FindAll(true)
	return res
}

func FindUser(id int) m.User {
	res := userDao.FindOne(id)
	return res
}

func CreateUser(user *m.User) *m.User {
	res := userDao.Insert(user)
	return res
}

func UpdateUser(id int, name string, age int) *m.User {
	res := userDao.Update(id, name, age)
	return res
}

func DeleteUser(id int) bool {
	res := userDao.Delete(id)
	return res
}
