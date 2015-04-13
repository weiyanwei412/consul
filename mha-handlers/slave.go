package main

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func slave(ip, port, username, password string) {
	dsName := username + ":" + password + "@tcp(" + "localhost" + ":" + port + ")/"
	//	db, err := sql.Open("mysql", "root:111111@tcp(192.168.2.61:3306)/")
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
	_, err = db.Query("stop slave io_thread")
	if err != nil {
		beego.Error("stop slave io_thread failure!", err)
		return
	}
	beego.Info("stop slave io_thread success!")
	row, err := db.Query("show slave status")
	if err != nil {
		beego.Error("Inquiry slave status failure!", err)
		return
	}
	beego.Info("Inquiry slave status success!")
	cols, _ := row.Columns()
	buffer := make([]interface{}, len(cols))
	data := make([]interface{}, len(cols))
	for i, _ := range buffer {
		buffer[i] = &data[i]
	}
	for row.Next() {
		err = row.Scan(buffer...)
		if err != nil {
			beego.Error("scan() traversal slave status failure!", err)
			return
		}
		//beego.Info("scan() traversal slave status success!")
	}
	mapField2Data := make(map[string]interface{}, len(cols))
	for k, col := range data {
		mapField2Data[cols[k]] = col
	}
	Master_Log_File := mapField2Data["Master_Log_File"]
	Read_Master_Log_Pos := mapField2Data["Read_Master_Log_Pos"]
	Slave_SQL_Running := mapField2Data["Slave_SQL_Running"]
	//	fmt.Printf("Slave_SQL_Running=%s\n", Slave_SQL_Running)
	if string(Slave_SQL_Running.([]uint8)) != "Yes" {
		beego.Error("SQL copy the thread is not normal! ", err)
		return
	}
	beego.Info("SQL copy the thread Normal!")
	sqlstr := "select master_pos_wait(?,?)"
	rowss, err := db.Query(sqlstr, Master_Log_File, Read_Master_Log_Pos)
	if err != nil {
		beego.Error("Carried out master_pos_wait function failure!", err)
		return
	}
	beego.Info("Carried out master_pos_wait function success!")
	var master_pos_wait string
	for rowss.Next() {
		err = rowss.Scan(&master_pos_wait)
		if err != nil {
			beego.Error("scan() traversal master_pos_wait data failure!", err)
			return
		}
		//	beego.Info("scan() traversal master_pos_wait data success!")
		if master_pos_wait < "0" && master_pos_wait == "null" {
			beego.Error("Switching database failure!", err)
			return
		}
		beego.Info("Switching database success!")
		_, err := db.Query("set global rpl_semi_sync__master_keepsyncrepl=0")
		if err != nil {
			beego.Error("Set rpl_semi_sync__master_keepsyncrepl=0 failure!")
			return
		}
		beego.Info("Set rpl_semi_sync__master_keepsyncrepl=0 success!")
		_, err = db.Query("set global rpl_semi_sync_master_trysyncrepl=0")
		if err != nil {
			beego.Error("Set rpl_semi_sync_master_trysyncrepl=0 failure!")
			return
		}
		beego.Info("Set rpl_semi_sync_master_trysyncrepl=0 success!")
		_, err = db.Query("set global read_only=0")
		if err != nil {
			beego.Error("Set read and write failure!", err)
			return
		}
		beego.Info("Set read and write success!")
	}
	SetConn(ip, port, username, password)
}
