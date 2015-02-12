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
		beego.Error("Connect to the innosql failed. Error: ", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		beego.Error("ping innosql fails. Error: ", err)
		return
	}
	//	beego.Info("连接mysql数据库成功")
	_, err = db.Query("stop slave io_thread")
	if err != nil {
		beego.Error("stop slave io_thread failed. Error: ", err)
		return
	}
	//	beego.Info("关闭slave线程成功")
	row, err := db.Query("show slave status")
	if err != nil {
		beego.Error("select slave status failed. Error: ", err)
		return
	}
	//	beego.Info("查询slave状态成功")
	cols, _ := row.Columns()
	buffer := make([]interface{}, len(cols))
	data := make([]interface{}, len(cols))
	for i, _ := range buffer {
		buffer[i] = &data[i]
	}
	for row.Next() {
		err = row.Scan(buffer...)
		if err != nil {
			beego.Error("scan err:", err)
			return
		}
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
		beego.Error("SQL Copy the thread is not normal! Error: ", err)
		return
	}
	//	beego.Info("SQL的复制线程正常")
	sqlstr := "select master_pos_wait(?,?)"
	rowss, err := db.Query(sqlstr, Master_Log_File, Read_Master_Log_Pos)
	if err != nil {
		beego.Error("Carried out master_pos_wait function failed. Error: ", err)
		return
	}
	//	beego.Info("执行master_pos_wait函数成功")
	var master_pos_wait string
	for rowss.Next() {
		err = rowss.Scan(&master_pos_wait)
		if err != nil {
			beego.Error("scan err", err)
			return
		}
		if master_pos_wait < "0" && master_pos_wait == "null" {
			beego.Error("Switching Database failed! Error: ", err)
			return
		}
		//		beego.Info("切换数据库成功")
		_, err := db.Query("set global read_only=0")
		if err != nil {
			beego.Error("Set readable and writable failed! Error: ", err)
			return
		}
		//		beego.Info("设置可读可写成功")
	}
	SetConn(ip, port, username, password)
}
