package main

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.SetLogger("file", `{"filename":"/root/mha-handlers/logs/mha-handlers.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	SessionAndChecks()
}
