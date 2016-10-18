package dao

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/golang/glog"
)

// type Entity Users

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

func FindAll(sex bool) []*Users {
	var ents []*Users
	err := engine.Iterate(&Users{}, func(idx int, bean interface{}) error {
		ent := bean.(*Users)
		ents = append(ents, ent)
		return nil
	})
	if err != nil {
		log.Error(err.Error())
	}
	return ents
}

func FindOne(id int) *Users {
	ent := &Users{}
	engine.Where("id = ?", id).Get(ent)
	return ent
}

func Insert(name string, age int, sex bool) *Users {
	ent := &Users{Name: name, Age: age, Sex: sex, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	affected, err := engine.Insert(ent)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		log.Info("nothing affected")
		return nil
	}
	// TODO: set created user's id
	return ent
}

func Update(id int, name string, age int) *Users {
	ent := &Users{Name: name, Age: age}
	affected, err := engine.Where("id =?", id).Update(ent)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if affected == 0 {
		log.Info("nothing affected")
		return nil
	}
	engine.Where("id = ?", id).Get(ent)
	return ent
}

func Delete(id int) bool {
	affected, err := engine.Where("id = ?", id).Delete(&Users{})
	if err != nil {
		log.Error(err.Error())
	}
	return affected > 0
}
