package main

import (
//	"time"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"/home/lindan/lindan/github/session-test/logs/session-test.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	RegisterService()
//	time.Sleep(5000000000)
	KvCreate()
	session := SessionCearte()
	beego.Info(session)
//	time.Sleep(500000000)
	KvAcquire(session)
	KvRelease(session)
}
