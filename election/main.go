package main

import (
	//	"fmt"
	//	"database/sql"
	//	"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/election.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	//beego.Trace(conf)
	for {
		err := GetLeader()
		if err != nil {
			beego.Error("Failure to obtain master information:", err)
			break
		}
	}
}
