package main

import (
//	"fmt"
	//	"database/sql"
	//	"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"/home/lindan/lindan/github/election/logs/election.log"}`)
	beego.SetLogFuncCall(true)
	beego.Info(`
	LevelEmergency	0
    	LevelAlert	1
   	LevelCritical	2
    	LevelError	3
    	LevelWarning	4
    	LevelNotice	5
    	LevelInformational  6
    	LevelDebug	7 //default`)
}

func main() {
	defer beego.BeeLogger.Close()
	//beego.Trace(conf)
	for ; ;{
		err := GetLeader()
		if err != nil {
                	beego.Error("Failure to obtain master information:", err)	
        	}
	}
}
