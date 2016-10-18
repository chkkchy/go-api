package user

import (
	"os"
	"time"

	m "../model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/golang/glog"
)

type Users struct {
	Id        int
	Name      string
	Age       int
	Sex       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

var engine *xorm.Engine

func init() {
	// ORM Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "testuser:testpass@tcp(ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:3306)/testdb")
	if err != nil {
		log.Error(err.Error())
	}
	// _ = engine.Sync2(new(Users))

	// Logs
	f, err := os.Create("logs/sql.log")
	if err != nil {
		log.Error(err.Error())
		return
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetLogger(xorm.NewSimpleLogger(f))

	// Connection pool
	engine.SetMaxIdleConns(100) // idleの総数
	engine.SetMaxOpenConns(500) // 全体の総数
}

func FindAll(sex bool) []m.User {
	var us []m.User
	err := engine.Iterate(&Users{}, func(idx int, bean interface{}) error {
		e := bean.(*Users)
		m := eToM(*e)
		us = append(us, m)
		return nil
	})
	if err != nil {
		log.Error(err.Error())
	}
	return us
}

func FindOne(id int) m.User {
	u := &Users{}
	engine.Where("id = ?", id).Get(u)
	ret := eToM(*u)
	return ret
}

func Insert(user *m.User) *m.User {
	u := &Users{Name: user.Name, Age: user.Age, Sex: user.Sex, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	affected, err := engine.Insert(u)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		log.Info("nothing affected")
		return nil
	}
	// TODO: set created user's id
	return user
}

func Update(id int, name string, age int) *m.User {
	u := &Users{Name: name, Age: age}
	affected, err := engine.Where("id =?", id).Update(u)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		log.Info("nothing affected")
		return nil
	}
	engine.Where("id = ?", id).Get(u)
	ret := eToM(*u)
	return &ret
}

func Delete(id int) bool {
	affected, err := engine.Where("id = ?", id).Delete(&Users{})
	if err != nil {
		log.Error(err.Error())
	}
	return affected > 0
}

func eToM(e Users) m.User {
	m := m.User{}
	m.Id = e.Id
	m.Name = e.Name
	m.Age = e.Age
	m.Sex = e.Sex
	return m
}
