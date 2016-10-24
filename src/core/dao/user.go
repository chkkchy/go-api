package dao

import "time"

// log "github.com/golang/glog"

// type Entity Users

type Users struct {
	Id        int
	Name      string
	Age       int
	Sex       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindUsers(sex bool) []*Users {
	var ents []*Users
	err := engine.Iterate(&Users{}, func(idx int, bean interface{}) error {
		ent := bean.(*Users)
		ents = append(ents, ent)
		return nil
	})
	if err != nil {
		// log.Error(err.Error())
	}
	return ents
}

func FindUser(id int) *Users {
	ent := &Users{}
	engine.Where("id = ?", id).Get(ent)
	return ent
}

func InsertUser(name string, age int, sex bool) *Users {
	ent := &Users{Name: name, Age: age, Sex: sex, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	affected, err := engine.Insert(ent)
	if err != nil {
		// log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		// log.Info("nothing affected")
		return nil
	}
	// TODO: set created user's id
	return ent
}

func UpdateUser(id int, name string, age int) *Users {
	ent := &Users{Name: name, Age: age}
	affected, err := engine.Where("id =?", id).Update(ent)
	if err != nil {
		// log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		// log.Info("nothing affected")
		return nil
	}
	engine.Where("id = ?", id).Get(ent)
	return ent
}

func DeleteUser(id int) bool {
	affected, err := engine.Where("id = ?", id).Delete(&Users{})
	if err != nil {
		// log.Error(err.Error())
	}
	return affected > 0
}
