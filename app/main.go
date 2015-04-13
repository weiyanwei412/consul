package main

import (
	//	"fmt"
	//	"database/sql"
	//	"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/election.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	//beego.Trace(conf)
	args := os.Args[1:]
	for _, arg := range args {
		//              if arg == "--" {
		//                      break
		//              }
		if arg == "-v" || arg == "--version" {
			beego.Info("version 0.4.0")
			return
		}
	}

	for {
		err := GetLeader()
		if err != nil {
			beego.Error("Failure to obtain master information:", err)
			break
		}
	}
}
