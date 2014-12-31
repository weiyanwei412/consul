package main

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"/root/handlers/logs/handlers.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	PUT()
}
