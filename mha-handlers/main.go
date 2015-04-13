package main

import (
	"github.com/astaxie/beego"
	"time"
	"os"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/mha-handlers.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	args := os.Args[1:]
	for _, arg := range args {
//		if arg == "--" {
//			break
//		}
		if arg == "-v" || arg == "--version" {
			beego.Info("version 0.4.0")
			return 
		}else{
			return
		}
	}
	SessionAndChecks()
}
