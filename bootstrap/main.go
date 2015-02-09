package main

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.SetLogger("file", `{"filename":"/home/lindan/lindan/github/consul/bootstrap/logs/bootstrap.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	SetConn()
}
