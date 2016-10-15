package user

import (
	"fmt"
	"time"

	m "../model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var users map[int]m.User = map[int]m.User{
	1: m.User{"yukichi", 26, false},
	2: m.User{"mendy", 25, true},
	3: m.User{"shibata", 28, false},
}

type User struct {
	Id        int
	Name      string
	Age       int
	Sex       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindUsers(sex string) []m.User {
	// TODO: filter by sex
	ret := mToS(users)
	return ret
}

func FindUser(id int) m.User {
	return users[id]
}

func CreateUser(user *m.User) *m.User {
	users[(len(users) + 1)] = *user
	engine, err := xorm.NewEngine("mysql", "testuser:testpass@ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com/testdb")
	if err != nil {
		fmt.Println("e1: ", err)
	}
	u := &User{Name: user.Name, Age: user.Age, Sex: user.Sex}
	affected, err := engine.Insert(u)
	if err != nil {
		fmt.Println("e2: ", err)
	}
	fmt.Println("affected: ", affected)
	return user
}

func UpdateUser(id int, name string, age int) m.User {
	user := users[id]
	user.Name = name
	user.Age = age
	users[id] = user
	return user
}

func DeleteUser(id int) bool {
	delete(users, id)
	_, ok := users[id]
	deleted := !ok
	return deleted
}

func mToS(mp map[int]m.User) []m.User {
	v := make([]m.User, 0, len(mp))
	for _, value := range mp {
		v = append(v, value)
	}
	return v
}
