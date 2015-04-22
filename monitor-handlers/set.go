package main

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Set(ip, port, username, password string) {
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
	_, err = db.Query("set global rpl_semi_sync_master_keepsyncrepl=0")
	if err != nil {
		beego.Error("Set rpl_semi_sync_master_keepsyncrepl=0 failure!")
		return
	}
	beego.Info("Set rpl_semi_sync_master_keepsyncrepl=0 success!")
	_, err = db.Query("set global rpl_semi_sync_master_trysyncrepl=0")
	if err != nil {
		beego.Error("Set rpl_semi_sync_master_trysyncrepl=0 failure!")
		return
	}
	beego.Info("Set rpl_semi_sync_master_trysyncrepl=0 success!")
}
