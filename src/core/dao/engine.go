package dao

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// log "github.com/golang/glog"
)

var engine *xorm.Engine

func init() {
	// ORM Engine
	var err error
	engine, err = xorm.NewEngine(
		"mysql",
		"testuser:testpass@tcp(ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:3306)/testdb",
	)
	if err != nil {
		// log.Error(err.Error())
		fmt.Println(err.Error())
	}
	// _ = engine.Sync2(new(Users))

	// Logs
	f, err := os.Create("src/api/logs/sql.log")
	if err != nil {
		// log.Error(err.Error())
		fmt.Println(err.Error())
		return
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetLogger(xorm.NewSimpleLogger(f))

	// Connection pool
	engine.SetMaxIdleConns(100) // idleの総数
	engine.SetMaxOpenConns(500) // 全体の総数
}
