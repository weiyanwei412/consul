package main

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.SetLogger("file", `{"filename":"/root/bootstrap/logs/bootstrap.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	SetConn()
}
