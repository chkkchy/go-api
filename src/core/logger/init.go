package logger

import (
	"flag"

	"github.com/golang/glog"
)

func Init() {
	flag.Parse()
	defer glog.Flush()
}
