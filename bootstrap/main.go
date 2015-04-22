package main

import (
	"github.com/astaxie/beego"
	"os"
	"time"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/bootstrap.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			beego.Info("version 0.5.0")
			return
		} else {
			return
		}
	}
	SetConn()
}
