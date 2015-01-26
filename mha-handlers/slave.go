package main

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func slave() {
	db, err := sql.Open("mysql", "root:111111@tcp(192.168.2.62:3306)/")
	if err != nil {
		beego.Error("Connect to the mysql fails:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		beego.Error("db.ping():", err)
	}
	_, err = db.Query("stop slave io_thread")
	if err != nil {
		beego.Error("show databases:", err)
	}

}
