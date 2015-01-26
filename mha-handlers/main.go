package main

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"/home/lindan/lindan/github/consul/mha-handlers/logs/handlers.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	Empty()
	slave()
}
