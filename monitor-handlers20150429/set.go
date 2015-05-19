package main

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Set(ip, port, username, password string, issynchronous int) {
	dsName := username + ":" + password + "@tcp(" + "localhost" + ":" + port + ")/"
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		beego.Error("Connection to the database failure!", err)
		return
	}
	beego.Info("Connection to the database success!")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		beego.Error("ping() database  failure!", err)
		return
	}
	beego.Info("ping() database success!")
	keepsyncrepl := "set global rpl_semi_sync_master_keepsyncrepl=" + string(issynchronous)
	_, err = db.Query(keepsyncrepl)
	if err != nil {
		beego.Error("Set rpl_semi_sync_master_keepsyncrepl=" + keepsyncrepl + "  failure!")
		return
	}
	beego.Info("Set rpl_semi_sync_master_keepsyncrepl=" + keepsyncrepl + " success!")
	trysyncrepl := "set global rpl_semi_sync_master_trysyncrepl=" + string(issynchronous)
	_, err = db.Query(trysyncrepl)
	if err != nil {
		beego.Error("Set rpl_semi_sync_master_trysyncrepl=" + trysyncrepl + " failure!")
		return
	}
	beego.Info("Set rpl_semi_sync_master_trysyncrepl=" + trysyncrepl + " success!")
}
