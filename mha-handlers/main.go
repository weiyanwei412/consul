package main

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"/root/mha-handlers/logs/mha-handlers.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	Empty()
	slave()
	SetConn()
}
