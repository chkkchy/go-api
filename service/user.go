package user

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"bool"` // false -> male, true -> female
}

var users map[int]User = map[int]User{
	1: User{"yukichi", 26, false},
	2: User{"mendy", 25, true},
	3: User{"shibata", 28, false},
}

func FindUsers(sex string) []User {
	// TODO: filter by sex
	ret := mToS(users)
	return ret
}

func FindUser(id int) User {
	return users[id]
}

func CreateUser(user User) User {
	users[(len(users) + 1)] = user
	return user
}

func UpdateUser(id int, name string, age int) User {
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

func mToS(m map[int]User) []User {
	v := make([]User, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}
